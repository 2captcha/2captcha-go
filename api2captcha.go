package api2captcha

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	BaseURL       = "https://2captcha.com"
	DefaultSoftId = 4583
)

type (
	Request struct {
		Params map[string]string
		Files  map[string]string
	}

	Client struct {
		BaseURL          *url.URL
		ApiKey           string
		SoftId           int
		Callback         string
		DefaultTimeout   int
		RecaptchaTimeout int
		PollingInterval  int

		httpClient *http.Client
	}

	Canvas struct {
		File            string
		Base64          string
		PreviousId      int
		CanSkip         bool
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	Capy struct {
		SiteKey   string
		Url       string
		ApiServer string
	}

	Coordinates struct {
		File            string
		Base64          string
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	FunCaptcha struct {
		SiteKey   string
		Url       string
		Surl      string
		UserAgent string
		Data      map[string]string
	}

	GeeTest struct {
		GT        string
		Challenge string
		Url       string
		ApiServer string
	}

	Grid struct {
		File            string
		Base64          string
		Rows            int
		Cols            int
		PreviousId      int
		CanSkip         bool
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	HCaptcha struct {
		SiteKey string
		Url     string
	}

	KeyCaptcha struct {
		UserId         int
		SessionId      string
		WebServerSign  string
		WebServerSign2 string
		Url            string
	}

	Normal struct {
		File            string
		Base64          string
		Phrase          bool
		CaseSensitive   bool
		Calc            bool
		Numberic        int
		MinLen          int
		MaxLen          int
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	ReCaptcha struct {
		SiteKey    string
		Url        string
		Invisible  bool
		Enterprise bool
		Version    string
		Action     string
		DataS      string
		Score      float64
		UserAgent  string
		Cookies    string
	}

	Rotate struct {
		Base64          string
		File            string
		Files           []string
		Angle           int
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	Text struct {
		Text string
		Lang string
	}

	AmazonWAF struct {
		Iv              string
		SiteKey         string
		Url             string
		Context         string
		ChallengeScript string
		CaptchaScript   string
	}

	GeeTestV4 struct {
		CaptchaId string
		Url       string
		ApiServer string
		Challenge string
	}

	Lemin struct {
		CaptchaId string
		DivId     string
		Url       string
		ApiServer string
	}

	CloudflareTurnstile struct {
		SiteKey   string
		Url       string
		Data      string
		PageData  string
		Action    string
		UserAgent string
	}

	CyberSiARA struct {
		MasterUrlId string
		Url         string
		UserAgent   string
	}

	DataDome struct {
		Url        string
		CaptchaUrl string
		Proxytype  string
		Proxy      string
		UserAgent  string
	}

	MTCaptcha struct {
		SiteKey string
		Url     string
	}

	Yandex struct {
		Url     string
		SiteKey string
	}

	Friendly struct {
		Url     string
		SiteKey string
	}

	CutCaptcha struct {
		MiseryKey  string
		DataApiKey string
		Url        string
	}

	Tencent struct {
		AppId string
		Url   string
	}

	AtbCAPTCHA struct {
		AppId     string
		ApiServer string
		Url       string
	}

	Audio struct {
		Base64 string
		Lang   string
	}

	Prosopo struct {
		Url     string
		SiteKey string
	}

	Captchafox struct {
		Url       string
		SiteKey   string
		Proxytype string
		Proxy     string
		UserAgent string
	}

	Temu struct {
		Body  string
		Part1 string
		Part2 string
		Part3 string
	}
)

var (
	ErrNetwork = errors.New("api2captcha: Network failure")
	ErrApi     = errors.New("api2captcha: API error")
	ErrTimeout = errors.New("api2captcha: Request timeout")
)

func NewClient(apiKey string) *Client {
	base, _ := url.Parse(BaseURL)
	return &Client{
		BaseURL:          base,
		ApiKey:           apiKey,
		SoftId:           DefaultSoftId,
		DefaultTimeout:   120,
		PollingInterval:  10,
		RecaptchaTimeout: 600,
		httpClient:       http.DefaultClient,
	}
}

func NewClientExt(apiKey string, client *http.Client) *Client {
	base, _ := url.Parse(BaseURL)
	return &Client{
		BaseURL:          base,
		ApiKey:           apiKey,
		DefaultTimeout:   120,
		PollingInterval:  10,
		RecaptchaTimeout: 600,
		httpClient:       client,
	}
}

func (c *Client) res(req Request) (*string, error) {
	rel := &url.URL{Path: "/res.php"}
	uri := c.BaseURL.ResolveReference(rel)

	req.Params["key"] = c.ApiKey
	c.httpClient.Timeout = time.Duration(c.DefaultTimeout) * time.Second

	var resp *http.Response = nil

	values := url.Values{}
	for key, val := range req.Params {
		values.Add(key, val)
	}
	uri.RawQuery = values.Encode()

	var err error = nil
	resp, err = c.httpClient.Get(uri.String())
	if err != nil {
		return nil, ErrNetwork
	}

	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	data := body.String()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrApi
	}

	if strings.HasPrefix(data, "ERROR_") {
		return nil, ErrApi
	}

	return &data, nil
}

func (c *Client) resAction(action string) (*string, error) {
	req := Request{
		Params: map[string]string{"action": action},
	}

	return c.res(req)
}

func (c *Client) Send(req Request) (string, error) {
	rel := &url.URL{Path: "/in.php"}
	uri := c.BaseURL.ResolveReference(rel)

	req.Params["key"] = c.ApiKey

	c.httpClient.Timeout = time.Duration(c.DefaultTimeout) * time.Second

	var resp *http.Response = nil
	if req.Files != nil && len(req.Files) > 0 {

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for name, path := range req.Files {
			file, err := os.Open(path)
			if err != nil {
				return "", err
			}
			defer file.Close()

			part, err := writer.CreateFormFile(name, filepath.Base(path))
			if err != nil {
				return "", err
			}
			_, err = io.Copy(part, file)
		}

		for key, val := range req.Params {
			_ = writer.WriteField(key, val)
		}

		err := writer.Close()
		if err != nil {
			return "", err
		}

		request, err := http.NewRequest("POST", uri.String(), body)
		if err != nil {
			return "", err
		}

		request.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err = c.httpClient.Do(request)
		if err != nil {
			return "", ErrNetwork
		}
	} else {
		values := url.Values{}
		for key, val := range req.Params {
			values.Add(key, val)
		}

		var err error = nil
		resp, err = c.httpClient.PostForm(uri.String(), values)
		if err != nil {
			return "", ErrNetwork
		}
	}

	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err := body.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	data := body.String()

	if resp.StatusCode != http.StatusOK {
		return "", ErrApi
	}

	if strings.HasPrefix(data, "ERROR_") {
		return "", ErrApi
	}

	if !strings.HasPrefix(data, "OK|") {
		return "", ErrApi
	}

	return data[3:], nil
}

func (c *Client) Solve(req Request) (string, string, error) {
	if c.Callback != "" {
		_, ok := req.Params["pingback"]
		if !ok {
			// set default pingback
			req.Params["pingback"] = c.Callback
		}
	}

	pingback, hasPingback := req.Params["pingback"]
	if pingback == "" {
		delete(req.Params, "pingback")
		hasPingback = false
	}

	_, ok := req.Params["soft_id"]
	if c.SoftId != 0 && !ok {
		req.Params["soft_id"] = strconv.FormatInt(int64(c.SoftId), 10)
	}

	id, err := c.Send(req)
	if err != nil {
		return "", "", err
	}

	// don't wait for result if Callback is used
	if hasPingback {
		return "", id, nil
	}

	timeout := c.DefaultTimeout
	if req.Params["method"] == "userrecaptcha" {
		timeout = c.RecaptchaTimeout
	}

	token, err := c.WaitForResult(id, timeout, c.PollingInterval)
	if err != nil {
		return "", "", err
	}

	return token, id, nil
}

func (c *Client) WaitForResult(id string, timeout int, interval int) (string, error) {
	start := time.Now()
	now := start
	for now.Sub(start) < (time.Duration(timeout) * time.Second) {

		time.Sleep(time.Duration(interval) * time.Second)

		code, err := c.GetResult(id)
		if err == nil && code != nil {
			return *code, nil
		}

		// ignore network errors
		if err != nil && err != ErrNetwork {
			return "", err
		}

		now = time.Now()
	}

	return "", ErrTimeout
}

func (c *Client) GetResult(id string) (*string, error) {
	req := Request{
		Params: map[string]string{"action": "get", "id": id},
	}

	data, err := c.res(req)
	if err != nil {
		return nil, err
	}

	if *data == "CAPCHA_NOT_READY" {
		return nil, nil
	}

	if !strings.HasPrefix(*data, "OK|") {
		return nil, ErrApi
	}

	reply := (*data)[3:]
	return &reply, nil
}

func (c *Client) GetBalance() (float64, error) {
	data, err := c.resAction("getbalance")
	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(*data, 64)
}

func (c *Client) Report(id string, correct bool) error {
	req := Request{
		Params: map[string]string{"id": id},
	}
	if correct {
		req.Params["action"] = "reportgood"
	} else {
		req.Params["action"] = "reportbad"
	}

	_, err := c.res(req)
	return err
}

func (req *Request) SetProxy(proxyType string, uri string) {
	req.Params["proxytype"] = proxyType
	req.Params["proxy"] = uri
}

func (req *Request) SetSoftId(softId int) {
	req.Params["soft_id"] = strconv.FormatInt(int64(softId), 10)
}

func (req *Request) SetCallback(callback string) {
	req.Params["pingback"] = callback
}

func (c *Canvas) ToRequest() Request {
	req := Request{
		Params: map[string]string{"canvas": "1", "recaptcha": "1"},
		Files:  map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
		req.Params["method"] = "post"
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
		req.Params["method"] = "base64"
	}
	if c.PreviousId != 0 {
		req.Params["previousID"] = strconv.FormatInt(int64(c.PreviousId), 10)
	}
	if c.CanSkip {
		req.Params["can_no_answer"] = "1"
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}
	if c.HintText != "" {
		req.Params["textinstructions"] = c.HintText
	}
	if c.HintImageBase64 != "" {
		req.Params["imginstructions"] = c.HintImageBase64
	}
	if c.HintImageFile != "" {
		req.Files["imginstructions"] = c.HintImageFile
	}

	return req
}

func (c *Normal) ToRequest() Request {
	req := Request{
		Params: map[string]string{},
		Files:  map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
		req.Params["method"] = "post"
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
		req.Params["method"] = "base64"
	}

	if c.Phrase {
		req.Params["phrase"] = "1"
	}
	if c.CaseSensitive {
		req.Params["regsense"] = "1"
	}
	if c.Calc {
		req.Params["calc"] = "1"
	}
	if c.Numberic != 0 {
		req.Params["numeric"] = strconv.FormatInt(int64(c.Numberic), 10)
	}
	if c.MinLen != 0 {
		req.Params["min_len"] = strconv.FormatInt(int64(c.MinLen), 10)
	}
	if c.MaxLen != 0 {
		req.Params["max_len"] = strconv.FormatInt(int64(c.MaxLen), 10)
	}

	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}
	if c.HintText != "" {
		req.Params["textinstructions"] = c.HintText
	}
	if c.HintImageBase64 != "" {
		req.Params["imginstructions"] = c.HintImageBase64
	}
	if c.HintImageFile != "" {
		req.Files["imginstructions"] = c.HintImageFile
	}

	return req
}

func (c *Capy) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "capy"},
	}
	if c.SiteKey != "" {
		req.Params["captchakey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.ApiServer != "" {
		req.Params["api_server"] = c.ApiServer
	}

	return req
}

func (c *Coordinates) ToRequest() Request {
	req := Request{
		Params: map[string]string{"coordinatescaptcha": "1"},
		Files:  map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}
	if c.HintText != "" {
		req.Params["textinstructions"] = c.HintText
	}
	if c.HintImageBase64 != "" {
		req.Params["imginstructions"] = c.HintImageBase64
	}
	if c.HintImageFile != "" {
		req.Files["imginstructions"] = c.HintImageFile
	}

	return req
}

func (c *FunCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "funcaptcha"},
	}
	if c.SiteKey != "" {
		req.Params["publickey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.Surl != "" {
		req.Params["surl"] = c.Surl
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}
	if c.Data != nil {
		for key, value := range c.Data {
			param := "data[" + key + "]"
			req.Params[param] = value
		}
	}

	return req
}

func (c *GeeTest) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "geetest"},
	}
	if c.GT != "" {
		req.Params["gt"] = c.GT
	}
	if c.Challenge != "" {
		req.Params["challenge"] = c.Challenge
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.ApiServer != "" {
		req.Params["api_server"] = c.ApiServer
	}

	return req
}

func (c *Grid) ToRequest() Request {
	req := Request{
		Params: map[string]string{},
		Files:  map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
	}
	if c.Rows != 0 {
		req.Params["recaptcharows"] = strconv.FormatInt(int64(c.Rows), 10)
	}
	if c.Cols != 0 {
		req.Params["recaptchacols"] = strconv.FormatInt(int64(c.Cols), 10)
	}
	if c.PreviousId != 0 {
		req.Params["previousID"] = strconv.FormatInt(int64(c.PreviousId), 10)
	}
	if c.CanSkip {
		req.Params["can_no_answer"] = "1"
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}
	if c.HintText != "" {
		req.Params["textinstructions"] = c.HintText
	}
	if c.HintImageBase64 != "" {
		req.Params["imginstructions"] = c.HintImageBase64
	}
	if c.HintImageFile != "" {
		req.Files["imginstructions"] = c.HintImageFile
	}

	return req
}

func (c *HCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "hcaptcha"},
	}
	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *KeyCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "keycaptcha"},
	}
	if c.UserId != 0 {
		req.Params["s_s_c_user_id"] = strconv.FormatInt(int64(c.UserId), 10)
	}
	if c.SessionId != "" {
		req.Params["s_s_c_session_id"] = c.SessionId
	}
	if c.WebServerSign != "" {
		req.Params["s_s_c_web_server_sign"] = c.WebServerSign
	}
	if c.WebServerSign2 != "" {
		req.Params["s_s_c_web_server_sign2"] = c.WebServerSign2
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *ReCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "userrecaptcha"},
	}
	if c.SiteKey != "" {
		req.Params["googlekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.Invisible {
		req.Params["invisible"] = "1"
	}
	if c.Enterprise {
		req.Params["enterprise"] = "1"
	}
	if c.Version != "" {
		req.Params["version"] = c.Version
	}
	if c.Action != "" {
		req.Params["action"] = c.Action
	}
	if c.DataS != "" {
		req.Params["data-s"] = c.DataS
	}
	if c.Score != 0 {
		req.Params["min_score"] = strconv.FormatFloat(c.Score, 'f', -1, 64)
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}
	if c.Cookies != "" {
		req.Params["cookies"] = c.Cookies
	}

	return req
}

func (c *Rotate) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "rotatecaptcha"},
		Files:  map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
	}
	if c.Files != nil {
		for i := 0; i < len(c.Files); i++ {
			name := "file_" + strconv.FormatInt(int64(i)+1, 10)
			req.Files[name] = c.Files[i]
		}
	}
	if c.Angle != 0 {
		req.Params["angle"] = strconv.FormatInt(int64(c.Angle), 10)
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}
	if c.HintText != "" {
		req.Params["textinstructions"] = c.HintText
	}
	if c.HintImageBase64 != "" {
		req.Params["imginstructions"] = c.HintImageBase64
	}
	if c.HintImageFile != "" {
		req.Files["imginstructions"] = c.HintImageFile
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
	}
	return req
}

func (c *Text) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "post"},
	}
	if c.Text != "" {
		req.Params["textcaptcha"] = c.Text
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}

	return req
}

func (c *AmazonWAF) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "amazon_waf"},
	}

	if c.Iv != "" {
		req.Params["iv"] = c.Iv
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}

	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	if c.Context != "" {
		req.Params["context"] = c.Context
	}

	if c.ChallengeScript != "" {
		req.Params["challenge_script"] = c.ChallengeScript
	}

	if c.CaptchaScript != "" {
		req.Params["captcha_script"] = c.CaptchaScript
	}

	return req
}

func (c *GeeTestV4) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "geetest_v4"},
	}
	if c.CaptchaId != "" {
		req.Params["captcha_id"] = c.CaptchaId
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *Lemin) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "lemin"},
	}

	if c.CaptchaId != "" {
		req.Params["captcha_id"] = c.CaptchaId
	}

	if c.DivId != "" {
		req.Params["div_id"] = c.DivId
	} else {
		req.Params["div_id"] = "lemin-cropped-captcha"
	}

	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	if c.ApiServer != "" {
		req.Params["api_server"] = c.ApiServer
	}
	return req
}

func (c *CloudflareTurnstile) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "turnstile"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.Data != "" {
		req.Params["data"] = c.Data
	}
	if c.PageData != "" {
		req.Params["pagedata"] = c.PageData
	}
	if c.Action != "" {
		req.Params["action"] = c.Action
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}

	return req
}

func (c *CyberSiARA) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "cybersiara"},
	}

	if c.MasterUrlId != "" {
		req.Params["master_url_id"] = c.MasterUrlId
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}

	return req
}

func (c *DataDome) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "datadome"},
	}

	if c.CaptchaUrl != "" {
		req.Params["captcha_url"] = c.CaptchaUrl
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.Proxytype != "" {
		req.Params["proxytype"] = c.Proxytype
	}
	if c.Proxy != "" {
		req.Params["proxy"] = c.Proxy
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}

	return req
}

func (c *MTCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "mt_captcha"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *Yandex) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "yandex"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *Friendly) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "friendly_captcha"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *Tencent) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "tencent"},
	}
	if c.AppId != "" {
		req.Params["app_id"] = c.AppId
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *AtbCAPTCHA) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "atb_captcha"},
	}
	if c.AppId != "" {
		req.Params["app_id"] = c.AppId
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.ApiServer != "" {
		req.Params["api_server"] = c.ApiServer
	}

	return req
}

func (c *CutCaptcha) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "cutcaptcha"},
	}
	if c.MiseryKey != "" {
		req.Params["misery_key"] = c.MiseryKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.DataApiKey != "" {
		req.Params["api_key"] = c.DataApiKey
	}

	return req
}

func (c *Audio) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "audio"},
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}

	return req
}

func (c *Prosopo) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "prosopo"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}

	return req
}

func (c *Captchafox) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "captchafox"},
	}

	if c.SiteKey != "" {
		req.Params["sitekey"] = c.SiteKey
	}
	if c.Url != "" {
		req.Params["pageurl"] = c.Url
	}
	if c.Proxytype != "" {
		req.Params["proxytype"] = c.Proxytype
	}
	if c.Proxy != "" {
		req.Params["proxy"] = c.Proxy
	}
	if c.UserAgent != "" {
		req.Params["userAgent"] = c.UserAgent
	}

	return req
}

func (c *Temu) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "temuimage"},
	}
	if c.Body != "" {
		req.Params["body"] = c.Body
	}
	if c.Part1 != "" {
		req.Params["part1"] = c.Part1
	}
	if c.Part2 != "" {
		req.Params["part2"] = c.Part2
	}
	if c.Part3 != "" {
		req.Params["part3"] = c.Part3
	}

	return req
}
