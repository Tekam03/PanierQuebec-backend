package main

import (
	"fmt"
	"log"

	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/tekam03/panierquebec-backend/gen/stores/v1/storesv1connect"
	handlerMerchant "github.com/tekam03/panierquebec-backend/internal/handler/merchant"
	serviceMerchant "github.com/tekam03/panierquebec-backend/internal/service/merchant"
	repoMerchant "github.com/tekam03/panierquebec-backend/internal/repo/merchant"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/tekam03/panierquebec-backend/internal/config"
	"github.com/tekam03/panierquebec-backend/internal/db"
)

func main() {
	config.LoadEnv()

	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// test the merchant repository
	merchantRepo := repoMerchant.NewPostgresMerchantRepo()
	merchantService := serviceMerchant.NewService(merchantRepo)
	merchantHandler := handlerMerchant.NewMerchantHandler(merchantService)


	mux := http.NewServeMux()
	path, handler := storesv1connect.NewMerchantServiceHandler(merchantHandler)
	mux.Handle(path, handler)
	reflector := grpcreflect.NewStaticReflector(
		storesv1connect.MerchantServiceName,
	)
	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	// If you don't need to support HTTP/2 without TLS (h2c), you can drop
	// x/net/http2 and use http.ListenAndServeTLS instead.

	fmt.Println("Starting server on localhost:8080")
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
		// Don't forget timeouts!
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

}
