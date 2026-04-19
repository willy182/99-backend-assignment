package domain

type Response struct {
	Result bool   `json:"result"`
	Users  []User `json:"users,omitempty"`
	User   *User  `json:"user,omitempty"`
	Error  string `json:"error,omitempty"`
}
