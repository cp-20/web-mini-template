package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) InitializeEcho(c echo.Context) error {
	if err := h.Initialize(c.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
