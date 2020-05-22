# Makefile for releasing news
#
# The release version is controlled from pkg/version

NAME:=news
GOPRIVATE:=gitlab.mcsolutions.ru
DOCKER_REPOSITORY:=gitlab.mcsolutions.ru:4567/find-psy/back/admins
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
DOCKER_PUBLIC_IMAGE:=timurkash/find-psy-news-sfpg-v1
GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"' | tr -d '[:space:]')
PORT:=3008

vendor:
	export GOPRIVATE
	go mod init gitlab.mcsolutions.ru/find-psy/back/admins/news
	go mod vendor

revendor:
	rm -rf vendor go.mod go.sum
	make vendor

update-vendor:
	make revendor

init:
	git init
	git add .
#	git remote add origin git@gitlab.mcsolutions.ru:find-psy/back/admins/news.git
	git remote add origin https://gitlab.mcsolutions.ru/find-psy/back/admins/news.git
	git commit -m "Init"
	git push origin master

commit:
	git add .
	git commit -m "autocommit"
	git push -u origin master

next:
	git add .
	git commit -m "next"
	git push -u origin master

run:
	GO111MODULE=on go run -ldflags "-s -w -X gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/version.REVISION=$(GIT_COMMIT)" cmd/news/* --level=debug

test:
	GO111MODULE=on go test -v -race ./...

build:
	GO111MODULE=on GIT_COMMIT=$$(git rev-list -1 HEAD) && GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor  -ldflags "-s -w -X gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/news ./cmd/news/*

build-charts:
	helm lint charts/*
	helm package charts/*

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .

run-container:
	@docker run -dp $(PORT):$(PORT) --name=news $(DOCKER_IMAGE_NAME):$(VERSION)

push-container:
	docker tag $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_IMAGE_NAME):latest
	docker push $(DOCKER_IMAGE_NAME):$(VERSION)
	docker push $(DOCKER_IMAGE_NAME):latest

push-container-public:
	docker tag  $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_PUBLIC_IMAGE):$(VERSION)
	docker tag  $(DOCKER_IMAGE_NAME):$(VERSION) $(DOCKER_PUBLIC_IMAGE):latest
	docker push $(DOCKER_PUBLIC_IMAGE):$(VERSION)
	docker push $(DOCKER_PUBLIC_IMAGE):latest

run-compose:
	docker-compose up -d

stop-container:
	@docker stop $(NAME)
	@docker ps

clear-container:
	@docker rm -f $(NAME)

net-compose:
	docker network create -d=bridge find-psy.net

screen:
	@echo Running $(NAME) binary
	screen -S $(NAME) ./bin/$(NAME)

version-set:
	@next="latest" && \
	current="$(VERSION)" && \
	sed -i '' "s/$$current/$$next/g" pkg/version/version.go && \
	sed -i '' "s/tag: $$current/tag: $$next/g" charts/news/values.yaml && \
	sed -i '' "s/appVersion: $$current/appVersion: $$next/g" charts/news/Chart.yaml && \
	sed -i '' "s/version: $$current/version: $$next/g" charts/news/Chart.yaml && \
	sed -i '' "s/news:$$current/news:$$next/g" kustomize/deployment.yaml && \
	echo "Version $$next set in code, deployment, chart and kustomize"

release:
	git tag $(VERSION)
	git push origin $(VERSION)

swagger:
	GO111MODULE=on go get github.com/swaggo/swag/cmd/swag
	cd pkg/api && $$(go env GOPATH)/bin/swag init -g server.go

rmi-win:
	@echo In Windows PowerShell
	#	docker rmi $(docker images --format "{{.Repository}}:{{.Tag}}" | findstr "gitlab.mcsolutions.ru:4567/find-psy/back/admins/news")

rmi:
	docker rmi $(docker images | grep $(DOCKER_IMAGE_NAME))

list:
	docker images $(DOCKER_IMAGE_NAME)

logs:
	docker logs $(NAME)

in:
	docker exec -it $(NAME) sh

port-forward-staging:
	kubectl -n staging port-forward $(kubectl get pods -n staging -l app.kubernetes.io/name=$(NAME) -o name) 1$(PORT):$(PORT)

port-forward-production:
	kubectl -n production port-forward $(kubectl get pods -n production -l app.kubernetes.io/name=$(NAME) -o name) 1$(PORT):$(PORT)

template:
	helm repo add find-psy https://timurkash.github.com/helm-example/
	helm repo update
	helm template events-sfpg find-psy/sfpg -f=values.yaml > template.yaml

uninstall:
	helm uninstall news-sfpg -n find-psy-staging
