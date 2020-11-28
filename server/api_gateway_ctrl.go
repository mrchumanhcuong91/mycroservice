package server

import (
	"log"
	a "com/tienxe/lib/service/authorization"
	"github.com/gin-gonic/gin"
)

func(s* ApiGateWayServer) CreateUserCtrl(c *gin.Context){
	log.Printf("call api CreateUserCtrl")
	user := a.AuthUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
    	return
	}
	// s.authService.RegisterUser(&a.AuthUser{})
	log.Printf("bind data %v",user.UserName)
	c.JSON(200, gin.H{"create ":"ok"})
	return
}