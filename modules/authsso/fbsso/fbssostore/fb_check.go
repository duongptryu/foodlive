package fbssostore

import (
	"context"
	"encoding/json"
	"foodlive/common"
	"foodlive/modules/authsso/fbsso/fbssomodel"
	"net/http"
	"time"
)

func (*authSsoStore) ValidateFbJwt(ctx context.Context, tokenString string) (*fbssomodel.FacebookUser, error) {
	facebookSsoRequest, err := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name&access_token="+tokenString, nil)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(facebookSsoRequest)

	if err != nil || resp.StatusCode != 200 {
		return nil, fbssomodel.ErrInvalidFacebookJwt
	}
	defer resp.Body.Close()

	var fbUser fbssomodel.FacebookUser

	err = json.NewDecoder(resp.Body).Decode(&fbUser)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return &fbUser, nil
}
