package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("URL")
	if url == "" {
		log.Fatal("URL env var can not be empty")
	}
	output := os.Getenv("OUTPUT_FOLDER")
	if output == "" {
		log.Fatal("OUTPUT_FOLDER env var can not be empty")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	pathFile := fmt.Sprintf("%s/result.json", output)
	err = ioutil.WriteFile(pathFile, data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
