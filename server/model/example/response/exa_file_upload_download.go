package response

import "github.com/kuihuar/huasu-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
