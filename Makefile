all: build

# go build -ldflags "-linkmode external -extldflags \"-static\" -s -w ${LDFLAGS}" -o bin/banner cmd/banner/main.go
.PHONY: build
build:
	mkdir -p bin
	go build -ldflags "-s -w ${LDFLAGS}" -o bin/banner cmd/banner/main.go

.PHONY: install
install:
	go install ./cmd/banner
