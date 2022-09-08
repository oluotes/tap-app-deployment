package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Status  string `json:"status"`
	ApiName string `json:"api-name"`
}

var healthStatus = Status{
	Status:  "Up",
	ApiName: "Address Book API",
}

// HealthCheck godoc
// @Summary Show the status of address nook server.
// @Description get the status of server.
// @Tags ServerStatus
// @Accept */*
// @Produce json
// @Success 200 {object} Status
// @Router /healthz [get]
func GetStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, healthStatus)
}
