.PHONY: all
all: build

.PHONY: build
build:
	go build

.PHONY: test
test:
	./stegosaurus --save hi --find hey

.PHONY: clean
clean:
	rm stegosaurus

