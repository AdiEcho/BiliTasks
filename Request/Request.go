package Request

import (
	cp "BiliTasks/Cprint"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var client *http.Client

func DoRequest(req *http.Request) (string, error) {
	req.Close = true
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36 Edg/86.0.622.48")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		cp.Println("error", "Do Error: ", err)
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		cp.Println("error", "ReadAll Error: ", err)
		_ = res.Body.Close()
		return "", err
	}
	bodyStr := string(body)
	_ = res.Body.Close()
	return bodyStr, nil
}

func Request(method, url, bodyStr string) (*http.Request, error) {
	switch method {
	case "POST", "Post", "post":
		req, err := http.NewRequest("POST", url, strings.NewReader(bodyStr))
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "GET", "Get", "get":
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "OPTIONS", "Options", "options":
		req, err := http.NewRequest("OPTIONS", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "PUT", "Put", "put":
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "HEAD", "Head", "head":
		req, err := http.NewRequest("HEAD", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "DELETE", "Delete", "delete":
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "TRACE", "Trace", "trace":
		req, err := http.NewRequest("TRACE", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	case "CONNECT", "Connect", "connect":
		req, err := http.NewRequest("CONNECT", url, nil)
		if err != nil {
			cp.Println("error", "NewRequest Error:", err)
			return nil, err
		}
		return req, nil
	default:
		return nil, errors.New("method Error")
	}

}
