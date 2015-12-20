package main

import (
    "fmt"
    "encoding/json"
    "net/http"
    "os"
    "./shopify"
)


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
    client := &http.Client{}

    req, _ := http.NewRequest("GET", configuration.UrlCount, nil)

    req.Header.Add("X-Shopify-Access-Token", configuration.Token)
    resp, _ := client.Do(req)

    defer resp.Body.Close()

    // decode result
    decoder = json.NewDecoder(resp.Body)
    countResult := shopify.CountResult{}
    err = decoder.Decode(&countResult)
    if err != nil {
      fmt.Println("error:", err)
      return;
    }

    fmt.Printf("%+v\n", countResult.Count)
}