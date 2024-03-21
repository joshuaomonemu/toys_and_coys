package structs

import "time"

type Users struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Phone_number string `json:"phone_number"`
	Password     string `json:"password"`
	D_o_b        string `json:"d_o_b"`
	Username     string `json:"username"`
}

type Events struct {
	Title    string `json:"title"`
	User     string `json:"user"`
	Content  string `json:"content"`
	Text     string `json:"text"`
	Time     string `json:"time"`
	Likes    int64  `json:"likes"`
	Event_id string `json:"event_id"`
}

type EventComment struct {
	Commentid string    `json:"commentid"`
	User      string    `json:"user"`
	Comment   string    `json:"comment"`
	Time      time.Time `json:"time"`
}

type EventPayload struct {
	Success bool   `json:"success"`
	Data    Events `json:"data"`
	Errors  string `json:"error"`
	Message string `json:"message"`
}

type CommentPayload struct {
	Success bool           `json:"success"`
	Data    []EventComment `json:"data"`
	Errors  string         `json:"error"`
	Message string         `json:"message"`
}

type UserPayload struct {
	Success bool   `json:"success"`
	Data    Users  `json:"data"`
	Errors  string `json:"errors"`
	Message string `json:"message"`
}

type MinUserPayload struct {
	Success bool   `json:"success"`
	Errors  error  `json:"errors"`
	Message string `json:"message"`
}

type JWTPayload struct {
	UserId string `json:"userId"`
}
