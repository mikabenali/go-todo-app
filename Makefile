build:
	sudo docker-compose build
	
run: build
	 sudo docker-compose up -d

run-dev: build
	sudo docker-compose -f docker-compose-dev.yml up -d
	
stop:
	sudo docker-compose stop
