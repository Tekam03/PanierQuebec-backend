package main

import (
	"fmt"
	"log"

	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/tekam03/panierquebec-backend/gen/products/v1/productsv1connect"
	dbgen "github.com/tekam03/panierquebec-backend/internal/db/gen"
	handlerExternalProduct "github.com/tekam03/panierquebec-backend/internal/handler/external_product"
	handlerMerchant "github.com/tekam03/panierquebec-backend/internal/handler/merchant"
	repoExternalProduct "github.com/tekam03/panierquebec-backend/internal/repository/external_product"
	repoMerchant "github.com/tekam03/panierquebec-backend/internal/repository/merchant"
	serviceExternalProduct "github.com/tekam03/panierquebec-backend/internal/service/external_product"
	serviceMerchant "github.com/tekam03/panierquebec-backend/internal/service/merchant"

	// repoMerchant "github.com/tekam03/panierquebec-backend/internal/repo/merchant"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/tekam03/panierquebec-backend/internal/config"
	"github.com/tekam03/panierquebec-backend/internal/db"
	"github.com/tekam03/panierquebec-backend/internal/migrate"
)

func main() {
	config.LoadEnv()

	// Connect to the database
	if err := db.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Run database migrations
	if err := migrate.RunMigrations("migrations"); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	// test the merchant repository
	// merchantRepo := repoMerchant.NewPostgresMerchantRepo()

	sqlcQueries := dbgen.New(db.Pool)

	merchantRepository := repoMerchant.NewRepo(sqlcQueries)
	externalProductRepository := repoExternalProduct.NewRepo(sqlcQueries)

	merchantService := serviceMerchant.NewService(merchantRepository)
	externalProductService := serviceExternalProduct.NewService(externalProductRepository)

	merchantHandler := handlerMerchant.NewMerchantHandler(merchantService)
	externalProductHandler := handlerExternalProduct.NewExternalProductHandler(externalProductService)

	mux := http.NewServeMux()
	path, handler := productsv1connect.NewMerchantServiceHandler(merchantHandler)
	mux.Handle(path, handler)
	path, handler = productsv1connect.NewExternalProductServiceHandler(externalProductHandler)
	mux.Handle(path, handler)
	reflector := grpcreflect.NewStaticReflector(
		productsv1connect.MerchantServiceName,
		productsv1connect.ExternalProductServiceName,
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
