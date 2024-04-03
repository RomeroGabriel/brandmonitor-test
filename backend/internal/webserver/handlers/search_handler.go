package handlers

import (
	"net/http"

	"github.com/RomeroGabriel/brandmonitor-challange/backend/internal/webserver/dto"
	"github.com/RomeroGabriel/brandmonitor-challange/backend/pkg/search"
	"github.com/labstack/echo/v4"
)

type SearchHandler struct{}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{}
}

func (h *SearchHandler) Search(c echo.Context) error {
	var data dto.SearchDTO

	if err := c.Bind(&data); err != nil {
		return err
	}

	paramEntity := data.ParseToSearchParameters()
	res, err := search.Search(paramEntity)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
