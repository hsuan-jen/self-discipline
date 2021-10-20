package system

import (
	"self-discipline/global"
	"self-discipline/model/system"
)

type NicknameService struct{}

func (*NicknameService) GetNickname() (info system.SysNickname, err error) {

	err = global.DB.Raw(`SELECT * FROM sys_nicknames JOIN ( SELECT round( rand()*( SELECT max( id ) FROM sys_nicknames )) AS idd 
		) AS sys_nicknames ON sys_nicknames.id > sys_nicknames.idd WHERE status = 0
		LIMIT 1;`).Scan(&info).Error
	return
}
