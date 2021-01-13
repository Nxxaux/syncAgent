package httprequest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"syncAgent-go/syncAgent/params"
)

//Get GET syncthing srv获取随机folderID
/*
url 地址
*/
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err //log
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-CSRF-Token-UE2TW", params.CSRF)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err //log
	}
	defer res.Body.Close()

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return d, nil
}

//PUT url request to syncthing
func PUT(url string, d []byte) error {
	body := bytes.NewBuffer(d)
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return err //log
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-CSRF-Token-UE2TW", params.CSRF)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err //log
	}
	defer res.Body.Close()

	return nil
}
