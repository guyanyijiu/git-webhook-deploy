### git-webhook-deploy

A simple tool for automatically deploying GitHub projects using webhooks.


### Getting started

##### Build

```shell
cd $GOPATH/src

git clone https://github.com/guyanyijiu/git-webhook-deploy.git

cd git-webhook-deploy

go build -o git-webhook-deploy main.go
```

##### Configuration

```shell
vim config.yaml
```

##### Run

```
./git-webhook-deploy -c config.yaml
```
or run as a daemon
```
./git-webhook-deploy -c config.yaml -d
```


### GitHub Webhook settings

![Example screenshot showing GitHub webhook settings](https://github.com/guyanyijiu/images/blob/master/Jietu20190325-200253.png)

