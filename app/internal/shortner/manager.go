package shortner

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"url_shortner/app/internal/shortner/data"
)

type manager struct {
	logger *zap.SugaredLogger
}

func (m manager) redirectToLongUrl(req data.GetRequest, ctx *gin.Context) {
	return
}

func (m manager) getListOfUrlsForUser(req data.GetRequest, ctx *gin.Context) ([]data.ShortUrl, error) {
	return nil, nil
}

func newManager(logger *zap.SugaredLogger) *manager {
	return &manager{logger: logger}
}
