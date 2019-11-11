package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writen, _ := io.Copy(os.Stdout, resp.Body)

	fmt.Println(writen)

}
