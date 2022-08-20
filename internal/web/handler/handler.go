package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebin-pavus/Assignment/internal/model"
)

// postCompute returs A/B and B/A
func PostCompute(c *gin.Context) {
	var newInput model.Input
	var newOutput model.Output

	// Call BindJSON to bind the received JSON to
	// newInput.
	if err := c.BindJSON(&newInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H(gin.H{"message": "Error binding json"}))
		return
	}

	if newInput.A == 0 || newInput.B == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H(gin.H{"message": "Division by zero. Fatal!"}))
		return
	}

	newOutput.First = newInput.A / newInput.B
	newOutput.Second = newInput.B / newInput.A

	c.IndentedJSON(http.StatusOK, newOutput)
}
