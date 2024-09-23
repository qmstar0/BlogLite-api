package auth

type User struct {
	ID   uint32
	Type string
	Name string
}

var Admin = User{
	ID:   1,
	Type: "admin",
	Name: "qmstar",
}
