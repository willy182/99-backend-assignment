package domains

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserRequest struct {
	Name string `json:"name"`
}

type UserServiceResponse struct {
	Result bool   `json:"result"`
	Users  []User `json:"users,omitempty"`
	User   *User  `json:"user,omitempty"`
	Error  string `json:"error,omitempty"`
}

type UserResponse struct {
	User  *User  `json:"user,omitempty"`
	Error string `json:"error,omitempty"`
}
