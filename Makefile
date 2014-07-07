all: build

build:
	docker build -t playlist/redis image

shell:
	docker run -t -i playlist/redis /bin/bash

clean:
	docker rmi playlist/redis
