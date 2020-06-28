run:
	go run main.go

build-docker:
	docker build -t ip2location -f Dockerfile .