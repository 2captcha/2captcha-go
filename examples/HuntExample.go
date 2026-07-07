package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	hunt := api2captcha.Hunt{
		Url:       "https://example.com/page-with-hunt",
		ApiGetLib: "https://example.com/hd-api/external/apps/app-id/api.js",
		Proxytype: "http",
		Proxy:     "username:password@1.2.3.4:5678",
	}

	req := hunt.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
