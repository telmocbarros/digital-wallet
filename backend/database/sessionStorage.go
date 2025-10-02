package database

type Session struct {
	UserId    string
	CreatedAt int64
	ExpiresAt int64
}

var sessionStorage = make(map[string]Session)

func SaveSession(sessionId string, session Session) {
	sessionStorage[sessionId] = session
}

func GetSession(sessionId string) (Session, bool) {
	session, exists := sessionStorage[sessionId]
	return session, exists
}

func DeleteSession(sessionId string) {
	delete(sessionStorage, sessionId)
}
