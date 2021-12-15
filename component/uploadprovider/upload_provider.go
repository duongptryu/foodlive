package uploadprovider

import (
	"context"
	"fooddelivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
