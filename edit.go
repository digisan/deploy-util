package deployutil

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/digisan/go-generics/v2"
	gio "github.com/digisan/gotk/io"
	nt "github.com/digisan/gotk/net-tool"
)

// return (err-index, error)
func EditOriginIP(old, new, onScheme string, onPort int, keepScheme, keepPort, only1st bool, bakFolder, ext string, fPaths ...string) (int, error) {
	for i, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return i, err
		}
		if len(bakFolder) > 0 {
			bakDir := filepath.Join(filepath.Dir(fPath), bakFolder)
			gio.MustCreateDir(bakDir)
			ext = strings.TrimPrefix(strings.TrimSpace(ext), ".")
			ext = IF(len(ext) > 0, "."+ext, ext)
			if err = os.WriteFile(filepath.Join(bakDir, filepath.Base(fPath)+ext), data, os.ModePerm); err != nil {
				return i, err
			}
		}

		if only1st {
			if err = nt.ModifyFile1stOriginIP(old, new, onScheme, onPort, keepScheme, keepPort, fPath); err != nil {
				return i, err
			}
		} else {
			if err = nt.ModifyFileOriginIP(old, new, onScheme, onPort, keepScheme, keepPort, fPath); err != nil {
				return i, err
			}
		}
	}
	return -1, nil
}

// return (err-index, error)
func EditSymbol(onlyInCmt bool, bakFolder, ext string, fPaths ...string) (int, error) {
	for i, fPath := range fPaths {
		data, err := os.ReadFile(fPath)
		if err != nil {
			return i, err
		}
		if len(bakFolder) > 0 {
			bakDir := filepath.Join(filepath.Dir(fPath), bakFolder)
			gio.MustCreateDir(bakDir)
			ext = strings.TrimPrefix(strings.TrimSpace(ext), ".")
			ext = IF(len(ext) > 0, "."+ext, ext)
			if err = os.WriteFile(filepath.Join(bakDir, filepath.Base(fPath)+ext), data, os.ModePerm); err != nil {
				return i, err
			}
		}

		if err = ReplaceSymbol(onlyInCmt, filepath.Ext(fPath), fPath); err != nil {
			return i, err
		}
	}
	return -1, nil
}
