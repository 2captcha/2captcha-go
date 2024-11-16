package main

import (
	"fmt"
	api2captcha "github.com/2captcha/2captcha-go"
	helper "github.com/2captcha/2captcha-go/examples/internal"
	"os"
)

func main() {
	client := api2captcha.NewClient(os.Args[1])

	assetsDir := helper.GetAssetsDir(os.Args[0])
	fileName := assetsDir + "/" + "normal.jpg"

	normal := api2captcha.Normal{
		File: fileName,
	}

	req := normal.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
