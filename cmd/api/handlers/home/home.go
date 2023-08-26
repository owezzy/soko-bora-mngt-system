package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for our home page.
func HomePageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Soko Bora"})
}
