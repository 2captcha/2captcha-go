package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	altcha := api2captcha.Altcha{
		Url: "https://mysite.com/page/with/altcha",
		ChallengeJson: `{"algorithm":"SHA-256","challenge":"...","signature":"..."}`,
	}

	req := altcha.ToRequest("altcha")

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
