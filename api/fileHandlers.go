package handlers

/*
文件夹相关操作
*/

//AgentCreateFolder 创建文件夹
/*
参数
folderID    文件夹ID
folderLabel 文件夹标签
folderCreatePath  文件夹创建路径
*/
func AgentCreateFolder(folderID, folderLabel,
	folderCreatePath string) (bool, error) {

	return true, nil
}

//AgentRemoveFolder 移除文件夹
/*
folderID    文件夹ID
*/
func AgentRemoveFolder(folderID string) (bool, error) {

	return true, nil
}

//AgentPauseFolder 暂停文件夹
/*
folderID    文件夹ID
*/
func AgentPauseFolder(folderID string) (bool, error) {

	return true, nil
}

//AgentRecoverFolder 恢复文件夹
/*
folderID    文件夹ID
*/
func AgentRecoverFolder(folderID string) (bool, error) {

	return true, nil
}

//AgentRescanFolder 重新扫描文件夹
/*
folderID    文件夹ID
*/
func AgentRescanFolder(folderID string) (bool, error) {

	return true, nil
}

//AgentsetFolderConfig 修改文件夹选项
/*
folderID    文件夹ID
*/
func AgentsetFolderConfig(folderID string) (bool, error) {

	return true, nil
}
