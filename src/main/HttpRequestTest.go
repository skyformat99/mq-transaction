package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	resp, err := http.Get("http://47.97.109.99:9995/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}