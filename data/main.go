package main

import (
	"fmt"

	lk "github.com/digisan/logkit"
)

var (
	fHttp2 = false //
	port   = 1323  // note: keep same as below @host
)

func init() {
	lk.Log("starting...main")
	lk.WarnDetail(false)
}

// @title National Education Data Dictionary API
// @version 1.0
// @description This is national education data dictionary backend-api server. Updated@ 2022-09-15T09:29:03+10:00
// @termsOfService
// @contact.name API Support
// @contact.url
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:1323
// @BasePath
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
func main() {
	fmt.Println("Only For Modify testing...")
}
