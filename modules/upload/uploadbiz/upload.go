package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"foodlive/common"
	"foodlive/component/uploadprovider"
	"foodlive/modules/upload/uploadmodel"
	log "github.com/sirupsen/logrus"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{imgStore: imgStore, provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDememsion(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	//img.CloudName = "s3" //set in Provider
	img.Extension = fileExt

	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	//delete img on s3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}

	return img, nil
}

func getImageDememsion(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
