build:
	go build .

run: build
	./pokedexcli

clean:
	rm -f ./pokedexcli

.PHONY: build run clean
