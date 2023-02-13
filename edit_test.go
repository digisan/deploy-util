package deployutil

import (
	"fmt"
	"testing"

	nt "github.com/digisan/gotk/net-tool"
)

func TestEdit(t *testing.T) {

	n, err := EditOriginIP("127.0.0.1", nt.LocalIP(), "http", -1, true, true, false, "backup", "bak", "./data/test.html")
	fmt.Println(n, err)
	if err != nil {
		return
	}

	n, err = EditSymbol(true, "", "bak", "./data/test.html")
	fmt.Println(n, err)
	if err != nil {
		return
	}
}
