package service

import (
	openapi "github.com/gergelyszaz/home-service/gen"
)

type ApiService struct {
	openapi.DefaultApiService
}

func NewApiService() openapi.DefaultApiServicer {
	return &ApiService{}
}
