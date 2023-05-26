package response

import "github.com/VINDA-98/Tenacity-Solutions/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
