all: build

build:
	docker build -t jacobwgillespie/redis .

shell:
	docker run -t -i jacobwgillespie/redis /bin/bash

clean:
	docker rmi jacobwgillespie/redis
