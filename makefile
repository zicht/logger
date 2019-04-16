test:
	if [ ! -d "./cover/" ]; then mkdir ./cover; fi
	go test -coverprofile ./cover/logger.out
	go tool cover -html=./cover/logger.out -o ./cover/logger.html