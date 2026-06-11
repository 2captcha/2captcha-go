package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	binance := api2captcha.Binance{
		SiteKey:    "login",
		Url:        "https://example.com/page-with-binance",
		ValidateId: "cb0bfef...e54ecd57b",
	}

	req := binance.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
