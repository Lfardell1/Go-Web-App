package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// SessionProvider defines the methods that a session provider must implement.
type SessionProvider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifetime int64)
}

// Session represents a user session.
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// Manager manages the session middleware.
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    SessionProvider
	maxLifetime int64
}

var providers = make(map[string]SessionProvider)

// NewManager creates a new session manager.
func NewManager(providerName, cookieName string, maxLifetime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q (forgot to register?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifetime: maxLifetime}, nil
}

// RegisterProvider registers a session provider.
func RegisterProvider(name string, provider SessionProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

// ... (Other methods: sessionId, SessionStart, SessionDestroy, etc.)

// SessionStart starts a new session or retrieves an existing session.
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) Session {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ := manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifetime)}
		http.SetCookie(w, &cookie)
		return session
	}

	sid, _ := url.QueryUnescape(cookie.Value)
	session, _ := manager.provider.SessionRead(sid)
	return session
}

// SessionDestroy destroys a session.
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.SessionDestroy(cookie.Value)
	expiration := time.Now()
	cookie = &http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(w, cookie)
}

// sessionId generates a session ID.
func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionInit initializes a new session with a given session ID.
func (manager *Manager) SessionInit(sid string) (Session, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	newSession, err := manager.provider.SessionInit(sid)
	if err != nil {
		return nil, err
	}
	return newSession, nil
}

// SessionRead reads an existing session with a given session ID.
func (manager *Manager) SessionRead(sid string) (Session, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	session, err := manager.provider.SessionRead(sid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// SessionGC performs session garbage collection.
func (manager *Manager) SessionGC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.SessionGC(manager.maxLifetime)
	time.AfterFunc(time.Duration(manager.maxLifetime), func() { manager.SessionGC() })
}
