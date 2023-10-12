package token_manager

import "sync"

var tokens = map[string]bool{}
var rwLocker = sync.RWMutex{}

func FindToken(token string) bool {
	rwLocker.RLock()
	defer rwLocker.RUnlock()

	return tokens[token]
}

func AddToken(token string) {
	rwLocker.Lock()
	defer rwLocker.Unlock()

	tokens[token] = true
}

var tokensSyncMap = sync.Map{}

func FindTokenFromSyncMap(token string) bool {
	_, ok := tokensSyncMap.Load(token)
	return ok
}

func AddTokenToSyncMap(token string) {
	tokensSyncMap.Store(token, true)
}
