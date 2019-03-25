package handler

import (
	"encoding/json"
	"git-webhook-deploy/config"
	"git-webhook-deploy/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

type pushEventPayload struct {
	Ref        string
	Repository struct {
		Clone_url string
	}
}

func Github(c *gin.Context) {

	event := c.GetHeader("X-Github-Event")
	if event == "" {
		log.Error("fail to get header X-GitHub-Event")
		return
	}

	delivery := c.GetHeader("X-GitHub-Delivery")
	if delivery == "" {
		log.Error("fail to get header X-GitHub-Delivery")
		return
	}

	sign := c.GetHeader("X-Hub-Signature")
	if sign == "" {
		log.Error("fail to get header X-Hub-Signature")
		return
	}
	
	_, _ = event, delivery

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error("fail to get request body, ", err)
		return
	}

	payload := &pushEventPayload{}
	err = json.Unmarshal(body, payload)
	if err != nil {
		log.Error("fail to parse request body, ", err)
	}

	repoConfig := config.FindRepositoryConfig("github", payload.Repository.Clone_url)
	if repoConfig == nil {
		log.Error("fail to find repository config, ", payload.Repository.Clone_url)
		return
	}

	if ! strings.HasSuffix(payload.Ref, repoConfig.Branch) {
		return
	}

	mySign := HmacSha1(body, []byte(repoConfig.Secret))
	if "sha1="+mySign != sign {
		log.Error("fail to verify sign")
		return
	}

	deploy(repoConfig)
}
