package types

type CreateClientRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PhoneNO   int    `json:"phonenumber"`
}

type Client struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	PhoneNO   int    `json:"phonenumber"`
}

func NewClient(firstName, lastName, email string, phoneNo int) *Client {
	return &Client{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		PhoneNO:   phoneNo,
	}
}
