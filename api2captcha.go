package api2captcha

import (
	"bytes"
	"errors"
	"io"
	"log"
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
	BaseURL = "https://2captcha.com"
)

type (

	Request struct {
		Params map[string]string
		Files map[string]string
	}

	Client struct {
		BaseURL *url.URL
		ApiKey string
		SoftId int
		Callback string
		DefaultTimeout int
		RecaptchaTimeout int
		PollingInterval int
		
		httpClient *http.Client
	}

	Canvas struct {
		File string
		Base64 string
		PreviousId int
		CanSkip bool
		Lang string
		HintText string
		HintImageBase64 string
		HintImageFile string
	}

	Capy struct {
		SiteKey string
		Url string
		ApiServer string
	}

	Coordinates struct {
		File string
		Base64 string
		Lang string
		HintText string
		HintImageBase64 string
		HintImageFile string
	}

	FunCaptcha struct {
		SiteKey string
		Url string
		Surl string
		UserAgent string
		Data map[string]string
	}

	GeeTest struct {
		GT string
		Challenge string
		Url string
		ApiServer string
	}

	Grid struct {
		File string
		Base64 string
		Rows int
		Cols int
		PreviousId int
		CanSkip bool
		Lang string
		HintText string
		HintImageBase64 string
		HintImageFile string
	}

	HCaptcha struct {
		SiteKey string
		Url string
	}

	KeyCaptcha struct {
		UserId int
		SessionId string
		WebServerSign string
		WebServerSign2 string
		Url string
	}

	Normal struct {
		File string
		Base64 string
		Phrase bool
		CaseSensitive bool
		Calc bool
		Numberic int
		MinLen int
		MaxLen int		
		Lang string
		HintText string
		HintImageBase64 string
		HintImageFile string
	}

	ReCaptcha struct {
		SiteKey string
		Url string
		Invisible bool
		Version string
		Action string
		Score float64
	}

	Rotate struct {
		File string
		Files []string
		Angle int
		Lang string
		HintText string
		HintImageBase64 string
		HintImageFile string
	}

	Text struct {
		Text string
		Lang string
	}
)

var (
	ErrNetwork = errors.New("api2captcha: Network failure")
	ErrApi = errors.New("api2captcha: API error")
	ErrTimeout = errors.New("api2captcha: Request timeout")
)

func NewClient(apiKey string) *Client {
	base, _ := url.Parse(BaseURL)
	return &Client{
		BaseURL: base,
		ApiKey:  apiKey,
		DefaultTimeout: 10,
		PollingInterval: 10,
		RecaptchaTimeout: 600,
		httpClient: &http.Client{},
	}
}

func NewClientExt(apiKey string, client *http.Client) *Client {
	base, _ := url.Parse(BaseURL)
	return &Client{
		BaseURL: base,
		ApiKey:  apiKey,
		DefaultTimeout: 10,
		PollingInterval: 10,
		RecaptchaTimeout: 600,
		httpClient: client,
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
	resp, err = http.Get(uri.String())
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
	log.Println("Status "+resp.Status+" data "+data)

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
		Params: map[string]string{"action":action},
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
		resp, err = http.PostForm(uri.String(), values)
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
	log.Println("Status "+resp.Status+" data "+data)

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

func (c *Client) Solve(req Request) (string, error) {

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
		req.Params["soft_id"] = strconv.FormatInt(int64(c.SoftId), 10);
	}

	id, err := c.Send(req)
	if err != nil {
		return "", err
	}

	// don't wait for result if Callback is used
	if hasPingback {
		return id, nil
	}

	timeout := c.DefaultTimeout	
	if req.Params["method"] == "userrecaptcha" {
		timeout = c.RecaptchaTimeout
	}

	return c.WaitForResult(id, timeout, c.PollingInterval)
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
		Params: map[string]string{"action":"get", "id": id},
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

func (c *Client) Report(id string, correct bool) (error) {
	req := Request{
		Params: map[string]string{"id":id},
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
	req.Params["soft_id"] = strconv.FormatInt(int64(softId), 64)
}

func (req *Request) SetCallback(callback string) {
	req.Params["pingback"] = callback
}

func (c *Canvas) ToRequest() Request {
	req := Request{
		Params: map[string]string{"canvas":"1", "recaptcha": "1"},
		Files: map[string]string{},
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
		req.Params["previousID"] = strconv.FormatInt(int64(c.PreviousId), 64)
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
		Files: map[string]string{},
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
		req.Params["numeric"] = strconv.FormatInt(int64(c.Numberic), 64)
	}
	if c.MinLen != 0 {
		req.Params["min_len"] = strconv.FormatInt(int64(c.MinLen), 64)
	}
	if c.MaxLen != 0 {
		req.Params["max_len"] = strconv.FormatInt(int64(c.MaxLen), 64)
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
		Params: map[string]string{"method":"capy"},
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
		Files: map[string]string{},
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
		Files: map[string]string{},
	}
	if c.File != "" {
		req.Files["file"] = c.File
	}
	if c.Base64 != "" {
		req.Params["body"] = c.Base64
	}
	if c.Rows != 0 {
		req.Params["recaptcharows"] = strconv.FormatInt(int64(c.Rows), 64)
	}
	if c.Cols != 0 {
		req.Params["recaptchacols"] = strconv.FormatInt(int64(c.Cols), 64)
	}
	if c.PreviousId != 0 {
		req.Params["previousID"] = strconv.FormatInt(int64(c.PreviousId), 64)
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
		req.Params["s_s_c_user_id"] = strconv.FormatInt(int64(c.UserId), 64)
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
	if c.Version != "" {
		req.Params["version"] = c.Version
	}
	if c.Action != "" {
		req.Params["action"] = c.Action
	}
	if c.Score != 0 {
		req.Params["min_score"] = strconv.FormatFloat(c.Score, 'f', -1, 64)
	}

	return req
}

func (c *Rotate) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method": "rotatecaptcha"},
		Files: map[string]string{},
	}
	if c.File != "" {
		req.Files["file_1"] = c.File
	}
	if c.Files != nil {
		for i := 0; i < len(c.Files); i++ {
			name := "file_" + strconv.FormatInt(int64(i) + 1, 64)
			req.Files[name] = c.Files[i]
		}
	}
	if c.Angle != 0 {
		req.Params["angle"] = strconv.FormatInt(int64(c.Angle), 64)
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

func (c *Text) ToRequest() Request {
	req := Request{
		Params: map[string]string{"method":"post"},
	}
	if c.Text != "" {
		req.Params["textcaptcha"] = c.Text
	}
	if c.Lang != "" {
		req.Params["lang"] = c.Lang
	}

	return req
}
