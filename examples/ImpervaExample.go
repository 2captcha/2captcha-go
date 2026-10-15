package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	imperva := api2captcha.Imperva{
		Url:                "https://example.com/login",
		IncapsulaScriptUrl: "/_Incapsula_Resource?SWUDNSAI=31&xinfo=1",
		IncapsulaCookies:   "visid_incap_123456=sample; incap_ses_123_123456=sample",
		Reese84UrlEndpoint: "https://example.com/_Incapsula_Resource?reese84=1",
		UserAgent:          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
		Proxytype:          "http",
		Proxy:              "username:password@1.2.3.4:5678",
	}

	req := imperva.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
