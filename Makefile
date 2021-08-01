base:
	docker build -t builder -f dockerfiles/Builder .

news: base
	docker build -t news -f dockerfiles/News .

user: base
	docker build -t user -f dockerfiles/User .
