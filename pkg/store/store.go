package store

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

type Store struct {
	db    map[string]int
	cache gcache.Cache
	ttl   time.Duration
}

func NewStore(db map[string]int, cache gcache.Cache, ttl time.Duration) Store {
	return Store{
		db:    db,
		cache: cache,
		ttl:   ttl,
	}
}

func (s *Store) Get(key string) int {
	fromCache, err := s.cache.Get(key)
	if err != nil {
		fmt.Println("Cache miss")
		fromDb, ok := s.db[key]
		if !ok {
			return -1
		}
		s.cache.SetWithExpire(key, fromDb, s.ttl)
		return fromDb
	}
	fmt.Println("Cache hit")
	return fromCache.(int)
}
