package business

import (
	"github.com/strwrd/driver-aood/src/storage"
)

type Case struct {
	Storage *storage.Storage
}

func NewCase(storage *storage.Storage) *Case {
	return &Case{Storage: storage}
}
