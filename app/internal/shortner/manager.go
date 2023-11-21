package shortner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"url_shortner/app/internal/shortner/data"
)

type manager struct {
	logger  *zap.SugaredLogger
	storage data.Storage
}

func (m manager) redirectToLongUrl(req data.GetRequest, ctx *gin.Context) error {
	longUrl, err := m.storage.GetShortedUrl(req.User, req.ShortUrl)
	if err != nil {
		return err
	}
	ctx.Redirect(http.StatusPermanentRedirect, longUrl)
	return nil
}

func (m manager) getListOfUrlsForUser(req data.GetRequest, ctx *gin.Context) ([]data.ShortUrl, error) {
	return nil, nil
}

func (m manager) generateAndGetShortenedUrl(clientId string, longUrl string) data.ShortUrl {
	shortUrl := fmt.Sprintf("shotgun_%v", time.Now().UnixNano())
	m.storage.RegisterShortenedUrl(clientId, longUrl, shortUrl)
	return data.NewShortUrl(shortUrl)
}

func newManager(logger *zap.SugaredLogger) *manager {
	return &manager{logger: logger, storage: data.NewStorage()}
}
