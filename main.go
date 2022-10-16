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

	cfg.Init("main", true, *depPtr)
	cfg.Show()

	var (
		nrIP  = cfg.CntArr[any]("IpRepl")
		nrSyb = cfg.CntArr[any]("SymbolRepl")
	)

	for i := 0; i < nrIP; i++ {
		ipRepl := cfg.Val[bool]("IpRepl", i, "Enable")
		if ipRepl {
			var (
				toPubIP   = cfg.Val[bool]("IpRepl", i, "ToPub")
				toLocIP   = cfg.Val[bool]("IpRepl", i, "ToLoc")
				aimIP     = cfg.Val[string]("IpRepl", i, "ToIP")
				onlyfirst = cfg.Val[bool]("IpRepl", i, "OnlyFirst")
				files     = cfg.ValArr[string]("IpRepl", i, "Files")
			)
			nt.ChangeLocalhost(false, onlyfirst, toPubIP, toLocIP, aimIP, files...)
		}
	}

	for i := 0; i < nrSyb; i++ {
		symbolRepl := cfg.Val[bool]("SymbolRepl", i, "Enable")
		if symbolRepl {
			var (
				onlyCmt = cfg.Val[bool]("SymbolRepl", i, "OnlyForCmt")
				files   = cfg.ValArr[string]("SymbolRepl", i, "Files")
			)
			ReplaceSymbol(onlyCmt, files...)
		}
	}
}
