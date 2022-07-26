package routes

import (
	"context"
	"fmt"
	"net/http"

	"klickhr-apigateway/pkg/hris/pb"

	"github.com/gin-gonic/gin"
)

type UploadHRISRequestBody struct {
	FileType       int    `json:"file_type"`
	RunType        int    `json:"run_type"`
	ImportType     int    `json:"import_type"`
	OrganizationId int    `json:"organization_id"`
	File           []byte `json:"file"`
}

func UploadHRIS(ctx *gin.Context, c pb.HRISServiceClient) {
	body := UploadHRISRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UploadHRIS(context.Background(), &pb.UploadHRISRequest{
		FileType:       int64(body.FileType),
		RunType:        int64(body.RunType),
		ImportType:     int64(body.ImportType),
		OrganizationId: int64(body.OrganizationId),
		File:           body.File,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	fmt.Println("a")
	ctx.JSON(int(res.Status), &res)
}
