package main

import(

"github.com/mkideal/cli"
sv "com/tienxe/lib/server"		

)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
	ResourceURL   string `cli:"*r,*resource" usage:"gateway resource url" dft:"$GATEWAY_RESOURCE_URL"`
}


func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		sv.NewApiGateWayServer(argv.ConnectionStr).Run()
		ctx.String("port=%d, x=%v, y=%v\n", argv.ConnectionStr, argv.ListenAddr, argv.ResourceURL)
		return nil
	})
	
}