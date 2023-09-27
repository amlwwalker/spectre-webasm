build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	cd web && npm run build
	go build -o	go-app-tuts

run: build
	./go-app-tuts

