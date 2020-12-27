module com/tienxe/lib

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/micro/micro/v3 v3.0.4
	github.com/mkideal/cli v0.2.3

    com/tienxe/lib/service/userService v0.0.0

)
replace (
    com/tienxe/lib/service/userService => ./service/userService
)