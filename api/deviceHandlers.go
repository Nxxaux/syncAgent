package handlers

/*
设备相关操作
*/

//AgentBindDevice 绑定设备
/*
localDevID 本地设备ID
remoteDevID 远程设备ID
remoteID 远程文件夹ID
localID 本地文件夹ID
*/
func AgentBindDevice(localDevID, remoteDevID,
	remoteID, localID string) (bool, error) {

	return true, nil
}

//AgentaddRemoteDevice 添加远程设备
/*
remoteDevID  远程设备ID
remoteDeviceNameID 远程设备名称（ID）
*/
func AgentaddRemoteDevice(remoteDevID,
	deviceNameID string) (bool, error) {

	return true, nil
}
