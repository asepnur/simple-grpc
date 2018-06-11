package handler

// User struct :: save value
type User struct {
	UserID     string `json:"id"`
	UserEmail  string `json:"email"`
	FullName   string `json:"FullName"`
	MSISDN     string `json:"msisdn"`
	CreateTime string `json:"create_time"`
}

// UserParams struct
type UserParams struct {
	Q string
}
