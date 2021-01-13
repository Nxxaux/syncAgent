package httprequest

import (
	"fmt"
	"log"
	"syncAgent-go/syncAgent/params"
	"testing"
)

//"C:\\Users\\57393\\xiicloud1"

//go test -v .\httprequest_test.go .\httpReqHandlers.go .\requtils.go
func TestRandomfolderID(t *testing.T) {
	t.Skipped()
	//CSRF获取+本地deviceID获取
	err := CSRF()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("CSRF:%s,deviceID:%s\n", params.CSRF, params.MetaDeviceID)
	//获取RandomfolderID
	id, err := randomFolderID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("This is random folder id:%s\n", id)
}

func TestGlobleConfigUpdate(t *testing.T) {
	t.Skipped()
	err := CSRF()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = GlobleConfigLoad()
	if err != nil {
		log.Println(err.Error())
		return
	}
	GlobleConfigUpdate()
}

func TestFolderADD(t *testing.T) {
	err := CSRF()
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = GlobleConfigLoad()
	if err != nil {
		log.Println(err.Error())
		return
	}

	id, err := randomFolderID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("This is random folder id:%s\n", id)

	fc := FolderCreate(id, params.MetaDeviceID, "C:\\Users\\57393\\xiicloud1")

	fmt.Printf("this is fc:%v\n", *fc)

	//ADD
	params.GCHK.Folders = append(params.GCHK.Folders, *fc)

	fmt.Printf("this is params.GCHK.Folders:%v\n", params.GCHK.Folders)

	err = GlobleConfigUpdate()
	if err != nil {
		fmt.Println(err.Error())
	}
}
