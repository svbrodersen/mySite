dev:
	templ generate
	go build -o ./tmp/main ./main.go 
	air

tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify


build: 
	make tailwind-build
	templ generate
	go build -o ./tmp/main ./main.go 


