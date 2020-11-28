package server
import ("github.com/gin-gonic/gin"
		a "com/tienxe/lib/service/authorization"
		)
//declare struct apigateway
type ApiGateWayServer struct{
	//web service
	* gin.Engine
	//authorization service
	authService a.AuthService

}
//constructor
func NewApiGateWayServer(mysql string) (api *ApiGateWayServer){
	api = &ApiGateWayServer{
		Engine: gin.Default(),
		authService: a.NewAuthenHandler(mysql)}
	api.createAuthenApi()
	return
}
func (api *ApiGateWayServer) createAuthenApi(){
	api.Group("/gateway/api")
	api.POST("/register",api.CreateUserCtrl)
}