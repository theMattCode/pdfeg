
core-test:
	cd packages/core && go test -json ./...

core-build-wasm:
	cd packages/core && GOOS=js GOARCH=wasm go build -o pdfeg.wasm
