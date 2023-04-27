package entities

type Video struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Person struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Age       int    `json:"age" validate:"gte=0,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}
