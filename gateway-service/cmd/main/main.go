package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"

	external "gateway-service/external"
	proto "gateway-service/external/proto"
	handler "gateway-service/internal/handler"
)

func main() {
	e := echo.New()

	conn, err := grpc.Dial("auth-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to auth-service: %v", err)
	}
	defer conn.Close()

	authClient := proto.NewAuthServiceClient(conn)

	h := &handler.Handler{
		AuthClient: authClient,
	}

	external.RegisterHandlers(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
