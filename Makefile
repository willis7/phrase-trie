# I usually keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision
VERSION         :=      $(shell cat ./VERSION)

all: install

install:
	go install -v

test:
	go test -v

fmt:
	go fmt


.PHONY: install test fmt
