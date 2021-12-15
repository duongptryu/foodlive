package ginupload

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/upload/uploadbiz"
	"fooddelivery/modules/upload/uploadstore"
	"github.com/gin-gonic/gin"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		//store image into own system
		//c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstore.NewSQLStore(appCtx.GetDatabase())
		biz := uploadbiz.NewUploadBiz(appCtx.GetUploadProvider(), imgStore)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.NewSimpleSuccessResponse(img))
	}
}
