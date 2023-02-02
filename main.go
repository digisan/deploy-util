package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/digisan/go-config"
	gio "github.com/digisan/gotk/io"
	nt "github.com/digisan/gotk/net-tool"
	lk "github.com/digisan/logkit"
)

func main() {

	cfgPtr := flag.String("config", "./deploy.json", "deploy config json file")
	backupPtr := flag.Bool("backup", true, "backup the original file?")
	filePtr := flag.String("file", "", "file path")
	pubPtr := flag.Bool("pub", false, "replace to public IP?")
	locPtr := flag.Bool("loc", false, "replace to local net IP?")
	firstPtr := flag.Bool("first", true, "only replace the first occurrence?")
	cmtPtr := flag.Bool("cmt", true, "only replace the occurrence in comment?")

	flag.Parse()

	bakFolder := "original"

	///////////////////////////////////////////////////////
	// cli is the first option

	if fPath := *filePtr; len(fPath) > 0 {
		if *backupPtr {
			data, err := os.ReadFile(fPath)
			lk.FailOnErr("%v", err)
			backupDir := filepath.Join(filepath.Dir(fPath), bakFolder)
			gio.MustCreateDir(backupDir)
			lk.FailOnErr("%v", os.WriteFile(filepath.Join(backupDir, filepath.Base(fPath)), data, os.ModePerm))
		}
		lk.FailOnErr("%v", nt.ChangeLocalhost(false, *firstPtr, *pubPtr, *locPtr, "127.0.0.1", fPath))
		lk.FailOnErr("%v", ReplaceSymbol(*cmtPtr, fPath))
		return
	}

	///////////////////////////////////////////////////////
	// config is the 2nd option

	cfgPath := *cfgPtr
	lk.FailOnErrWhen(len(cfgPath) == 0, "%v", fmt.Errorf("deploy config is missing"))

	cfg.Init("main", true, cfgPath)
	cfg.Show()

	var (
		nrIP   = cfg.CntObjects("IpRepl")
		nrSyb  = cfg.CntObjects("SymbolRepl")
		backup = cfg.Val[bool]("Backup")
	)

	if backup {

		for i := 0; i < nrIP; i++ {
			if cfg.Val[bool]("IpRepl", i, "Enable") {
				for _, fPath := range cfg.ValArr[string]("IpRepl", i, "Files") {
					data, err := os.ReadFile(fPath)
					lk.FailOnErr("%v", err)
					backupDir := filepath.Join(filepath.Dir(fPath), bakFolder)
					gio.MustCreateDir(backupDir)
					lk.FailOnErr("%v", os.WriteFile(filepath.Join(backupDir, filepath.Base(fPath)), data, os.ModePerm))
				}
			}
		}

		for i := 0; i < nrSyb; i++ {
			if cfg.Val[bool]("SymbolRepl", i, "Enable") {
				for _, fPath := range cfg.ValArr[string]("SymbolRepl", i, "Files") {
					data, err := os.ReadFile(fPath)
					lk.FailOnErr("%v", err)
					backupDir := filepath.Join(filepath.Dir(fPath), bakFolder)
					gio.MustCreateDir(backupDir)
					lk.FailOnErr("%v", os.WriteFile(filepath.Join(backupDir, filepath.Base(fPath)), data, os.ModePerm))
				}
			}
		}

	}

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
