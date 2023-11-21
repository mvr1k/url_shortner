package data

import "errors"

type Storage map[string]map[string]string

func NewStorage() Storage {
	return make(map[string]map[string]string)
}

func (store Storage) RegisterShortenedUrl(clientId string, longUrl string, shortUrl string) {
	var clientLevelStoreage = store[clientId]
	if clientLevelStoreage == nil {
		store[clientId] = make(map[string]string)
		clientLevelStoreage = store[clientId]
	}
	clientLevelStoreage[shortUrl] = longUrl
}

const (
	ErrorClientNotFound            = "could not find the client id in mappings"
	ErrorMappingNotFoundForLongUrl = "could not find the mapping for the short url"
)

func (store Storage) GetShortedUrl(clientId string, shortUrl string) (res string, err error) {
	clientLevelStorage := store[clientId]
	res = ""
	if clientLevelStorage != nil {
		found := false
		res, found = clientLevelStorage[shortUrl]
		if !found {
			err = errors.New(ErrorMappingNotFoundForLongUrl)
		}
	} else {
		err = errors.New(ErrorClientNotFound)
	}
	return
}

func (store Storage) GetAllUrlsByClientId(clientId string) (res map[string]string, err error) {
	res = store[clientId]
	if res == nil {
		err = errors.New(ErrorClientNotFound)
	}
	return
}
