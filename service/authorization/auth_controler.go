package authorization

import (
	// "github.com/gin-gonic/gin"
)
type AuthCtrl struct {
	authService* AuthService
}
// func NewAuthCtrl(id string) *AuthCtrl{
// 	return &AuthCtrl{authService : &NewAuthenHandler(id)}
// } 
// func(auth * AuthCtrl) CreateUserCtrl(c *gin.Context){
// 	user := AuthUser{}
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//     	return
// 	}
// 	auth.authService.RegisterUser(&AuthUser{})
// }