package handler

import (
	"git-webhook-deploy/config"
	"git-webhook-deploy/log"
	"os"
	"strings"
)

var deployChan = make(chan *config.Repository, deployChanCap)
var deployChanCap = 3

func init() {
	go deployHandler()
}

func deploy(c *config.Repository) {
	if len(deployChan) < deployChanCap {
		deployChan <- c
	}
}

func deployHandler() {
	for c := range deployChan {
		if PathIsExist(c.Path + "/" + c.Name) {
			gitPull(c.Path, c.Name, c.Branch, c.Script)
		} else {
			gitClone(c.Url, c.Path, c.Name, c.Branch, c.Script)
		}
	}
}

func gitClone(url string, path string, name string, branch string, script string) bool {
	if ! PathIsWritable(path) {
		log.Error("fail to git clone, ", path+" not writable")
		return false
	}
	err := os.Chdir(path)
	if err != nil {
		log.Error("fail to git clone, ", err)
		return false
	}

	out, err := ExecCommand(config.Config.Git, "clone", url, name)
	if err != nil {
		log.Error("fail to git clone, ", err)
		return false
	}
	err = os.Chdir(path + "/" + name)
	if err != nil {
		log.Error("fail to git clone, ", err)
		return false
	}

	out, err = ExecCommand(config.Config.Git, "checkout", branch)
	if err != nil {
		log.Error("fail to git checkout, ", err)
		return false
	}

	log.Info("success to git clone")
	log.Info(out)
	ExecScript(script)
	return true
}

func gitPull(path string, name string, branch string, script string) bool {
	project := path + "/" + name
	err := os.Chdir(project)
	if err != nil {
		log.Error("fail to git pull, ", err)
		return false
	}
	out, err := ExecCommand(config.Config.Git, "checkout", branch)
	if err != nil {
		log.Error("fail to git checkout, ", err)
		return false
	}

	out, err = ExecCommand(config.Config.Git, "pull")
	if err != nil {
		log.Error("fail to git pull, ", err)
		return false
	}

	log.Info("success to git pull")
	log.Info(out)
	ExecScript(script)
	return true
}

func ExecScript(script string) {
	bins := strings.Fields(script)
	if len(bins) > 0 {
		out, err := ExecCommand(bins[0], bins[1:]...)
		if err != nil {
			log.Error("fail to execute script, ", err)
		} else {
			log.Info("success to execute script")
			log.Info(out)
		}
	}
}
