package database

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Data []User = []User{
	{Id: "b18b851a-c8c4-4957-b68a-14362a1810c6", Email: "john@example.com", Password: "password123"},
	{Id: "b5ed9407-681b-4dbb-b2d3-997803e8bbfc", Email: "jane@example.com", Password: "securepass"},
}

func GetUserById(id string) (User, bool) {
	for _, user := range Data {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}
