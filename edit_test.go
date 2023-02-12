package deployutil

import (
	"fmt"
	"testing"

	nt "github.com/digisan/gotk/net-tool"
)

func TestEdit(t *testing.T) {
	if err := EditOriginIP("127.0.0.1", nt.LocalIP(), "http", -1, true, true, false, "backup", "./data/main.go"); err != nil {
		fmt.Println(err)
		return
	}
	if err := EditSymbol(true, "", "./data/main.go"); err != nil {
		fmt.Println(err)
		return
	}
}
