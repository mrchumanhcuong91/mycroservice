package main

import(

"github.com/mkideal/cli"
sv "com/tienxe/lib/server"		

)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
}


func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		sv.NewApiGateWayServer(argv.ConnectionStr).Run(argv.ListenAddr +":8181")
		return nil
	})
	
}