package httprequest

import (
	"fmt"
	"syncAgent-go/syncAgent/params"
	"testing"
)

//go test -v .\httprequest_test.go .\httpReqHandlers.go .\requtils.go
func TestRandomfolderID(t *testing.T) {
	//CSRF获取+本地deviceID获取
	ok, err := CSRF()
	if !ok || err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("CSRF:%s,deviceID:%s\n", params.CSRF, params.MetaDeviceID)
	//获取RandomfolderID
	id, err := randomFolderID()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("this is random folder id:%s\n", id)
}
