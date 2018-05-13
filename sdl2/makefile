DIRS := 01
SOURCES = $(wildcard $@/*.go)

DEP_FILE := Gopkg.lock

build: dep $(DIRS)

$(DEP_FILE):
	dep init

dep: $(DEP_FILE)
	dep ensure

$(DIRS):
	go build -o bin/$@ $(SOURCES)

.PHONY: build $(DIRS) dep