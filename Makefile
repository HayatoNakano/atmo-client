OUT_DIR = out

all: test build

build: clean
	mkdir -p ${OUT_DIR}
	GOOS=linux GOARCH=arm64 go build -o ./$(OUT_DIR)/co2client ./cmd/co2client/main.go
	# GOOS=linux GOARCH=arm64 go build -o ./$(OUT_DIR)/natureclient ./cmd/natureclient/main.go

test:
	go test ./...

clean:
	rm -rf ${OUT_DIR}