BINARY_NAME=IsShounicFull

run:
	@echo "since you didn't typed anything other than make, this target (run) will we run by default."
	@echo "If you want to build this program instead, type 'make build'. "
	@go run main.go

build:
	@echo "Building For Release... (May take longer if this is the first time running or if you ran 'go clean' beforehand)"
	@echo "Building Windows x86..."
	CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -x -o ${BINARY_NAME}-windows-x86 main.go 
	@echo "Building Linux x86_64..."
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -x -o ${BINARY_NAME}-linux-x86_64 main.go 
	@echo "Building Windows x86_64..."
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -x -o ${BINARY_NAME}-windows-x86_64 main.go 
	@echo "Building Windows arm64..."
	CGO_ENABLED=1 CC=aarch64-w64-mingw32-gcc GOOS=windows GOARCH=arm64 go build -x -o ${BINARY_NAME}-windows-arm64 main.go 

clean:
	go clean