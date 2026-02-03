package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	captchafox := api2captcha.Captchafox{
		SiteKey:   "sk_ILKWNruBBVKDOM7dZs59KHnDLEWiH",
		Url:       "https://mysite.com/page/with/captchafox",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		Proxy:     "username:password@1.2.3.4:5678",
		Proxytype: "http",
	}

	req := captchafox.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
