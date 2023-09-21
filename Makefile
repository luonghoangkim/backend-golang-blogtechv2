pro:
	docker rmi -f web-blogtechv2:1.0
	docker-compose up
dev:
	cd cmd/dev; go run main.go