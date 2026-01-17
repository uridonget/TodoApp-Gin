name=todo-app-gin

build:
	docker build -t $(name) .

run: build
	docker run -p 8080:8080 --name $(name) $(name)

stop:
	-docker stop $(name)
	-docker rm $(name)

clean: stop
	-docker rmi $(name)

.PHONY: build run stop clean