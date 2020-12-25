package server

import (
	"log"
	m "com/tienxe/lib/server/user_service/models"

	"github.com/gin-gonic/gin"
)

func(s* ApiGateWayServer) CreateUserCtrl(c *gin.Context){
	log.Printf("call api CreateUserCtrl")
	user := m.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
    	return
	}
	// s.authService.RegisterUser(&a.AuthUser{})
	log.Printf("bind data %v",user.Name)
	s.uDao.AddUser(&user)
	c.JSON(200, gin.H{"create ":"ok"})
	return
}