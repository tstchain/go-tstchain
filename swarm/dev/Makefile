.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/tstchain/go-tstchain/cmd/swarm
	go build -o bin/gtst     github.com/tstchain/go-tstchain/cmd/gtst
	go build -o bin/bootnode github.com/tstchain/go-tstchain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/tstchain/go-tstchain/swarm/...
