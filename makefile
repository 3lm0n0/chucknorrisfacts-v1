##################################################
# Usage:
##################################################
# make          # compile all binary
# make hello    # prints hello
# make init     # creates module
# make setup    # sets up the microservice
# make build    # builds the microservice
# make run      # runs the microservice
# make test     # run tests
# make clean    # remove ALL binaries and objects

BINARY_NAME=fact

.PHONY:= hello setup build run test
.DEFAULT_GOAL:= setup build run

hello:
	echo "Hello"

init:
	@echo "=> Go module fact initializing"
	@go mod init 'fact'

setup:
	@echo "=> Stetting microservice"
	@export GOSUMDB=off
	@go mod tidy
	@go mod download
	@echo "=> Setup completed"

build:
	@echo "=> Building microservice"
	@go build -o ./bin/${BINARY_NAME}
	
run:
	./bin/${BINARY_NAME}

test:
	go test -v ./...

clean:
	@echo "Cleaning up all binaries, objects and sum ..."
	@go clean
	@rm -rvf *.o ./bin/${BINARY_NAME} go.sum