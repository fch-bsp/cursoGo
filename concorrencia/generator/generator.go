package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chain - Canal somente-leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // compilou expressão regular
			c <- r.FindStringSubmatch(string(html))[1]

		}(url) //criando funão e invogando

	}
	return c

}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros Colocados:", <-t1, "|", <-t2)
	fmt.Println("Seguntos Colocados :", <-t1, "|", <-t2)
}
