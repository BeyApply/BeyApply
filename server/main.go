package main

import (
	"BeyApply/latexUtils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct for pdf text from frontend
type PdfTextRequest struct {
	Text string `json:"text"`
}

func main() {
	fmt.Println("wssg gang")
	latexUtils.PrintHi() // Call the function properly

	r := gin.Default()

	r.POST("/getResumeText", func(c *gin.Context) {
		var request PdfTextRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Process the pdfText (request.Text)
		// For example, print it or save to a database
		println("Received PDF text:", request.Text)

		// Respond to frontend
		c.JSON(http.StatusOK, gin.H{
			"message": "PDF text received successfully",
		})
	})

	r.Run(":8080") // Run the server on port 8080
}
