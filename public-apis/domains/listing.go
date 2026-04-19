package domains

type Listing struct {
	ID          int64  `json:"id"`
	UserID      int    `json:"user_id,omitempty"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	UserData    *User  `json:"user,omitempty"`
}

type ListingRequest struct {
	UserID      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
}

type ListingServiceResponse struct {
	Result   *bool     `json:"result,omitempty"`
	Listings []Listing `json:"listings,omitempty"`
	Listing  *Listing  `json:"listing,omitempty"`
	Error    string    `json:"error,omitempty"`
}
