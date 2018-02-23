package service

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/omgwtflaserguns/matomat-server/config"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
	"sync"
	"time"
)

func WrapGrpcServer(grpcServer *grpc.Server, wg *sync.WaitGroup) {

	conf := config.GetConfig()
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	grpcHttpHandler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		if wrappedGrpc.IsGrpcWebRequest(req) {
			logger.Debugf("grpc: %s", req.RequestURI)
			wrappedGrpc.ServeHTTP(resp, req)
		} else {
			logger.Debugf("http: %s %s %s", req.Method, req.Proto, req.RequestURI)
		}

		// Fall back to other servers.
		http.DefaultServeMux.ServeHTTP(resp, req)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "https://matomat.de", "https://www.matomat.de"},
		AllowedHeaders:   []string{"X-Grpc-Web", "Content-Type"},
		AllowCredentials: true,
		Debug:            conf.DebugCors,
	})

	httpServer := &http.Server{
		Addr:           conf.GrpcWeb.Address,
		Handler:        c.Handler(grpcHttpHandler),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		IdleTimeout:    3 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// ErrorLog: logger, TODO maybe add error log?
	}

	wg.Add(1)
	go runHttpServer(httpServer, wg)

	logger.Debug("gRPC-web Server started at %s", conf.GrpcWeb.Address)
}

func runHttpServer(httpServer *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	err := httpServer.ListenAndServe()

	if err != nil {
		logger.Panicf("ListenAndServe failed: %s", err)
	}
}
