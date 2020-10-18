.PHONY: proto data build

GOPATH=/Users/huchenhao/go

proto:
	for d in api services; do \
		for f in $$d/**/api/*.proto; do \
			protoc -I$(GOPATH)/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20201016211454-9db36e164668 -I. -I$(GOPATH)/pkg/mod/github.com/protocolbuffers/protobuf@v3.13.0+incompatible/src --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done

lint:
	./bin/lint.sh

build:
	./bin/build.sh

data:
	go-bindata -o data/bindata.go -pkg data data/*.json

run:
	docker-compose build
	docker-compose up
