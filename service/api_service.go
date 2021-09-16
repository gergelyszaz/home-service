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

	servicesArr := []api.ServiceStatus{}
	for _, v := range services {
		servicesArr = append(servicesArr, v)
	}

	return api.Response(http.StatusOK, servicesArr), nil
}

// HealthServicePost -
func (s *ApiService) HealthServicePost(ctx context.Context, service string) (api.ImplResponse, error) {
	updatedAt, _ := time.Now().MarshalJSON()

	services[service] = api.ServiceStatus{
		Service:   service,
		UpdatedAt: string(updatedAt),
	}
	return api.Response(http.StatusOK, "OK"), nil
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
	return api.Response(http.StatusAccepted, "OK"), nil
}

func disableWake() {
	time.Sleep(wolTimeOut)
	wake = false
}
