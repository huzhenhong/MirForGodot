/**
#*****************************************************************************
# @file    table.go
# @author  MakerYang(https://www.makeryang.com)
# @statement 免费课程配套开源项目，任何形式收费均为盗版
#*****************************************************************************
*/

package GameServerData

import (
	"Game/framework/database"
	"Game/framework/utils"
)

var TableName = "game_server_data"

type Data struct {
	ServerId     int    `gorm:"primary_key;AUTO_INCREMENT;unique_index;not null;column:server_id"`
	ServerName   string `gorm:"column:server_name"`
	ServerStatus int    `gorm:"column:server_status"`
	Database.DefaultField
}

type Return struct {
	Token      string `json:"token"`
	ServerName string `json:"server_name"`
}

func ReturnData(dataStruct *Data) Return {

	data := Return{}

	if dataStruct.ServerId > 0 {
		data.Token = Utils.EncodeId(32, dataStruct.ServerId)
		data.ServerName = dataStruct.ServerName
	}

	return data
}
