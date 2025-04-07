package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	external "gateway-service/external"
	proto "gateway-service/external/proto"
	handler "gateway-service/internal/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	httpPort := 8080

	grpcTarget := "host.docker.internal:50051"

	log.Printf("[INFO] Connecting to gRPC service at %s", grpcTarget)

	conn, err := grpc.NewClient(
		grpcTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to gRPC: %v", err)
	}

	log.Printf("[INFO] Successfully connected to gRPC service at %s", grpcTarget)

	defer func(conn *grpc.ClientConn) {
		log.Println("[INFO] Closing gRPC connection...")
		if err := conn.Close(); err != nil {
			log.Fatalf("[ERROR] Failed to close gRPC connection: %v", err)
		}
		log.Println("[INFO] gRPC connection closed.")
	}(conn)

	authClient := proto.NewUserServiceClient(conn)
	log.Println("[INFO] AuthServiceClient initialized.")

	h := &handler.Handler{
		AuthClient: authClient,
	}
	log.Println("[INFO] Handler is set up with AuthServiceClient.")

	external.RegisterHandlers(e, h)
	log.Println("[INFO] Routes are registered via openapi.gen.go.")

	addr := fmt.Sprintf(":%d", httpPort)
	log.Printf("[INFO] Starting HTTP server on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatalf("[ERROR] Echo server failed: %v", err)
	}
}
