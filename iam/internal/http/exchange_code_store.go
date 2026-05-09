package http

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/owezzy/soko-bora-mngt-system/iam/internal/application"
)

type ExchangeCodeStore struct {
	mu    sync.Mutex
	codes map[string]exchangeCodeEntry
}

type exchangeCodeEntry struct {
	session   application.AuthSession
	expiresAt time.Time
}

func NewExchangeCodeStore() *ExchangeCodeStore {
	return &ExchangeCodeStore{codes: make(map[string]exchangeCodeEntry)}
}

func (s *ExchangeCodeStore) Issue(session application.AuthSession, ttl time.Duration) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pruneLocked(time.Now())
	code := randomCode()
	s.codes[code] = exchangeCodeEntry{
		session:   session,
		expiresAt: time.Now().Add(ttl),
	}

	return code
}

func (s *ExchangeCodeStore) Consume(code string) (application.AuthSession, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, ok := s.codes[code]
	if !ok {
		return application.AuthSession{}, false
	}
	delete(s.codes, code)
	if time.Now().After(entry.expiresAt) {
		return application.AuthSession{}, false
	}

	return entry.session, true
}

func (s *ExchangeCodeStore) pruneLocked(now time.Time) {
	for code, entry := range s.codes {
		if now.After(entry.expiresAt) {
			delete(s.codes, code)
		}
	}
}

func randomCode() string {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(buf)
}
