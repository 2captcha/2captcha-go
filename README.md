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
  - [Solve captcha](#solve-captcha)
    - [Normal Captcha](#normal-captcha)
    - [Text](#text-captcha)
    - [ReCaptcha v2](#recaptcha-v2)
    - [ReCaptcha v3](#recaptcha-v3)
    - [reCAPTCHA Enterprise](#recaptcha-enterprise)
    - [FunCaptcha](#funcaptcha)
    - [GeeTest](#geetest)
    - [hCaptcha](#hcaptcha)
    - [KeyCaptcha](#keycaptcha)
    - [Capy](#capy)
    - [Grid (ReCaptcha V2 Old Method)](#grid)
    - [Canvas](#canvas)
    - [ClickCaptcha](#clickcaptcha)
    - [Rotate](#rotate)
    - [AmazonWAF](#amazon-waf)
    - [CloudflareTurnstile](#cloudflare-turnstile)
    - [Lemin Cropped Captcha](#lemin-cropped-captcha)
    - [GeeTestV4](#geetestv4)
    - [CyberSiARA](#cybersiara)
    - [DataDome](#datadome)
    - [MTCaptcha](#mtcaptcha)
    - [Tencent](#tencent)
    - [AtbCAPTCHA](#atbcaptcha)
    - [Cutcaptcha](#cutcaptcha)
    - [FriendlyCaptcha](#friendlycaptcha)
  - [Other methods](#other-methods)
    - [send / getResult](#send--getresult)
    - [balance](#balance)
    - [report](#report)
  - [Proxies](#proxies)
- [Get in touch](#get-in-touch)
- [Join the team 👪](#join-the-team-)

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
|soft_id|4583|Your software ID obtained after publishing in [2captcha sofware catalog]|
|callback|-|URL of your web-sever that receives the captcha recognition result. The URl should be first registered in [pingback settings] of your account|
|default_timeout|120|Timeout in seconds for all captcha types except ReCaptcha. Defines how long the module tries to get the answer from `res.php` API endpoint|
|recaptcha_timeout|600|Timeout for ReCaptcha in seconds. Defines how long the module tries to get the answer from `res.php` API endpoint|
|polling_interval|10|Interval in seconds between requests to `res.php` API endpoint, setting values less than 5 seconds is not recommended|

>  **IMPORTANT:** once *callback URL* is defined for `client` instance, all methods return only the captcha ID and DO NOT poll the API to get the result. The result will be sent to the callback URL.
To get the answer manually use [GetResult method](#send--getresult)

## Solve captcha
When you submit any image-based captcha use can provide additional options to help 2captcha workers to solve it properly.

### Captcha options
|Option|Default Value|Description|
|---|---|---|
|numeric|0|Defines if captcha contains numeric or other symbols [see more info in the API docs][post options]|
|min_len|0|minimal answer lenght|
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
This method can be used to bypass a captcha that requires to answer a question provided in clear text.

```go
captcha:= api2captcha.Text{
   Text: "If tomorrow is Saturday, what day is today?",
   Lang: "en",
}
```

### ReCaptcha v2
Use this method to solve ReCaptcha V2 and obtain a token to bypass the protection.

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

### ReCaptcha v3
This method provides ReCaptcha V3 solver and returns a token.

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

### hCaptcha
Method to solve GeeTest puzzle captcha. Returns a set of tokens as JSON.

```go
captcha:= api2captcha.HCaptcha{
   SiteKey: "10000000-ffff-ffff-ffff-000000000001",
   Url: "https://mysite.com/captcha.html",
}
req := captchaToRequest()
req.SetProxy("HTTPS", "login:password@IP_address:PORT")
code, err := client.Solve(req)
```

### KeyCaptcha
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
Grid method is originally called Old ReCaptcha V2 method. The method can be used to bypass any type of captcha where you can apply a grid on image and need to click specific grid boxes. Returns numbers of boxes.

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
This method can be used to solve a captcha that asks to rotate an object. Mostly used to bypass FunCaptcha. Returns the rotation angle.

```go
captcha:= api2captcha.Rotate{
    File: "path/to/captcha.jpg",
    Angle: 40,
    Lang: "en",
    HintImageFile: "path/to/hint.jpg",
    HintText: "Put the images in the correct way",
}
```
### GeeTestV4
Use this method to solve GeeTest v4. Returns the response in JSON.

```go
captcha:= api2captcha.GeeTestV4{
    CaptchaId: "e392e1d7fd421dc63325744d5a2b9c73",
    Url: "https://www.site.com/page/",
}
```

### Lemin Cropped Captcha
Use this method to solve Lemin Captcha challenge. Returns JSON with answer containing the following values: answer, challenge_id.

```go
captcha:= Lemin{
   CaptchaId: "CROPPED_3dfdd5c_d1872b526b794d83ba3b365eb15a200b",
   Url:   "https://www.site.com/page/",
   DivId:     "lemin-cropped-captcha",
   ApiServer: "api.leminnow.com",
}
```

### Cloudflare Turnstile
Use this method to solve Cloudflare Turnstile. Returns JSON with the token.

```go
captcha:= api2captcha.CloudflareTurnstile{
   SiteKey: "0x1AAAAAAAAkg0s2VIOD34y5",
   Url: "http://mysite.com/",
}
```

### CyberSiARA
Use this method to solve CyberSiARA and obtain a token to bypass the protection.
```go
captcha:= api2captcha.CyberSiARA{
   MasterUrlId: "12333-3123123",
   Url: "https://test.com",
   UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
}
```

### DataDome
Use this method to solve DataDome and obtain a token to bypass the protection.
To solve the DataDome captcha, you must use a proxy.
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
  Url: "https://rutube.ru",
  SiteKey: "Y5Lh0tiycconMJGsFd3EbbuNKSp1yaZESUOIHfeV",
}
```

### Friendlycaptcha
Use this method to solve Friendly Captcha and obtain a token to bypass the protection.
```go
captcha:= api2captcha.Friendly{
  Url: "https://example.com",
  SiteKey: "2FZFEVS1FZCGQ9",
}
```

### Cutcaptcha
Use this method to solve Cutcaptcha and obtain a token to bypass the protection.
```go
captcha:= api2captcha.Cutcaptcha{
   MiseryKey: "a1488b66da00bf332a1488993a5443c79047e752",
   DataApiKey: "SAb83IIB",
   Url: "https://example.cc/foo/bar.html",
}
```

### Amazon WAF
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

### Tencent
Use this method to solve Tencent and obtain a token to bypass the protection.
```go
	tencentCaptcha := api2captcha.Tencent{
    AppId: "2092215077",
    Url:   "http://lcec.lclog.cn/cargo/NewCargotracking?blno=BANR01XMHB0004&selectstate=BLNO",
}
```

### AtbCAPTCHA
Use this method to solve AtbCAPTCHA and obtain a token to bypass the protection.
```go
	atbCaptcha := api2captcha.AtbCAPTCHA{
    AppId:     "af23e041b22d000a11e22a230fa8991c",
    Url:       "https://www.playzone.vip/",
    ApiServer: "https://cap.aisecurius.com",
}
```

### Cutcaptcha
Use this method to solve Cutcaptcha and obtain a token to bypass the protection.
```go
	cutCaptcha := api2captcha.Cutcaptcha{
    MiseryKey:  "a1488b66da00bf332a1488993a5443c79047e752",
    Url:        "https://filecrypt.co/Container/237D4D0995.html",
    DataApiKey: "SAb83IIB",
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
Use this method to get your account's balance

```go
balance, err := client.GetBalance()
if err != nil {
   log.Fatal(err);
}
```
### report
Use this method to report good or bad captcha answer.

```c++
err := client.Report(id, true) // solved correctly
err := client.Report(id, false) // solved incorrectly

```

## Proxies
You can pass your proxy as an additional argument for methods: recaptcha, funcaptcha, geetest, geetest v4, hcaptcha, keycaptcha, capy puzzle, lemin, turnstile, amazon waf, CyberSiARA, DataDome, MTCaptcha and etc. The proxy will be forwarded to the API to solve the captcha.

We have our own proxies that we can offer you. [Buy residential proxies](https://2captcha.com/proxy/residential-proxies) for avoid restrictions and blocks. [Quick start](https://2captcha.com/proxy?openAddTrafficModal=true).

## Get in touch

<a href="mailto:support@2captcha.com"><img src="https://github.com/user-attachments/assets/539df209-7c85-4fa5-84b4-fc22ab93fac7" width="80" height="30"></a>
<a href="https://2captcha.com/support/tickets/new"><img src="https://github.com/user-attachments/assets/be044db5-2e67-46c6-8c81-04b78bd99650" width="81" height="30"></a>

## Join the team 👪

There are many ways to contribute, of which development is only one! Find your next job. Open positions: AI experts, scrapers, developers, technical support, and much more! 😍

<a href="mailto:job@2captcha.com"><img src="https://github.com/user-attachments/assets/36d23ef5-7866-4841-8e17-261cc8a4e033" width="80" height="30"></a>

<!-- Shared links -->
[2Captcha]: https://2captcha.com/
[2captcha sofware catalog]: https://2captcha.com/software
[pingback settings]: https://2captcha.com/setting/pingback
[post options]: https://2captcha.com/2captcha-api#normal_post
[list of supported languages]: https://2captcha.com/2captcha-api#language
