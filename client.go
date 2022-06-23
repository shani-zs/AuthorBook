package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8060/shani")
	if err != nil {
		fmt.Errorf("error. error: %s", err.Error())
	}

	defer resp.Body.Close()

	body, er := io.ReadAll(resp.Body)
	if er != nil {
		fmt.Errorf("failed for fetch response body. got error:%v", er)
	}

	fmt.Println(string(body))
}
