package usecases

import "github.com/gabrielerzinger/urlShrtnr/repositories"

type EntryUsecases interface {
}

type entryUsecases struct {
	repo repositories.Repository
}


// NewUsecase ctor
func NewUsecase(newRepo repositories.Repository) *entryUsecases {
	return &entryUsecases{repo: newRepo}
}

// todo: move those cosntants ${base} and ${mod} elsewhere.
func (e entryUsecases) hashString(s string) (hash uint64) {
	var idx uint64 = 1
	hash = 0
	var base uint64 = 62
	var mod uint64 = 1e9+9

	for i := 0; i < len(s) ; i++  {
		hash = (hash + ( uint64(s[i]) - uint64('a') + uint64(1) * idx) % mod)
		idx = (idx * base) % mod
	}
	return

}

// ShortenUrl takes a full url and return it's short version
func (e entryUsecases) ShortenUrl(url string) (ret string) {
	var s string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hash := e.hashString(url)
	for hash > 0  {
		ret = ret + string(s[hash%62])
		hash = hash / 62
	}
	return
}
