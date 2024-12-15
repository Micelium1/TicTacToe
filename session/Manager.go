package Session

import (
	"math/rand"
	"strconv"
	"sync"
)

type Manager struct {
	Sessions map[string]*Session
	mu       sync.Mutex
}

var instance *Manager
var once sync.Once

func GetManager() *Manager {
	once.Do(func() {
		instance = &Manager{
			Sessions: make(map[string]*Session),
		}
	})
	return instance
}

func (m *Manager) GenerateUniqueID() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	for {
		id := generateRandomID()
		if _, exists := m.Sessions[id]; !exists {
			m.Sessions[id] = &Session{SessionId_: id}
			return id
		}
	}
}

func generateRandomID() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

func (m *Manager) DeleteSession(session Session) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.Sessions, session.SessionId_)

}
