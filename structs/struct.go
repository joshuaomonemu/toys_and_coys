package structs

import "cloud.google.com/go/firestore"

type Users struct {
	Firstname    string   `json:"firstname"`
	Lastname     string   `json:"lastname"`
	Email        string   `json:"email"`
	Phone_number string   `json:"phone_number"`
	Password     string   `json:"password"`
	D_o_b        string   `json:"d_o_b"`
	Username     string   `json:"username"`
	Event        []string `json:"event"`
}

type Events struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Content string `json:"content"`
	Text    string `json:"text"`
	Time    string `json:"time"`
	Likes   int64  `json:"likes"`
}

type EventComment struct {
	Ref     firestore.DocumentRef `json:"ref"`
	User    string                `json:"user"`
	Comment string                `json:"comment"`
	Time    string                `json:"time"`
}

type EventPayload struct {
	Succeeded bool   `json:"succeeded"`
	Data      Events `json:"data"`
	Errors    error  `json:"error"`
	Message   string `json:"message"`
}

type CommentPayload struct {
	Succeeded bool         `json:"succeeded"`
	Data      EventComment `json:"data"`
	Errors    error        `json:"error"`
	Message   string       `json:"message"`
}

type UserPayload struct {
	Succeeded bool   `json:"succeeded"`
	Data      Users  `json:"data"`
	Errors    error  `json:"errors"`
	Message   string `json:"message"`
}

type MinUserPayload struct {
	Succeeded bool   `json:"succeeded"`
	Errors    error  `json:"errors"`
	Message   string `json:"message"`
}
