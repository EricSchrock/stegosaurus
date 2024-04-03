.PHONY: all
all: build

.PHONY: build
build:
	go fmt
	go build

.PHONY: test
test:
	./stegosaurus
	./stegosaurus --save hi
	./stegosaurus --find hi
	./stegosaurus --save hi --find hi

.PHONY: clean
clean:
	rm stegosaurus

