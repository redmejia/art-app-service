down_service: clean_dist
	@echo "DOWN SERVICES"
	docker-compose down


up_service: build_service
	@echo "...STOP ALL"
	docker-compose down
	@echo "BUILD & START SERVICES..."
	docker-compose up --build

# build art service
build_service:
	@echo building service
	env GOOS=linux CGO_ENABLED=0 go build -o dist/art_service main.go

clean_dist:
	rm dist/art_service