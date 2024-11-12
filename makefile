.PHONY: dev start build

dev: 
	air

start:
	go run main.go

build:
	go build -o build/app.exe

clean:
	rm -Rf build/app.exe
	rmdir build