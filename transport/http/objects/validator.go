package objects

import (
	svc "sample-service"

	"github.com/gin-gonic/gin"
)

type ObjectValidator struct {
	ID uint64 `binding:"required" json:"id"`
	IP string `binding:"required" json:"ip"`
}

func Bind(c *gin.Context) (*svc.Object, error) {
	var json ObjectValidator
	if err := c.ShouldBindJSON(&json); err != nil {
		return nil, err
	}

	author := &svc.Object{
		ID: json.ID,
		IP: json.IP,
	}

	return author, nil
}
