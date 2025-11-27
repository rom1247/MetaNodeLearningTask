package cache

import (
	"time"

	"github.com/bluele/gcache"
)

var TokenCache gcache.Cache

func InitTokenCache(defaultExpire time.Duration, size int) {
	TokenCache = gcache.New(size).
		LRU().
		Expiration(defaultExpire).
		Build()
}

func StoreToken(token string, userID uint) {
	_ = TokenCache.Set(token, userID)
}

func GetToken(token string) (uint, bool) {
	val, err := TokenCache.Get(token)
	if err != nil {
		return 0, false
	}
	return val.(uint), true
}

func DeleteToken(token string) {
	TokenCache.Remove(token)
}
