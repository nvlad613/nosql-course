package http

import "github.com/labstack/echo/v5"

func (h *Handler) Health(c *echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}
