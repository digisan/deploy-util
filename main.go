package main

import (
	"flag"
	"fmt"

	cfg "github.com/digisan/go-config"
	nt "github.com/digisan/gotk/net-tool"
)

func main() {

	fmt.Printf("Usage: \n\t-deploy: deploy json file [./deploy.json]\n\n")

	depPtr := flag.String("deploy", "./deploy.json", "deploy json file")
	flag.Parse()

	cfg.Init(true, *depPtr)
	cfg.Show()

	var (
		ipRepl     = cfg.Val[bool]("IpRepl.Enable")
		symbolRepl = cfg.Val[bool]("SymbolRepl.Enable")
	)

	if ipRepl {
		var (
			aimIP     = cfg.Val[string]("IpRepl.Loc2Ip")
			toPubIP   = cfg.Val[bool]("IpRepl.Loc2Pub")
			onlyfirst = cfg.Val[bool]("IpRepl.OnlyFirst")
			files     = cfg.ValArr[string]("IpRepl.Files")
		)
		nt.ChangeLocalIP(false, onlyfirst, toPubIP, aimIP, files...)
	}

	if symbolRepl {
		var (
			onlyCmt = cfg.Val[bool]("SymbolRepl.OnlyForCmt")
			files   = cfg.ValArr[string]("SymbolRepl.Files")
		)
		ReplaceSymbol(onlyCmt, files...)
	}
}