package register

type Response struct {
	PublicId  string `json:"public_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
