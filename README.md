<a href="https://github.com/2captcha/2captcha-python"><img src="https://github.com/user-attachments/assets/37e1d860-033b-4cf3-a158-468fc6b4debc" width="82" height="30"></a>
<a href="https://github.com/2captcha/2captcha-javascript"><img src="https://github.com/user-attachments/assets/4d3b4541-34b2-4ed2-a687-d694ce67e5a6" width="36" height="30"></a>
<a href="https://github.com/2captcha/2captcha-go"><img src="https://github.com/user-attachments/assets/5e37ab36-f32f-464b-9d33-a335e2e1b13e" width="63" height="30"></a>
<a href="https://github.com/2captcha/2captcha-ruby"><img src="https://github.com/user-attachments/assets/0270d56f-79b0-4c95-9b09-4de89579914b" width="75" height="30"></a>
<a href="https://github.com/2captcha/2captcha-cpp"><img src="https://github.com/user-attachments/assets/36de8512-acfd-44fb-bb1f-b7c793a3f926" width="45" height="30"></a>
<a href="https://github.com/2captcha/2captcha-php"><img src="https://github.com/user-attachments/assets/e8797843-3f61-4fa9-a155-ab0b21fb3858" width="52" height="30"></a>
<a href="https://github.com/2captcha/2captcha-java"><img src="https://github.com/user-attachments/assets/a3d923f6-4fec-4c07-ac50-e20da6370911" width="50" height="30"></a>
<a href="https://github.com/2captcha/2captcha-csharp"><img src="https://github.com/user-attachments/assets/f4d449de-780b-49ed-bb0a-b70c82ec4b32" width="38" height="30"></a>

# Golang Module for 2Captcha API
The easiest way to quickly integrate [2Captcha] into your code to automate solving of any type of captcha.
Examples of API requests for different captcha types are available on the [Golang captcha solver](https://2captcha.com/lang/go) page.
- [Golang Module for 2Captcha API](#golang-module-for-2captcha-api)
  - [Installation](#installation)
  - [Configuration](#configuration)
      - [Client instance options](#client-instance-options)
  - [Solve captcha](#solve-captcha)
    - [Captcha options](#captcha-options)
    - [Basic example](#basic-example)
    - [Normal Captcha](#normal-captcha)
    - [Text Captcha](#text-captcha)
    - [reCAPTCHA v2](#recaptcha-v2)
    - [reCAPTCHA v3](#recaptcha-v3)
    - [reCAPTCHA Enterprise](#recaptcha-enterprise)
    - [FunCaptcha](#funcaptcha)
    - [GeeTest](#geetest)
    - [GeeTest V4](#geetest-v4)
    - [KeyCaptcha](#keycaptcha)
    - [Capy](#capy)
    - [Grid](#grid)
    - [Canvas](#canvas)
    - [ClickCaptcha](#clickcaptcha)
    - [Rotate](#rotate)
    - [Amazon WAF](#amazon-waf)
    - [Cloudflare Turnstile](#cloudflare-turnstile)
    - [Lemin Cropped Captcha](#lemin-cropped-captcha)
    - [CyberSiARA](#cybersiara)
    - [DataDome](#datadome)
    - [MTCaptcha](#mtcaptcha)
    - [Yandex](#yandex)
    - [Tencent](#tencent)
    - [atbCAPTCHA](#atbcaptcha)
    - [Cutcaptcha](#cutcaptcha)
    - [Friendly Captcha](#friendly-captcha)
    - [Audio Captcha](#audio-captcha)
    - [Prosopo](#prosopo)
    - [Captchafox](#captchafox)
  - [Other methods](#other-methods)
    - [send / getResult](#send--getresult)
    - [balance](#balance)
    - [report](#report)
  - [Proxies](#proxies)
  - [Examples](#examples)
- [Get in touch](#get-in-touch)
- [Join the team 👪](#join-the-team-)
- [License](#license)
   - [Graphics and Trademarks](#graphics-and-trademarks)

## Installation
To install the api client, use this:

```bash
go get -u github.com/2captcha/2captcha-go
```

## Configuration

Import the module like this:
```go
import (
        "github.com/2captcha/2captcha-go"
)
```
`Client` instance can be created like this:
```go
client := api2captcha.NewClient("YOUR_API_KEY")
```
There are few options that can be configured:
```go
client.SoftId = 123
client.Callback = "https://your.site/result-receiver"
client.DefaultTimeout = 120
client.RecaptchaTimeout = 600
client.PollingInterval = 100
```

### Client instance options

|Option|Default value|Description|
|---|---|---|
|soft_id|4583|Your software ID obtained after publishing in [2captcha software catalog]|
|callback|-|URL of your web-sever that receives the captcha recognition result. The URl should be first registered in [pingback settings] of your account|
|default_timeout|120|Timeout in seconds for all captcha types except reCAPTCHA. Defines how long the module tries to get the answer from `res.php` API endpoint|
|recaptcha_timeout|600|Timeout for reCAPTCHA in seconds. Defines how long the module tries to get the answer from `res.php` API endpoint|
|polling_interval|10|Interval in seconds between requests to `res.php` API endpoint, setting values less than 5 seconds is not recommended|

>  **IMPORTANT:** once *callback URL* is defined for `client` instance, all methods return only the captcha ID and DO NOT poll the API to get the result. The result will be sent to the callback URL.

To get the answer manually use [GetResult method](#send--getresult)

## Solve captcha
When you submit any image-based captcha use can provide additional options to help 2captcha workers to solve it properly.

### Captcha options
|Option|Default Value|Description|
|---|---|---|
|numeric|0|Defines if captcha contains numeric or other symbols [see more info in the API docs][post options]|
|min_len|0|minimal answer length|
|max_len|0|maximum answer length|
|phrase|0|defines if the answer contains multiple words or not|
|case_sensitive|0|defines if the answer is case sensitive|
|calc|0|defines captcha requires calculation|
|lang|-|defines the captcha language, see the [list of supported languages] |
|hint_img|-|an image with hint shown to workers with the captcha|
|hint_text|-|hint or task text shown to workers with the captcha|

Below you can find basic examples for every captcha type, check out the code below.

### Basic example
Example below shows a basic solver call example with error handling.

```go
captcha := api2captcha.Normal{
   File: "/path/to/normal.jpg",
}

code, err := client.Solve(captcha.ToRequest())
if err != nil {
	if err == api2captcha.ErrTimeout {
		log.Fatal("Timeout");
	} else if err == api2captcha.ErrApi {
		log.Fatal("API error");
	} else if err == api2captcha.ErrNetwork {
		log.Fatal("Network error");
	} else {
		log.Fatal(err);
	}
}
fmt.Println("code "+code)
```

### Normal Captcha

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_normal_captcha)</sup>

To bypass a normal captcha (distorted text on image) use the following method. This method also can be used to recognize any text on the image.

```go
captcha:= api2captcha.Normal{
   File: "/path/to/normal.jpg",
   Numeric: 4,
   MinLen: 4,
   MaxLen: 20,
   Phrase: true,
   CaseSensitive: true,
   Lang: "en",
   HintImgFile: "/path/to/hint.jpg",
   HintText: "Type red symbols",
}
```

### Text Captcha

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_text_captcha)</sup>

This method can be used to bypass a captcha that requires to answer a question provided in clear text.

```go
captcha:= api2captcha.Text{
   Text: "If tomorrow is Saturday, what day is today?",
   Lang: "en",
}
```

### reCAPTCHA v2

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_recaptchav2_new)</sup>

Use this method to solve reCAPTCHA V2 and obtain a token to bypass the protection.

```go
captcha := api2captcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Invisible: true,
   Action: "verify",
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### reCAPTCHA v3

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_recaptchav3)</sup>

This method provides reCAPTCHA V3 solver and returns a token.

```go
captcha := api2captcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Version: "v3",
   Action: "verify",
   Score: 0.3,
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### reCAPTCHA Enterprise

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_recaptcha_enterprise)</sup>

reCAPTCHA Enterprise can be used as reCAPTCHA V2 and reCAPTCHA V3. Below is a usage example for both versions.

```go
// reCAPTCHA V2
captcha:=  api2captcha.ReCaptcha({
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Invisible: true,
   Action: "verify",
   Enterprise: true,
})

// reCAPTCHA V3
captcha := api2captcha.ReCaptcha{
   SiteKey: "6Le-wvkSVVABCPBMRTvw0Q4Muexq1bi0DJwx_mJ-",
   Url: "https://mysite.com/page/with/recaptcha",
   Version: "v3",
   Action: "verify",
   Score: 0.3,
   Enterprise: true,
}

req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### FunCaptcha

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_funcaptcha_new)</sup>

FunCaptcha (Arkoselabs) solving method. Returns a token.

```go
captcha := api2captcha.FunCaptcha{
   SiteKey: "69A21A01-CC7B-B9C6-0F9A-E7FA06677FFC",
   Url: "https://mysite.com/page/with/funcaptcha",
   Surl: "https://client-api.arkoselabs.com",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
   Data: map[string]string{"anyKey":"anyValue"},
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### GeeTest

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_geetest)</sup>

Method to solve GeeTest puzzle captcha. Returns a set of tokens as JSON.

```go
captcha := api2captcha.GeeTest{
   GT: "f2ae6cadcf7886856696502e1d55e00c",
   ApiServer: "api-na.geetest.com",
   Challenge: "12345678abc90123d45678ef90123a456b",
   Url: "https://mysite.com/captcha.html",
}
req := captcha.ToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### GeeTest V4

<sup>[API method description.](https://2captcha.com/2captcha-api#geetest-v4)</sup>

Use this method to solve GeeTest v4. Returns the response in JSON.

```go
captcha:= api2captcha.GeeTestV4{
    CaptchaId: "e392e1d7fd421dc63325744d5a2b9c73",
    Url: "https://www.site.com/page/",
}
```

### KeyCaptcha

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_keycaptcha)</sup>

Token-based method to solve KeyCaptcha.

```go
captcha:= api2captcha.KeyCaptcha{
   UserId: 10,
   SessionId: "493e52c37c10c2bcdf4a00cbc9ccd1e8",
   WebServerSign: "9006dc725760858e4c0715b835472f22",
   WebServerSign2: "9006dc725760858e4c0715b835472f22",
   Url: "https://www.keycaptcha.ru/demo-magnetic/",
}
req := captchaToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### Capy

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_capy)</sup>

Token-based method to bypass Capy puzzle captcha.

```go
captcha:= api2captcha.Capy{
   SiteKey: "PUZZLE_Abc1dEFghIJKLM2no34P56q7rStu8v",
   Url: "https://www.mysite.com/captcha/",
}
req := captchaToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### Grid

<sup>[API method description.](https://2captcha.com/2captcha-api#grid)</sup>

Grid method is originally called Old reCAPTCHA V2 method. The method can be used to bypass any type of captcha where you can apply a grid on image and need to click specific grid boxes. Returns numbers of boxes.

```go
captcha:= api2captcha.Grid{
    File: "path/to/captcha.jpg",
    Rows: 3,
    Cols: 3,
    PreviousId: 0,
    CanSkip: false,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Select all images with an Orange",
}
```

### Canvas

<sup>[API method description.](https://2captcha.com/2captcha-api#canvas)</sup>

Canvas method can be used when you need to draw a line around an object on image. Returns a set of points' coordinates to draw a polygon.

```go
captcha:= api2captcha.Canvas{
    File: "path/to/captcha.jpg",
    PreviousId: 0,
    CanSkip: false,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Draw around apple",
}
```

### ClickCaptcha

<sup>[API method description.](https://2captcha.com/2captcha-api#coordinates)</sup>

ClickCaptcha method returns coordinates of points on captcha image. Can be used if you need to click on particular points on the image.

```go
captcha:= api2captcha.Coordinates{
    File: "path/to/captcha.jpg",
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Connect the dots",
}
```

### Rotate

<sup>[API method description.](https://2captcha.com/2captcha-api#solving_rotatecaptcha)</sup>

This method can be used to solve a captcha that asks to rotate an object. Mostly used to bypass FunCaptcha. Returns the rotation angle.

```go
captcha:= api2captcha.Rotate{
    File: "path/to/captcha.jpg",
    Angle: 15,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Put the images in the correct way",
}
```

### Amazon WAF

<sup>[API method description.](https://2captcha.com/2captcha-api#amazon-waf)</sup>

Use this method to solve Amazon WAF Captcha also known as AWS WAF Captcha is a part of Intelligent threat mitigation for Amazon AWS. Returns JSON with the token.

```go
captcha:= api2captcha.AmazonWAF {
   Iv: "CgAHbCe2GgAAAAAj",
   SiteKey: "0x1AAAAAAAAkg0s2VIOD34y5",
   Url: "https://non-existent-example.execute-api.us-east-1.amazonaws.com/latest",
   Context: "9BUgmlm48F92WUoqv97a49ZuEJJ50TCk9MVr3C7WMtQ0X6flVbufM4n8mjFLmbLVAPgaQ1Jydeaja94iAS49ljb",
   ChallengeScript: "https://41bcdd4fb3cb.610cd090.us-east-1.token.awswaf.com/41bcdd4fb3cb/0d21de737ccb/cd77baa6c832/challenge.js"
   CaptchaScript: "https://41bcdd4fb3cb.610cd090.us-east-1.captcha.awswaf.com/41bcdd4fb3cb/0d21de737ccb/cd77baa6c832/captcha.js"
}
```

### Cloudflare Turnstile

<sup>[API method description.](https://2captcha.com/2captcha-api#turnstile)</sup>

Use this method to solve Cloudflare Turnstile. Returns JSON with the token.

```go
captcha:= api2captcha.CloudflareTurnstile{
   SiteKey: "0x1AAAAAAAAkg0s2VIOD34y5",
   Url: "http://mysite.com/",
}
```

### Lemin Cropped Captcha

<sup>[API method description.](https://2captcha.com/2captcha-api#lemin)</sup>

Use this method to solve Lemin Captcha challenge. Returns JSON with answer containing the following values: answer, challenge_id.

```go
captcha:= Lemin{
   CaptchaId: "CROPPED_3dfdd5c_d1872b526b794d83ba3b365eb15a200b",
   Url:   "https://www.site.com/page/",
   DivId:     "lemin-cropped-captcha",
   ApiServer: "api.leminnow.com",
}
```

### CyberSiARA

<sup>[API method description.](https://2captcha.com/2captcha-api#cybersiara)</sup>

Use this method to solve CyberSiARA and obtain a token to bypass the protection.
```go
captcha:= api2captcha.CyberSiARA{
   MasterUrlId: "12333-3123123",
   Url: "https://test.com",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}
```

### DataDome

<sup>[API method description.](https://2captcha.com/2captcha-api#datadome)</sup>

Use this method to solve DataDome and obtain a token to bypass the protection.

> [!IMPORTANT]
> To solve the DataDome captcha, you must use a proxy. It is recommended to use [residential proxies].


```go
captcha:= api2captcha.DataDome{
   Url: "https://test.com",
   CaptchaUrl: "https://test.com/captcha/",
   Proxytype: "http",
   Proxy: "proxyuser:strongPassword@123.123.123.123:3128",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}
```

### MTCaptcha

<sup>[API method description.](https://2captcha.com/2captcha-api#mtcaptcha)</sup>

Use this method to solve MTCaptcha and obtain a token to bypass the protection.
```go
captcha:= api2captcha.MTCaptcha{
   Url: "https://service.mtcaptcha.com/mtcv1/demo/index.html",
   SiteKey: "MTPublic-DemoKey9M",
}
```

### Yandex
Use this method to solve Yandex and obtain a token to bypass the protection.
```go
captcha:= api2captcha.Yandex{
   Url: "https://example.com",
   SiteKey: "Y5Lh0tiycconMJGsFd3EbbuNKSp1yaZESUOIHfeV",
}
```

### Tencent

<sup>[API method description.](https://2captcha.com/2captcha-api#tencent)</sup>

Use this method to solve Tencent and obtain a token to bypass the protection.
```go
tencentCaptcha := api2captcha.Tencent{
   AppId: "2092215077",
   Url:   "http://lcec.lclog.cn/cargo/NewCargotracking?blno=BANR01XMHB0004&selectstate=BLNO",
}
```

### atbCAPTCHA

<sup>[API method description.](https://2captcha.com/2captcha-api#atb-captcha)</sup>

Use this method to solve atbCAPTCHA and obtain a token to bypass the protection.
```go
atbCaptcha := api2captcha.AtbCAPTCHA{
   AppId:     "af23e041b22d000a11e22a230fa8991c",
   Url:       "https://www.playzone.vip/",
   ApiServer: "https://cap.aisecurius.com",
}
```

### Cutcaptcha

<sup>[API method description.](https://2captcha.com/2captcha-api#cutcaptcha)</sup>

Use this method to solve Cutcaptcha and obtain a token to bypass the protection.
```go
captcha:= api2captcha.Cutcaptcha{
   MiseryKey: "a1488b66da00bf332a1488993a5443c79047e752",
   DataApiKey: "SAb83IIB",
   Url: "https://example.cc/foo/bar.html",
}
```

### Friendly Captcha

<sup>[API method description.](https://2captcha.com/2captcha-api#friendly-captcha)</sup>

Use this method to solve Friendly Captcha and obtain a token to bypass the protection.
```go
captcha:= api2captcha.Friendly{
   Url: "https://example.com",
   SiteKey: "2FZFEVS1FZCGQ9",
}
```

### Audio Captcha

<sup>[API method description.](https://2captcha.com/2captcha-api#audio)</sup>

Use this method to solve Audio captcha and obtain a token to bypass the protection.
```go
audio := api2captcha.Audio{
   Base64: fileBase64Str,
   Lang:   "en",
}
```

### Prosopo

<sup>[API method description.](https://2captcha.com/2captcha-api#prosopo-procaptcha)</sup>

Use this method to solve Prosopo Captcha and obtain a token to bypass the protection.
```go
	prosopo := api2captcha.Prosopo{
		SiteKey: "5EZVvsHMrKCFKp5NYNoTyDjTjetoVo1Z4UNNbTwJf1GfN6Xm",
		Url:     "https://www.twickets.live/",
	}
```

### Captchafox

<sup>[API method description.](https://2captcha.com/2captcha-api#captchafox)</sup>

Use this method to bypass Captchafox.

```go
	captchafox := api2captcha.Captchafox{
		SiteKey:   "sk_ILKWNruBBVKDOM7dZs59KHnDLEWiH",
		Url:       "https://mysite.com/page/with/captchafox",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		Proxy:     "username:password@1.2.3.4:5678",
		Proxytype: "http",
	}
```


## Other methods

### Send / GetResult
These methods can be used for manual captcha submission and answer polling.

```go
id, err := client.Send(captchaToRequest())
if err != nil {
   log.Fatal(err);
}

time.Sleep(10 * time.Second)

code, err := client.GetResult(id)
if err != nil {
   log.Fatal(err);
}

if code == nil {
   log.Fatal("Not ready")
}

fmt.Println("code "+*code)
```

### balance

<sup>[API method description.](https://2captcha.com/2captcha-api#additional-methods)</sup>

Use this method to get your account's balance

```go
balance, err := client.GetBalance()
if err != nil {
   log.Fatal(err);
}
```
### report

<sup>[API method description.](https://2captcha.com/2captcha-api#complain)</sup>

Use this method to report good or bad captcha answer.

```go
err := client.Report(id, true) // solved correctly
err := client.Report(id, false) // solved incorrectly
```

## Error Handling

When using the 2Captcha API, it's important to handle errors properly to ensure smooth interaction with the service. Below is an example of error handling in the `2captcha-go` library:

```go
func main() {
	result, err := client.Text("If tomorrow is Saturday, what day is today?")
	if err != nil {
		switch e := err.(type) {
		case *api.ValidationException:
			// Invalid parameters passed
			log.Fatalf("Validation error: %v", e)
		case *api.NetworkException:
			// Network error occurred
			log.Fatalf("Network error: %v", e)
		case *api.ApiException:
			// API responded with an error
			log.Fatalf("API error: %v", e)
		case *api.TimeoutException:
			// CAPTCHA is not solved within the expected time
			log.Fatalf("Timeout error: %v", e)
		default:
			// Unknown error
			log.Fatalf("Unexpected error: %v", err)
		}
	}

	fmt.Printf("CAPTCHA solved! Result: %s\n", result.Text)
}
```

## Proxies
You can pass your proxy as an additional argument for methods: recaptcha, funcaptcha, geetest, geetest v4, keycaptcha, capy puzzle, lemin, turnstile, amazon waf, CyberSiARA, DataDome, MTCaptcha and etc. The proxy will be forwarded to the API to solve the captcha.

We have our own proxies that we can offer you. [Buy residential proxies](https://2captcha.com/proxy/residential-proxies) for avoid restrictions and blocks. [Quick start](https://2captcha.com/proxy?openAddTrafficModal=true).

## Examples

Examples of solving all supported captcha types are located in the [examples] directory.

## Get in touch

<a href="mailto:support@2captcha.com"><img src="https://github.com/user-attachments/assets/539df209-7c85-4fa5-84b4-fc22ab93fac7" width="80" height="30"></a>
<a href="https://2captcha.com/support/tickets/new"><img src="https://github.com/user-attachments/assets/be044db5-2e67-46c6-8c81-04b78bd99650" width="81" height="30"></a>

## Join the team 👪

There are many ways to contribute, of which development is only one! Find your next job. Open positions: AI experts, scrapers, developers, technical support, and much more! 😍

<a href="mailto:job@2captcha.com"><img src="https://github.com/user-attachments/assets/36d23ef5-7866-4841-8e17-261cc8a4e033" width="80" height="30"></a>

## License

The code in this repository is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.

### Graphics and Trademarks

The graphics and trademarks included in this repository are not covered by the MIT License. Please contact <a href="mailto:support@2captcha.com">support</a> for permissions regarding the use of these materials.


<!-- Shared links -->
[2Captcha]: https://2captcha.com/
[2captcha software catalog]: https://2captcha.com/software
[pingback settings]: https://2captcha.com/setting/pingback
[post options]: https://2captcha.com/2captcha-api#normal_post
[list of supported languages]: https://2captcha.com/2captcha-api#language
[residential proxies]: https://2captcha.com/proxy/residential-proxies
[examples]: ./examples/
