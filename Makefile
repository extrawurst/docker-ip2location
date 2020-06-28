
ZIP_FILE=ip2location.zip
DB_TYPE=DB1LITEBIN
DB_URL=http://www.ip2location.com/download/

update-db:
	@echo download db using token: ${IP2LOC_TOKEN}
	@wget "${DB_URL}?token=${IP2LOC_TOKEN}&file=${DB_TYPE}" -O $(ZIP_FILE)
	@unzip -o $(ZIP_FILE) -d ./data/
	@echo updated db

run:
	go run main.go

build-docker:
	docker build -t extrawurst/ip2location -f Dockerfile .