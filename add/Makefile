all:
	GOOS=wasip1 GOARCH=wasm tinygo build -o core.wasm main.go
	wasm-tools objdump core.wasm
	wasm-tools component embed ../wit/adder.wit core.wasm -o core.embedded.wasm
	wasm-tools objdump core.embedded.wasm
	wasm-tools component new core.embedded.wasm --adapt=wasi_snapshot_preview1.reactor.wasm -o core.component.wasm
	wasm-tools component wit core.component.wasm
