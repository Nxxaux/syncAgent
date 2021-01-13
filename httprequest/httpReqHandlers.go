package httprequest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

//获取随机的folderID
func randomFolderID() (string, error) {
	d, err := Get(params.TryRandomFolderID)
	if err != nil {
		return "", err
	}
	//添加之后访问添加文件夹
	FolderID := func(d []byte) string {
		f := params.FolderIDRandom{}
		err := json.Unmarshal(d, &f)
		if err != nil {
			//shuold not happen
			return "" //log
		}
		sl := strings.ToLower(string(f.Random))
		s1 := sl[:5]
		s2 := sl[5:]
		s3 := s1 + "-" + s2

		return s3
	}(d)
	return FolderID, nil
}

//CSRF GET 请求并返回结果cookie
/*
url： 地址
flag：
0 默认从文件获取
1 从syncthing 服务器获取
*/
func CSRF() (bool, error) {
	var csrf string
	sa, err := readCsrfRecord()
	if err != nil {
		return false, err //log
	}
	csrf = sa
	ok, err := csrfVarify(params.TryERR, csrf)
	if !ok {
		sb, err := csrfGet()
		if err != nil {
			return false, err //log
		}
		csrf = sb
	}
	params.CSRF = csrf
	return true, nil
}
