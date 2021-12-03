/**
 * @Author: gaoerpeng
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/12/2 11:42 上午
 */

package transport

import (
	endpts "ch6-discovery/endpoint"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

//MakeHttpHandler make http handler  use mux
func MakeHttpHandler(ctx context.Context, endpoints endpts.DiscoveryEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	//定义处理器
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	//say-hello 接口
	r.Methods("GET").Path("say-hello").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...,
	))

	//服务发现接口
	r.Methods("GET").Path("/discovery").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...,
	))

	//健康检查接口
	r.Methods("GET").Path("health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeJsonResponse,
		options...,
	))

	return r
}

func decodeDiscoveryRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	serverName := r.URL.Query().Get("serviceName")
	if serverName == "" {
		return nil, ErrorBadRequest
	}
	return endpts.DiscoveryRequest{ServiceName: serverName}, nil
}

func decodeHealthCheckRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return endpts.HealthRequest{}, nil
}

func encodeJsonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset:utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeSayHelloRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return endpts.SayHelloRequest{}, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json;charset:utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
