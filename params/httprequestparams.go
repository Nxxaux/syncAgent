package params

//FolderIDRandom for folderID random form sysncthing server
type FolderIDRandom struct {
	Random string `json:"random"`
}

//CSRFlength varify CSRF
const CSRFlength = 32

//GCHK config hanlde key
var GCHK = Configuration{}
