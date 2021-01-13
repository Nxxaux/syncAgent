package httprequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"syncAgent-go/syncAgent/params"
)

const csrfRecordPath = "./record.txt"

//获取csrf
func csrfGet() (string, error) {
	req, err := http.NewRequest("GET", params.TryLoginURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	for k, v := range res.Header {
		if k == "X-Syncthing-Id" {
			params.MetaDeviceID = v[0]
		}
	}
	return res.Cookies()[0].Value, nil
}

//校验csrf 是否有效
func csrfVarify(url string, csrf string) (bool, error) {
	req, err := http.NewRequest("GET", params.TryERR, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-CSRF-Token-UE2TW", csrf)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

//从文件中读取csrf
func readCsrfRecord() (string, error) {
	fd, err := os.OpenFile(csrfRecordPath, os.O_CREATE|os.O_RDONLY, 0755)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	return string(b), nil
}

//写入csrf文件
func writeCsrfRecord(content string) (bool, error) {
	fd, err := os.OpenFile(csrfRecordPath, os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return false, err
	}
	defer fd.Close()
	fmt.Fprintln(fd, content)
	return true, nil
}
