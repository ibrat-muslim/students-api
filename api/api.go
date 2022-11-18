package api

import (
	"github.com/ibrat-muslim/students-api/config"
	"github.com/ibrat-muslim/students-api/storage"
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}
