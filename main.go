package main

import (
	"PowerTranslate/router"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
)

func main() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
	env := cfg.Section("env")
	addr, _ := env.GetKey("addr")
	runMode, _ := env.GetKey("runMode")

	gin.SetMode(runMode.String())
	r := router.InitRouter()
	err = r.Run(addr.String())
	if err != nil {
		log.Fatal(err)
	}
}
