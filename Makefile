r:
	go build -gcflags '-N -l' -o build/app
	./build/app
build:
	go build -o build/app -v -x
run:
	./build/app
