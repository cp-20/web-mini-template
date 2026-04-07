package handler

import (
	"net/http"
	"strings"

	"github.com/cp-20/web-mini-template/server/internal/gen/openapi"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateTaskEcho(c echo.Context) error {
	var req createTaskRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json")
	}

	err := h.CreateTask(c.Request().Context(), &openapi.CreateTaskRequest{
		Title:    req.Title,
		MemberID: req.MemberID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "title and member_id are required") {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
