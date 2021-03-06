GOFMT_FILES?=$$(find . -name '*.go')

all: fmt build

fmt:
	gofmt -w $(GOFMT_FILES)

build:
	find . -name ".DS_Store" -delete
	go mod tidy
	go run main.go pack
	go build -ldflags "-X github.com/swiftcarrot/dashi/cmd.Version=`git rev-parse HEAD`" -o bin/dashi
	go install

clean:
	rm bin/dashi
