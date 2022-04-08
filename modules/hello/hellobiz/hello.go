package hellobiz

import (
	"context"
	"errors"
)

type helloBiz struct {

}

func NewHelloBiz () *helloBiz {
	return &helloBiz{}
}

func (biz *helloBiz) WayBiz (ctx context.Context, id int) (string, error) {
	if id == 0 {
		return "", errors.New("id is empty")
	}
	return "Duong", nil
}