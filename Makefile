dev:
	air

build: templ
	go build -o ./tmp/main .

templ:
	templ generate