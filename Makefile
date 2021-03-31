GOPKG ?=	moul.io/alfred-workflow-u
DOCKER_IMAGE ?=	moul/alfred-workflow-u
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ alfred-workflow-u hello world' > .tmp/usage.txt
	alfred-workflow-u hello world 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate

lint:
	cd tool/lint; make
.PHONY: lint
