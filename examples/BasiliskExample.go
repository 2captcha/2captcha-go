package main

import (
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	basilisk := api2captcha.Basilisk{
		SiteKey: "b7890h...19fb2600897",
		Url:     "https://example.com/page-with-basilisk",
	}

	req := basilisk.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
