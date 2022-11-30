package v1

import (
	"github.com/ibrat-muslim/students-api/api/models"
	"github.com/ibrat-muslim/students-api/config"
	"github.com/ibrat-muslim/students-api/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
