package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"user-service/domain"
	"user-service/repository"
)

type RestHandler struct {
	repo repository.IUserServiceRepository
}

func NewRestHandler(repo repository.IUserServiceRepository) *RestHandler {
	return &RestHandler{repo: repo}
}

func (h *RestHandler) CreateOrGetAllUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllUsers(w, r)
	case http.MethodPost:
		h.CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domain.Response{
			Result: false,
			Error:  "invalid form data",
		})
		return
	}

	name := r.FormValue("name")
	if name == "" {
		errMsg := "name is required"
		log.Println(errMsg)
		respondWithJSON(w, http.StatusBadRequest, domain.Response{
			Result: false,
			Error:  errMsg,
		})
		return
	}

	data, err := h.repo.CreateUser(name)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, domain.Response{
			Result: false,
			Error:  err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusCreated, domain.Response{
		Result: true,
		User:   &data,
	})
}

func (h *RestHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	pageNumStr := r.URL.Query().Get("page_num")
	pageSizeStr := r.URL.Query().Get("page_size")

	if pageNumStr == "" {
		pageNumStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}

	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum < 1 {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domain.Response{
			Result: false,
			Error:  "invalid parameter page_num",
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domain.Response{
			Result: false,
			Error:  "invalid parameter page_size",
		})
		return
	}

	data, err := h.repo.GetAllUsers(pageNum, pageSize)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, domain.Response{
			Result: false,
			Error:  err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, domain.Response{
		Result: true,
		Users:  data,
	})
}

func (h *RestHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domain.Response{
			Result: false,
			Error:  "invalid user ID",
		})
		return
	}

	data, err := h.repo.GetUserByID(id)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println(err)
		respondWithJSON(w, http.StatusNotFound, domain.Response{
			Result: false,
			Error:  "user not found",
		})
		return
	case err != nil:
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, domain.Response{
			Result: false,
			Error:  err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, domain.Response{
		Result: true,
		User:   &data,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
