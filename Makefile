.PHONY: all
all: build

.PHONY: build
build:
	go fmt
	go build

.PHONY: test
test: errors ok

.PHONY: errors
errors:
	./stegosaurus
	./stegosaurus --save hi --find hi
	./stegosaurus --save abcdefghijklmnopqrstu
	./stegosaurus --find abcdefghijklmnopqrstu

.PHONY: ok
ok:
	./stegosaurus --save hi
	./stegosaurus --find hi

.PHONY: clean
clean:
	rm stegosaurus

