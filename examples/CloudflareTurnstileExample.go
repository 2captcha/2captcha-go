package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	cloudflareTurnstile := api2captcha.CloudflareTurnstile{
		SiteKey: "0x4AAAAAAAChNiVJM_WtShFf",
		Url:     "https://ace.fusionist.io",
	}

	req := cloudflareTurnstile.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
