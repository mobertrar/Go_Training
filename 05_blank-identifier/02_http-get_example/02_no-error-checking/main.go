package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

//The Get function from the http package
//it returns a response or result (res)
//and an error ( _ )
//"look I know an error is coming and I'm not going to do anything with it"
