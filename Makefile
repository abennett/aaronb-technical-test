base:
	docker build -t builder -f dockerfiles/Builder .

news: base
	docker build -t news -f dockerfiles/News .

user: base
	docker build -t user -f dockerfiles/User .

all: base user news

test: base user news
	docker-compose up -d
	sleep 1
	go test . -race

clean:
	docker-compose down
