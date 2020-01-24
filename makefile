up : 
	cd client/ && yarn install
	docker-compose up --build
down :
	docker-compose down