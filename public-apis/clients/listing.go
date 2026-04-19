package clients

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"public-apis/domains"
	"public-apis/utils"
)

type IListingService interface {
	CreateListing(param domains.ListingRequest) (domains.ListingServiceResponse, error)
	GetAllListings(pageNum, pageSize string) (domains.ListingServiceResponse, error)
	GetListingByID(id string) (domains.ListingServiceResponse, error)
}

type listingService struct {
	host    string
	httpReq *utils.HTTPClient
}

func NewListingService(baseURL string, req *utils.HTTPClient) IListingService {
	return &listingService{host: baseURL, httpReq: req}
}

// CreateListing implements IListingService.
func (u *listingService) CreateListing(param domains.ListingRequest) (domains.ListingServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var res domains.ListingServiceResponse

	form := url.Values{}
	form.Add("user_id", strconv.Itoa(param.UserID))
	form.Add("listing_type", param.ListingType)
	form.Add("price", strconv.Itoa(param.Price))

	err := u.httpReq.Call(http.MethodPost, u.host, strings.NewReader(form.Encode()), &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

// GetAllListings implements IListingService.
func (u *listingService) GetAllListings(pageNum string, pageSize string) (domains.ListingServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fullURL := fmt.Sprintf("%s?page_num=%s&page_size=%s", u.host, pageNum, pageSize)

	var res domains.ListingServiceResponse

	err := u.httpReq.Call(http.MethodGet, fullURL, nil, &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

// GetListingByID implements IListingService.
func (u *listingService) GetListingByID(id string) (domains.ListingServiceResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fullURL := fmt.Sprintf("%s/%s", u.host, id)

	var res domains.ListingServiceResponse

	err := u.httpReq.Call(http.MethodGet, fullURL, nil, &res, headers)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}
