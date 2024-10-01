.PHONY: all build run-datanode run-director run-namenode run-mercenarios clean

all: build

build:
    docker build -f Dockerfile-director -t director:latest .
    docker build -f Dockerfile-datanode -t datanode:latest .
    docker build -f Dockerfile-namenode -t namenode:latest .

run-datanode:
    docker run --rm -d --name datanode datanode:latest

run-director:
    docker run --rm -d --name director director:latest

run-namenode:
    docker run --rm -d --name namenode namenode:latest

run-mercenarios:
    ssh dist@dist050.inf.santiago.usm.cl "cd /path/to/mercenarios && go run main.go"
    ssh dist@dist051.inf.santiago.usm.cl "cd /path/to/mercenarios && go run main.go"
    ssh dist@dist052.inf.santiago.usm.cl "cd /path/to/mercenarios && go run main.go"

clean:
    docker stop datanode director namenode
