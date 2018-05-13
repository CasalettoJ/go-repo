DIRS := 01 02 03
SOURCES = $(wildcard $@/*.go)

DEP_FILE := Gopkg.lock

build: dep $(DIRS)

$(DEP_FILE):
	dep init

dep: $(DEP_FILE)
	dep ensure

$(DIRS):
	go build -o bin/$@ $(SOURCES)
	if [ -d "$@/assets" ]; then \
		cp -R $@/assets/ bin/assets/; \
    fi;

.PHONY: build $(DIRS) dep