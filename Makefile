up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

build:
	docker-compose build --no-cache

rebuild:
	docker-compose up -d --build

stop:
	docker-compose stop

start:
	docker-compose start
