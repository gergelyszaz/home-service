package service

import (
	"context"
	"net/http"
	"time"

	api "github.com/gergelyszaz/home-service/gen"
)

type ApiService struct {
}

func NewApiService() api.DefaultApiServicer {
	return &ApiService{}
}

var wolTimeOut = 1 * time.Minute

var wake = false
var services = map[string]api.ServiceStatus{}

// HealthGet -
func (s *ApiService) HealthGet(ctx context.Context) (api.ImplResponse, error) {
	return api.Response(http.StatusOK, services), nil
}

// HealthServicePost -
func (s *ApiService) HealthServicePost(ctx context.Context, service string) (api.ImplResponse, error) {
	services[service] = api.ServiceStatus{
		Service:   service,
		UpdatedAt: time.Now().String(),
	}
	return api.Response(http.StatusOK, nil), nil
}

// WakeGet -
func (s *ApiService) WakeGet(ctx context.Context) (api.ImplResponse, error) {
	return api.Response(http.StatusOK, api.WakeStatus{WakeDevice: wake}), nil
}

// WakePost -
func (s *ApiService) WakePost(ctx context.Context) (api.ImplResponse, error) {
	if !wake {
		go disableWake()
	}
	wake = true
	return api.Response(http.StatusAccepted, nil), nil
}

func disableWake() {
	time.Sleep(wolTimeOut)
	wake = false
}
