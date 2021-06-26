docker-build:
	docker build -t widgethubbackend -f Dockerfile-golang .
	docker build -t express-app -f express-app/Dockerfile-express .

docker-run:
	docker-compose -f docker-compose.yml up
