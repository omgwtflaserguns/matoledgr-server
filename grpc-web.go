package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func wrapGrpcServer(grpcServer *grpc.Server) {

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithAllowedRequestHeaders([]string{"*"}))

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
		// TODO harden this
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-Grpc-Web", "Content-Type"},
		Debug:          config.debugCors,
	})

	httpServer := &http.Server{
		Addr:           config.grpcWeb.address,
		Handler:        c.Handler(grpcHttpHandler),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		IdleTimeout:    3 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// ErrorLog: logger, TODO maybe add error log?
	}

	wg.Add(1)
	go runHttpServer(httpServer)

	logger.Debug("gRPC-web Server started at %s", config.grpcWeb.address)
}

func runHttpServer(httpServer *http.Server) {
	defer wg.Done()

	err := httpServer.ListenAndServe()

	if err != nil {
		logger.Panicf("ListenAndServe failed: %s", err)
	}
}
