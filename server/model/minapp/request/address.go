package request

type CreateAddress struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Detail   string `json:"detail"`
	PostCode string `json:"postCode"`
}

type UpdateAddress struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Detail   string `json:"detail"`
	PostCode string `json:"postCode"`
}

type DeleteAddress struct {
	ID uint `json:"id"`
}
