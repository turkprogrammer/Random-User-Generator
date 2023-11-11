package user

type RandomUserResponse struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email      string     `json:"email"`
		Gender     *string    `json:"gender"`
		Location   Location   `json:"location"`
		Login      Login      `json:"login"`
		DOB        DOB        `json:"dob"`
		Registered Registered `json:"registered"`
		Phone      string     `json:"phone"`
		Cell       string     `json:"cell"`
		ID         ID         `json:"id"`
		Picture    Picture    `json:"picture"`
		Nat        string     `json:"nat"`
	} `json:"results"`
	// Add other fields as needed
}
