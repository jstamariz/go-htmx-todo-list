run_dev:
	docker compose up -d
	air
run:
	docker compose up -d
	make build
	./tmp/main.exe
build:
	make tailwind
	go build -o ./tmp/main.exe
tailwind:
	tailwind -i ./resources/css/input.css -o ./dist/output.css
dev_setup:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.2/tailwindcss-windows-x64.exe
	chmod +x tailwindcss-windows-x64.exe
	mv tailwindcss-windows-x64.exe tailwind
	mv tailwind C:\bin\tailwind
	go install github.com/cosmtrek/air@latest