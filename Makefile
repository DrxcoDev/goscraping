.PHONY: all docker run clean

all: build/scraper

docker:
	docker build -t goscraping .

run:
	./build/scraper

clean:
	rm -f build/scraper
	rm -f build/quotes.json
