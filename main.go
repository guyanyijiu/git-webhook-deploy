package main

import (
	"flag"
	"fmt"
	"git-webhook-deploy/config"
	"git-webhook-deploy/log"
	"git-webhook-deploy/router"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)

var (
	configFile string
	daemon     bool
)

func init() {
	flag.StringVar(&configFile, "c", "", `set configuration file (default "config.yaml")`)
	flag.BoolVar(&daemon, "d", false, "run in daemon mode")
	flag.Parse()

	if configFile == "" {
		flag.Usage()
		return
	}

	if daemon {
		cmd := exec.Command(os.Args[0], "-c", configFile)
		err := cmd.Start()
		if err != nil {
			fmt.Println("fail to run in daemon mode")
			os.Exit(0)
		}
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}

func main() {
	err := config.Init(configFile)
	if err != nil {
		fmt.Println("fail to read configuration file")
		return
	}
	log.Init(config.Config.Log)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Writer()
	r := gin.Default()
	router.InitRouters(r)

	err = r.Run(config.Config.Host + ":" + config.Config.Port)
	if err != nil {
		fmt.Println("fail to start http server, ", err)
	}
}
