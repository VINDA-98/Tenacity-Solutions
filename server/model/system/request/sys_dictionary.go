package request

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/model/common/request"
	"github.com/VINDA-98/Tenacity-Solutions/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
