package response

import "github.com/VINDA-98/Tenacity-Solutions/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
