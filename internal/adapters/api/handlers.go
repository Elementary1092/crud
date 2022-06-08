package api

import (
	"github.com/elem1092/crud/internal/adapters"
)

type handler struct {
	service Service
}

func NewHandler(service Service) adapters.Handler {
	return &handler{service}
}
