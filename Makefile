docker_build:
	docker build -t natomanga-cli .

docker_run:
	docker run -it -v ~/Desktop/natomanga-cli:/root/Desktop/natomanga-cli natomanga-cli
