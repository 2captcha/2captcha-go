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

	backgroundFile := assetsDir + "/dragdrop/background.jpeg"
	image1File := assetsDir + "/dragdrop/image1.jpeg"
	image2File := assetsDir + "/dragdrop/image2.jpeg"

	backgroundBs := helper.ReadFile2BiteSlice(backgroundFile)
	image1Bs := helper.ReadFile2BiteSlice(image1File)
	image2Bs := helper.ReadFile2BiteSlice(image2File)

	backgroundB64 := base64.StdEncoding.EncodeToString(backgroundBs)
	image1B64 := base64.StdEncoding.EncodeToString(image1Bs)
	image2B64 := base64.StdEncoding.EncodeToString(image2Bs)

	dragDrop := api2captcha.DragDrop{
		Base64:   backgroundB64,
		Images:   []string{image1B64, image2B64},
		HintText: "Drag the images to proper position",
	}

	req := dragDrop.ToRequest()

	token, captchaId, err := client.Solve(req)

	fmt.Println("token ::: " + token)
	fmt.Println("captchaId ::: " + captchaId)
	fmt.Print("error ::: ")
	fmt.Println(err)
}
