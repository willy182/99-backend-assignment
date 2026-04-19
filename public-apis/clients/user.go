package clients

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"public-apis/domains"
	"public-apis/utils"
)

type IUserService interface {
	CreateUser(param domains.UserRequest) (domains.UserServiceResponse, error)
	GetAllUsers(pageNum, pageSize string) (domains.UserServiceResponse, error)
	GetUserByID(id string) (domains.UserServiceResponse, error)
}

type userService struct {
	host    string
	httpReq *utils.HTTPClient
}

func NewUserService(baseURL string, req *utils.HTTPClient) IUserService {
	return &userService{host: baseURL, httpReq: req}
}

// CreateUser implements IUserService.
func (u *userService) CreateUser(param domains.UserRequest) (domains.UserServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var res domains.UserServiceResponse

	form := url.Values{}
	form.Add("name", param.Name)

	err := u.httpReq.Call(http.MethodPost, u.host, strings.NewReader(form.Encode()), &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

// GetAllUsers implements IUserService.
func (u *userService) GetAllUsers(pageNum string, pageSize string) (domains.UserServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fullURL := fmt.Sprintf("%s?page_num=%s&page_size=%s", u.host, pageNum, pageSize)

	var res domains.UserServiceResponse

	err := u.httpReq.Call(http.MethodGet, fullURL, nil, &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

// GetUserByID implements IUserService.
func (u *userService) GetUserByID(id string) (domains.UserServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fullURL := fmt.Sprintf("%s/%s", u.host, id)

	var res domains.UserServiceResponse

	err := u.httpReq.Call(http.MethodGet, fullURL, nil, &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}
