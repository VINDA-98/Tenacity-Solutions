package system

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
