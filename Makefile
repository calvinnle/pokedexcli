default:
	echo 'do nothing bro'

build:
	go build .

run: build
	./pokedexcli

test: 
	go test ./... -v

clean:
	rm -f ./pokedexcli

.PHONY: build run clean
