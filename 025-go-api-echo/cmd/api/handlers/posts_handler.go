package handlers

import(
	"net/http"

	"github.com/FuaAlfu/Go-Projects/025-go-api-echo/cmd/api/service"
)

func PoastIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway,"Unable to procces data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.json(http.StatusOK, res)
}

func PoastSingleHandler(c echo.Context) error {
	id := c.Param("")
}