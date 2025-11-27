# Incluir el archivo .env
-include .env
export
# Go parameters
GOCMD=go
APP=api/api.go
GORUN=$(GOCMD) run $(APP) #for clean-architecture structure
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=md2pdf
BINARY_UNIX=$(BINARY_NAME)
BINARY_WIN=$(BINARY_NAME).exe
BINARY_MAC=$(BINARY_NAME).macos
PORT=3030
IP=127.0.0.1
URL=https://$(IP):$(PORT)
EXPOSE=$(IP):$(PORT):$(PORT) 
REPO_NAME=yber/$(BINARY_NAME)
VERSION=0.0.1
TAG=${VERSION}
COMMIT=$$(git rev-parse --short HEAD)

# install dependencies
mod:
	$(GOMOD) tidy

# ejecuta el api.go
run:
	$(GORUN)

air:
	air

build:
	CGO_ENABLED=0 $(GOBUILD) -trimpath -ldflags="-s -w -extldflags '-static'" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo -tags osusergo,netgo -o $(BINARY_NAME) -v $(APP)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOAMD64=v2 $(GOBUILD) -a -trimpath -ldflags="-s -w -extldflags '-static'" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo -tags osusergo,netgo -o $(BINARY_UNIX) -v $(APP)

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 GOAMD64=v2 $(GOBUILD) -a -trimpath -ldflags="-s -w -extldflags '-static'" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo -tags osusergo,netgo -o $(BINARY_MAC) -v $(APP)

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GOAMD64=v2 $(GOBUILD) -a -trimpath -ldflags="-s -w -extldflags '-static'" -gcflags="all=-m=0 -l=2 -dwarf=false" -installsuffix cgo -tags osusergo,netgo -o $(BINARY_WIN) -v $(APP)

#hace el build de la imagen Docker
dbuildcert:
	docker build --target cert -t $(REPO_NAME):$(TAG)-cert -f ./Containerfile .

dbuildprod:
	docker build --target prod -t $(REPO_NAME):$(TAG)-prod -f ./Containerfile .

#corre el contenedor en segundo plano
drund:
	docker run -dp $(EXPOSE) --name $(BINARY_NAME) -v md2pdf_data:/pdf $(REPO_NAME):$(TAG)-cert
# 	docker run -dp $(EXPOSE) --name $(BINARY_NAME) -v md2pdf_data:/pdf $(REPO_NAME):$(TAG)-prod
# 	docker run -dp $(EXPOSE) --name $(BINARY_NAME) --env-file=.env -v md2pdf_data:/pdf $(REPO_NAME):$(TAG)

#hace el stop del contenedor Docker
dclean:
	docker stop $(BINARY_NAME) || true
	docker rm $(BINARY_NAME) || true
	docker rmi $(REPO_NAME):$(TAG) || true

# git tag
version:
	git describe --tags --abbrev=0

gct:
	git tag -a ${VERSION} ${COMMIT} -m "version ${VERSION}"

gpt:
	git push ${REMOTE} ${VERSION}

# muestra el log de github (no añadir | head -n 10 | tail -n 5 | sed 's/^/  /')
log:
	git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative