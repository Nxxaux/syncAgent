package params

const (
	//TryERR GET  测试地址
	TryERR = "http://127.0.0.1:8384/rest/system/error"
	//TryRandomFolderID 随机获取folderID （测试专用）
	TryRandomFolderID = "http://127.0.0.1:8384/rest/svc/random/string?length=10"

	//TryLoginURL GET 获取本地DeviceID 和 CSRF
	TryLoginURL = "http://127.0.0.1:8384/"
	//TryConfigURL PUT  提交配置到syncthing 服务器
	TryConfigURL = "http://127.0.0.1:8384/rest/config"

	///////////////////////删除文件夹///////////////////////////////

	///////////////////////其他用途////////////////////////////////
	//for other method
	tryConnURL      = "http://127.0.0.1:8384/rest/system/connections"
	tryAddURL       = "http://127.0.0.1:8384/rest/svc/random/string?length=10"
	tryAddactionURL = "http://127.0.0.1:8384/rest/system/browse?current=C:%5CUsers%5C57393%5C"
)
