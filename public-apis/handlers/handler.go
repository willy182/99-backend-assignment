package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"public-apis/clients"
	"public-apis/domains"
	"public-apis/utils"
)

type RestHandler struct {
	userService    clients.IUserService
	listingService clients.IListingService
}

func NewRestHandler(userService clients.IUserService, listingService clients.IListingService) *RestHandler {
	return &RestHandler{userService: userService, listingService: listingService}
}

func (h *RestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var payload domains.UserRequest
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domains.UserResponse{
			Error: err.Error(),
		})
		return
	}

	if payload.Name == "" {
		errMsg := "name is mandatory"
		log.Println(errMsg)
		respondWithJSON(w, http.StatusBadRequest, domains.UserResponse{
			Error: errMsg,
		})
		return
	}

	data, err := h.userService.CreateUser(payload)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, domains.UserResponse{
		User: data.User,
	})
}

func (h *RestHandler) CreateOrGetListings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAllListings(w, r)
	case http.MethodPost:
		h.CreateListing(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RestHandler) CreateListing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithJSON(w, http.StatusMethodNotAllowed, nil)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var payload domains.ListingRequest
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Error: err.Error(),
		})
		return
	}

	if payload.UserID == 0 {
		errMsg := "user_id is mandatory"
		log.Println(errMsg)
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Error: errMsg,
		})
		return
	}

	if payload.ListingType == "" {
		errMsg := "listing_type is mandatory"
		log.Println(errMsg)
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Error: errMsg,
		})
		return
	}

	if payload.Price == 0 {
		errMsg := "price cannot be 0"
		log.Println(errMsg)
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Error: errMsg,
		})
		return
	}

	data, err := h.listingService.CreateListing(payload)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, domains.ListingServiceResponse{
		Listing: data.Listing,
	})
}

func (h *RestHandler) GetAllListings(w http.ResponseWriter, r *http.Request) {
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
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Result: utils.ToBoolPtr(false),
			Error:  "invalid parameter page_num",
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		log.Println(err)
		respondWithJSON(w, http.StatusBadRequest, domains.ListingServiceResponse{
			Result: utils.ToBoolPtr(false),
			Error:  "invalid parameter page_size",
		})
		return
	}

	resListing, err := h.listingService.GetAllListings(pageNumStr, pageSizeStr)
	if err != nil {
		log.Println(err)
		respondWithJSON(w, http.StatusInternalServerError, err)
		return
	}

	userIds := make([]int, 0)
	resUser, err := h.userService.GetUserByIds([]int)
	if err != nil {
		log.Println(err)
	}

	dataUser := map[string]User
	for _, v := range resUser {
		dataUser[v.id] = v
	}

	data := make([]domains.Listing, 0)
	for _, val := range resListing.Listings {
		data = append(data, domains.Listing{
			ID:          val.ID,
			ListingType: val.ListingType,
			Price:       val.Price,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
			UserData:    dataUser[val.UserID],
		})
	}

	respondWithJSON(w, http.StatusCreated, domains.ListingServiceResponse{
		Result:   utils.ToBoolPtr(true),
		Listings: data,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
