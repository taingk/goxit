package Helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	w.JSON(status, payload)
}
