package authhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout clears the HttpOnly JWT cookies (accessToken & refreshToken)
// by setting them with MaxAge = -1, instructing the browser to delete them immediately.
func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("accessToken", "", -1, "/", "localhost", false, true)
	c.SetCookie("refreshToken", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully.",
	})
}
