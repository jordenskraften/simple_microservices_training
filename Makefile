t1:
	cd ./test1-micro && go run main.go

t2:
	cd ./test2-micro && go run main.go

dt1:
	cd ./test1-micro && docker build -t test1-micro -f micro.dockerfile .

dt2:
	cd ./test2-micro && docker build -t test2-micro -f micro.dockerfile .

dtr1:
	docker run -d -p 8080:8080 test1-micro

dtr2:
	docker run -d -p 50051:50051 test2-micro

don:
	docker-compose up -d

doff:
	docker-compose down