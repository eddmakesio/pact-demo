# make -j3
target: run-service-a run-service-b run-service-c

run-service-a:
	./service_a/service_a

run-service-b:
	./service_b/service_b

run-service-c:
	./service_c/service_c

build:
	cd ./service_a && go build
	cd ./service_b && go build
	cd ./service_c && go build

clean:
	rm ./service_a/service_a
	rm ./service_b/service_b
	rm ./service_c/service_c