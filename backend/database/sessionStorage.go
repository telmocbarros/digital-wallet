package database

type Session struct {
	UserId    string
	CreatedAt int64
	ExpiresAt int64
}

var SessionStorage = make(map[string]Session)
