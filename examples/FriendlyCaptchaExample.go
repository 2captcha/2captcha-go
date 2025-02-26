package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	friendlyCaptcha := api2captcha.Friendly{
		SiteKey: "FCMST5VUMCBOCGQ9",
		Url:     "https://geizhals.de/455973138?fsean=5901747021356",
	}

	request := friendlyCaptcha.ToRequest()

	token, captchaId, err := client.Solve(request)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
