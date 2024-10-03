dev:
	air

build: templ
	go build -o ./tmp/main .

templ:
	templ generate

tailwind:
	tailwindcss -i css/main.css -o dist/main.css