package pokeapi

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func (c *Client) CatchPokemon(name string) (bool, error) {
	if len(name) == 0 {
		return false, fmt.Errorf("name is not defined")
	}

	// store full url to make use of caching
	endpoint := "/pokemon/" + name
	fullUrl := BaseURL + endpoint

	var pokemon Pokemon
	// cache hit
	data, ok := c.clientCache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit!")
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return false, err
		}

		//battle
		fmt.Println("enter a number from 1 to 10: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		numStr := scanner.Text()

		// Parse the user's input as an integer
		userNum, err := strconv.Atoi(numStr)
		if err != nil {
			return false, err
		}

		// Define the maximum value as 10 (for range 1-10)
		max := big.NewInt(10)

		// Generate a random number in the range [0, 9]
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Fatal(err)
		}

		// Convert big.Int to int64, then to int for comparison
		// Add 1 to shift range from [0,9] to [1,10]
		randomNum := int(n.Int64()) + 1

		if userNum == randomNum {
			fmt.Printf("Gotcha! You caught %s!", pokemon.Name)
            return true, nil
		} else {
			fmt.Printf("The number was %d, not %d. You didn't catch %s!", randomNum, userNum, pokemon.Name)
		}

		return false, nil
	}

	fmt.Println("cache miss!")
	// cache miss
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return false, fmt.Errorf("error creating new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

    // for more verbose error handling
	if resp.StatusCode == http.StatusNotFound {
		return false, fmt.Errorf("%s is not a valid pokemon name", name)
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad status: %d %s",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("Error reading body")
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return false, err
	}

	// cache the data
	c.clientCache.Set(fullUrl, data)

	//battle

	fmt.Println("enter a number from 1 to 10: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numStr := scanner.Text()

	// Parse the user's input as an integer
	userNum, err := strconv.Atoi(numStr)
	if err != nil {
		return false, err
	}

	// Define the maximum value as 10 (for range 1-10)
	max := big.NewInt(10)

	// Generate a random number in the range [0, 9]
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}

	// Convert big.Int to int64, then to int for comparison
	// Add 1 to shift range from [0,9] to [1,10]
	randomNum := int(n.Int64()) + 1

	if userNum == randomNum {
		fmt.Printf("Gotcha! You caught %s!", pokemon.Name)
        return true, nil
	} else {
		fmt.Printf("The number was %d, not %d. You didn't catch %s!", randomNum, userNum, pokemon.Name)
	}

	return false, nil
}
