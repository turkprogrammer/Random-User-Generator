package user

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	// Добавим новые поля
	Gender     string
	Location   Location
	Login      Login
	DOB        DOB
	Registered Registered
	Phone      string
	Cell       string
	ID         ID
	Picture    Picture
	Nat        string
}

type Location struct {
	Street      Street
	City        string
	State       string
	Country     string
	Postcode    int
	Coordinates Coordinates
	Timezone    Timezone
}

type Street struct {
	Number int
	Name   string
}

type Coordinates struct {
	Latitude  string
	Longitude string
}

type Timezone struct {
	Offset      string
	Description string
}

type Login struct {
	UUID     string
	Username string
	Password string
	Salt     string
	MD5      string
	SHA1     string
	SHA256   string
}

type DOB struct {
	Date string
	Age  int
}

type Registered struct {
	Date string
	Age  int
}

type ID struct {
	Name  string
	Value string
}

type Picture struct {
	Large     string
	Medium    string
	Thumbnail string
}
