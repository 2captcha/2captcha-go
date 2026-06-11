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
		GenericCaptcha: api2captcha.GenericCaptcha{
			UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
			Proxytype: "http",
			Proxy:     "username:password@1.2.3.4:5678",
		},
	}

	req := binance.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
