package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

url := "http://www.chicagomusicmagazine.com/"

func main() {
  res, _ := http.Get(url)
  // res, _ := http.Get("http://www.chicagomusicmagazine.com/")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()
  fmt.Printf("%s", page)
}
