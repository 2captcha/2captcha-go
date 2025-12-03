package main

import (
	"encoding/base64"
	"fmt"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
	helper "github.com/2captcha/2captcha-go/examples/internal"
)

func main() {

	client := api2captcha.NewClient(os.Args[1])

	assetsDir := helper.GetExamplesAssetsDir(os.Args[0])

	bodyFile := assetsDir + "/temu/body.png"
	part1File := assetsDir + "/temu/part1.png"
	part2File := assetsDir + "/temu/part2.png"
	part3File := assetsDir + "/temu/part3.png"

	bodyBs := helper.ReadFile2BiteSlice(bodyFile)
	part1Bs := helper.ReadFile2BiteSlice(part1File)
	part2Bs := helper.ReadFile2BiteSlice(part2File)
	part3Bs := helper.ReadFile2BiteSlice(part3File)

	body := base64.StdEncoding.EncodeToString(bodyBs)
	part1 := base64.StdEncoding.EncodeToString(part1Bs)
	part2 := base64.StdEncoding.EncodeToString(part2Bs)
	part3 := base64.StdEncoding.EncodeToString(part3Bs)

	temuCaptcha := api2captcha.Temu{
		Body:  body,
		Part1: part1,
		Part2: part2,
		Part3: part3,
	}

	req := temuCaptcha.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
