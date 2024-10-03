dev:
	air

build: templ tailwind
	go build -o ./tmp/main .

test: templ tailwind
	go test ./...

templ:
	templ generate

tailwind:
	tailwindcss -i css/main.css -o dist/main.css