package main

import (
	"flag"
	"fmt"

	cfg "github.com/digisan/go-config"
	nt "github.com/digisan/gotk/net-tool"
	lk "github.com/digisan/logkit"
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
		if cfg.Val[bool]("IpRepl", i, "Enable") {
			var (
				toPubIP   = cfg.Val[bool]("IpRepl", i, "ToPub")
				toLocIP   = cfg.Val[bool]("IpRepl", i, "ToLoc")
				aimIP     = cfg.Val[string]("IpRepl", i, "ToIP")
				newPort   = cfg.Val[int]("IpRepl", i, "NewPort")
				onlyFirst = cfg.Val[bool]("IpRepl", i, "OnlyFirst")
				files     = cfg.ValArr[string]("IpRepl", i, "Files")
			)
			if newPort > 0 {
				lk.FailOnErr("%v", nt.ChangeLocalUrlPort(false, onlyFirst, -1, newPort, files...))
			}
			lk.FailOnErr("%v", nt.ChangeLocalhost(false, onlyFirst, toPubIP, toLocIP, aimIP, files...))
		}
	}

	for i := 0; i < nrSyb; i++ {
		if cfg.Val[bool]("SymbolRepl", i, "Enable") {
			var (
				onlyCmt = cfg.Val[bool]("SymbolRepl", i, "OnlyForCmt")
				files   = cfg.ValArr[string]("SymbolRepl", i, "Files")
			)
			lk.FailOnErr("%v", ReplaceSymbol(onlyCmt, files...))
		}
	}
}
