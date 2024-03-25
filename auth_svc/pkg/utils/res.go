package utils

type Response struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type MethodResponse struct {
	Count      int32    `json:"count"`
	FirstNames []string `json:"firstnames"`
}
