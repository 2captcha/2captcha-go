package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	reCaptcha := api2captcha.ReCaptcha{
		SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
		Url:     "ttps://mysite.com/page/with/recaptcha",
	}

	req := reCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
