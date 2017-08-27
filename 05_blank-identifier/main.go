package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	/*res, err := http.Get("http://www.uol.com.br")
	ir err := nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err := nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)*/

	res, _ := http.Get("http://www.uol.com.br/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
