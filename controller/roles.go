package controller

import (
	"admin_go/models"
	"admin_go/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CeeatRole(c *gin.Context) {
	var form models.SysRole
	if err := c.ShouldBind(&form); err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	fmt.Println("121", form)
	// if c.Validate(form) {
	// 	// do something
	// }
}
