package response

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
