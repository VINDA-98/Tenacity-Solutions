package system

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
