docker-build:
	docker build -t widgethubbackend Dockerfile-golang
	docker build -t express-app Dockerfile-express

docker-run:
	docker-compose up