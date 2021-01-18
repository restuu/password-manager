run:
	docker-compose up --build

stop:
	docker-compose down

.PHONE: run stop