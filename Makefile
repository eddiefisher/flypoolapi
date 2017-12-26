.PHONY: build run

BINARY=flypoolapi

.DEFAULT_GOAL: $(BINARY)

# Builds the project
$(BINARY):
	go build -v -o ../bin/${BINARY} ./*.go

run:
	go run ./*.go