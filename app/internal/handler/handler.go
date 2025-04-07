package handler

import (
	"context"
	"net/http"

	external "gateway-service/external"
	proto "gateway-service/external/proto"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	AuthClient proto.UserServiceClient
}

func (h *Handler) CreateUser(ctx echo.Context) error {
	var req external.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	resp, err := h.AuthClient.CreateUser(
		context.Background(),
		&proto.CreateUserRequest{Name: req.Name},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "CreateUser failed: "+err.Error())
	}

	user := external.User{
		Id:   &resp.User.Id,
		Name: &resp.User.Name,
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(ctx echo.Context, id string) error {
	resp, err := h.AuthClient.GetUser(
		context.Background(),
		&proto.GetUserRequest{Id: id},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "GetUserByID failed: "+err.Error())
	}

	if resp.Users == nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	user := external.User{
		Id:   &resp.Users.Id,
		Name: &resp.Users.Name,
	}

	return ctx.JSON(http.StatusOK, user)
}
