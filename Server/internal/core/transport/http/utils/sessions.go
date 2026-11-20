package core_http_utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"
)

var (
	sessionMu sync.Mutex
	sessions  = map[string]int{}   
)

func newSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func CreateSession(w http.ResponseWriter, userID int) {
	sid := newSessionID()

	sessionMu.Lock()
	sessions[sid] = userID
	sessionMu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sid,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 7,
	})
}

func GetUserIDFromSession(r *http.Request) (int, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0, false
	}
	sessionMu.Lock()
	defer sessionMu.Unlock()
	id, ok := sessions[cookie.Value]
	return id, ok
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("session_id"); err == nil {
		sessionMu.Lock()
		delete(sessions, cookie.Value)
		sessionMu.Unlock()
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
	})
}