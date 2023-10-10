package main

import (
    "fmt"
    "log"
    "github.com/2captcha/2captcha-go"
)

func main() {
    client := api2captcha.NewClient("0cf96237fa3953ea12e27b74d2cdae5a")
    cap := api2captcha.Text{
        Text: "If tomorrow is Saturday, what day is today?",
        Lang: "en",
    }
    code, err := client.Solve(cap.ToRequest())
    if err != nil {
        log.Fatal(err);
    }
    fmt.Println("code "+code)
}
