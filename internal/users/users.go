package users

var (
	Users map[string]int
)

// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func IsValidUser(username, password string) (bool, error) {
	return true, nil
}

func UserNames() ([]string, error) {
	var usernames = []string{"jack", "john"}
	return usernames, nil
}

func Lookup(username string) User {

	u := User{
		Name:  username,
		Email: username + "@labstack.com",
	}

	return u
}
