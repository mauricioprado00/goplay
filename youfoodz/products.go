package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "os"
    //"io/ioutil"
    "./shopify"
)

/**
 * https://gobyexample.com/json
 */
func main() {

    file, err := os.Open("config.json")

    if err != nil {
        fmt.Println("Please copy config.json.sample to config.json and configure your application endpoint")
        return
    }

    decoder := json.NewDecoder(file)
    configuration := shopify.Configuration{}
    err = decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
      return;
    }

    // http query
    // https://golang.org/pkg/net/http/
    client := &http.Client{}

    req, _ := http.NewRequest("GET", configuration.UrlList, nil)

    req.Header.Add("X-Shopify-Access-Token", configuration.Token)
    resp, _ := client.Do(req)

    defer resp.Body.Close()


    // body, _ := ioutil.ReadAll(resp.Body)

    // fmt.Printf("%s\n", body)

    // return;

    // decode result
    // https://golang.org/pkg/encoding/json/
    // http://blog.golang.org/json-and-go
    // https://coderwall.com/p/4c2zig/decode-top-level-json-array-into-a-slice-of-structs-in-golang
    // https://gobyexample.com/json
    decoder = json.NewDecoder(resp.Body)
    products := shopify.Products{}
    err = decoder.Decode(&products)
    if err != nil {
      fmt.Println("error:", err)
      return;
    }

    fmt.Printf("%+v\n", products)
    fmt.Printf("%+v\n", configuration)
}