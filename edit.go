package deployutil

import (
	"os"
	"path/filepath"

	gio "github.com/digisan/gotk/io"
	nt "github.com/digisan/gotk/net-tool"
)

func EditOriginIP(old, new, onScheme string, onPort int, keepScheme, keepPort, only1st bool, bakFolder string, fPaths ...string) error {
	for _, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		if len(bakFolder) > 0 {
			bakDir := filepath.Join(filepath.Dir(fPath), bakFolder)
			gio.MustCreateDir(bakDir)
			if err = os.WriteFile(filepath.Join(bakDir, filepath.Base(fPath)), data, os.ModePerm); err != nil {
				return err
			}
		}

		if only1st {
			if err = nt.ModifyFile1stOriginIP(old, new, onScheme, onPort, keepScheme, keepPort, fPath); err != nil {
				return err
			}
		} else {
			if err = nt.ModifyFileOriginIP(old, new, onScheme, onPort, keepScheme, keepPort, fPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func EditSymbol(onlyInCmt bool, bakFolder string, fPaths ...string) error {
	for _, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return err
		}
		if len(bakFolder) > 0 {
			bakDir := filepath.Join(filepath.Dir(fPath), bakFolder)
			gio.MustCreateDir(bakDir)
			if err = os.WriteFile(filepath.Join(bakDir, filepath.Base(fPath)), data, os.ModePerm); err != nil {
				return err
			}
		}

		if err = ReplaceSymbol(onlyInCmt, filepath.Ext(fPath), fPath); err != nil {
			return err
		}
	}
	return nil
}
