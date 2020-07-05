
ZIP_FILE=ip2location.zip

update-db:
	@wget "https://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.BIN.ZIP" -O $(ZIP_FILE)
	@unzip -o $(ZIP_FILE) -d ./data/
	@wget "https://download.ip2location.com/lite/IP2LOCATION-LITE-DB1.IPV6.BIN.ZIP" -O $(ZIP_FILE)
	@unzip -o $(ZIP_FILE) -d ./data/
	@echo updated db

run:
	go run main.go

build-docker:
	docker build -t extrawurst/ip2location -f Dockerfile .