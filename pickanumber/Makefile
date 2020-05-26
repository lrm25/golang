GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
EXECUTABLE=pickanumber
BUILD_FOLDER=build

all: compile
compile:
	mkdir -p $(BUILD_FOLDER)
	$(GOBUILD) -o $(BUILD_FOLDER)/$(EXECUTABLE)
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_FOLDER)
