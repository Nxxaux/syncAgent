package httprequest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	if csrf == "" || len(csrf) != params.CSRFlength {
		return false, nil
	}
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
func writeCsrfRecord(content string) error {
	fd, err := os.OpenFile(csrfRecordPath, os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer fd.Close()
	fmt.Fprintf(fd, "%s", content)
	return nil
}

//获取随机的folderID(测试专用)
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
func CSRF() error {
	var csrf string
	sa, err := readCsrfRecord()
	if err != nil {
		return err //log 不需要返回，失败也没事
	}
	csrf = sa
	ok, err := csrfVarify(params.TryERR, csrf)
	if err != nil {
		return err
	}
	if !ok {
		fmt.Println("文件中找的csrf验证未通过") //log
		sb, err := csrfGet()
		if err != nil {
			return err //log
		}
		csrf = sb
	} else {
		_, err := csrfGet()
		if err != nil {
			return err //log
		}
	}
	params.CSRF = csrf
	err = writeCsrfRecord(params.CSRF)
	if err != nil {
		return err //log 不需要返回，失败也没事
	}
	return nil
}

//------------------------------------------------------------------------------------------------------

//GlobleConfigLoad for init
func GlobleConfigLoad() error {
	params.GCHK = params.Configuration{}
	d, err := Get(params.TryConfigURL)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, &params.GCHK)
	if err != nil {
		return err
	}
	return nil
}

//GlobleConfigUpdate update to syncthing
func GlobleConfigUpdate() error {
	d, err := json.Marshal(params.GCHK)
	if err != nil {
		return err
	}
	err = PUT(params.TryConfigURL, d)
	if err != nil {
		return err
	}
	return nil
}

//------------------------------------------------------------------------------------------------------

//FolderCreate create a folder
func FolderCreate(folderID, folderLable,
	folderCreatePath string) *params.FolderConfiguration {
	dev := params.FolderDeviceConfiguration{
		DeviceID: params.MetaDeviceID, //固定参数
	}
	devs := make([]params.FolderDeviceConfiguration, 0)
	devs = append(devs, dev)
	mdisk := params.Size{
		Value: 1,
		Unit:  "%",
	}
	p := make(map[string]string, 0)
	ver := params.VersioningConfiguration{
		Type:             "",
		Params:           p,
		CleanupIntervalS: 0,
	}
	return &params.FolderConfiguration{
		ID:                      folderID,
		Label:                   folderID,
		FilesystemType:          "basic",
		Path:                    folderCreatePath, //测试用，先不修改
		Type:                    "sendreceive",
		Devices:                 devs,
		RescanIntervalS:         3600,
		FSWatcherEnabled:        true,
		FSWatcherDelayS:         10,
		IgnorePerms:             false,
		AutoNormalize:           true,
		MinDiskFree:             mdisk,
		Versioning:              ver,
		Copiers:                 0,
		PullerMaxPendingKiB:     0,
		Hashers:                 0,
		Order:                   "random",
		IgnoreDelete:            false,
		ScanProgressIntervalS:   0,
		PullerPauseS:            0,
		MaxConflicts:            10,
		DisableSparseFiles:      false,
		DisableTempIndexes:      false,
		Paused:                  false,
		WeakHashThresholdPct:    25,
		MarkerName:              ".stfolder",
		CopyOwnershipFromParent: false,
		RawModTimeWindowS:       0,
		MaxConcurrentWrites:     2,
		DisableFsync:            false,
		BlockPullOrder:          "standard",
		CopyRangeMethod:         "standard",
		CaseSensitiveFS:         false,
		JunctionsAsDirs:         false,
	}
}
