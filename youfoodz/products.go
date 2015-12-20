package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "os"
)

type Configuration struct {
    Token      string
    UrlCount   string
}

func main() {

    file, err := os.Open("config.json")

    if err != nil {
        fmt.Println("Please copy config.json.sample to config.json and configure your application endpoint")
        return
    }

    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err = decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }

    client := &http.Client{}

    req, _ := http.NewRequest("GET", configuration.UrlCount, nil)

    req.Header.Add("X-Shopify-Access-Token", configuration.Token)
    resp, _ := client.Do(req)

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    fmt.Printf("Response: %s\n", body)
}