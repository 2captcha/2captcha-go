package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	yidun := api2captcha.Yidun{
		SiteKey: "6b4d7e0c4f5a4c7db2f3a1e8c9d6f123",
		Url:     "https://example.com/page-with-yidun",
	}

	req := yidun.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
