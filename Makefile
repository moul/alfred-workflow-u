GOBINS = .
GOPKG ?=	moul.io/alfred-workflow-u
VERSION ?= `git describe --tags --always`

include rules.mk

bundle:
	go build -o ./workflow/alfred-workflow-u .
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o "./workflow/alfred-workflow-u"; \
	VERSION=v$(VERSION) envsubst > ./workflow/info.plist < ./workflow/info.plist.template;
	cd workflow && zip -r "./alfred-workflow-u.zip" ./*
	cd workflow && zip -d "./alfred-workflow-u.zip" ./info.plist.template
		mv "workflow/alfred-workflow-u.zip" "./alfred-workflow-u-v$(VERSION).alfredworkflow"

lint:
	cd tool/lint; make
.PHONY: lint
