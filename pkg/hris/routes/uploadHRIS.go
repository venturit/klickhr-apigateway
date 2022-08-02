package routes

import (
	"io"

	"klickhr-apigateway/pkg/hris/pb"

	"github.com/gin-gonic/gin"
)

type UploadHRISRequestBody struct {
	FileType       int `form:"file_type"`
	RunType        int `form:"run_type"`
	ImportType     int `form:"import_type"`
	OrganizationId int `form:"organization_id"`
}

func UploadHRIS(ctx *gin.Context, c pb.HRISServiceClient) {
	body := UploadHRISRequestBody{}
	ctx.Bind(&body)
	//get file
	file, _ := ctx.FormFile("file")
	// Maximum 10Mb size per stream.
	buf := make([]byte, 1024*10000)
	stream, err := c.UploadHRIS(ctx)
	if err != nil {
		return
	}
	fil, err := file.Open() //os.Open("./" + file.Filename)
	if err != nil {
		return
	}
	for {
		num, err := fil.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}

		if err := stream.Send(&pb.UploadHRISRequest{
			FileType:       int64(body.FileType),
			RunType:        int64(body.RunType),
			ImportType:     int64(body.ImportType),
			OrganizationId: int64(body.OrganizationId),
			FileName:       file.Filename,
			FileBytes:      buf[:num],
		}); err != nil {
			return
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		return
	}
	ctx.JSON(int(res.Status), &res)
}
