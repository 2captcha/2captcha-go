package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	tspd := api2captcha.Tspd{
		Url:            "https://example.com/page-with-tspd",
		TspdCookie:     "TS386a400d029=082670...010245; TS386a400d078=082670...dbb3b0c",
		HtmlPageBase64: "PCFET0NUWVBFIGh0bWw+...",
		Proxytype:      "http",
		Proxy:          "username:password@1.2.3.4:5678",
	}

	req := tspd.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
