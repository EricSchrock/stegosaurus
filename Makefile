.PHONY: all
all: build

.PHONY: build
build:
	go fmt
	go build

.PHONY: test
test:
	touch test.png test.jpg test.jpeg
	./stegosaurus
	./stegosaurus --save hi --find hi
	./stegosaurus --save abcdefghijklmnopqrstu
	./stegosaurus --find abcdefghijklmnopqrstu
	./stegosaurus --save hi
	./stegosaurus --find hi
	rm test.png test.jpg test.jpeg

.PHONY: clean
clean:
	rm stegosaurus

