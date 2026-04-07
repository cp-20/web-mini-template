package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetFeedEcho(c echo.Context) error {
	res, err := h.GetFeed(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
