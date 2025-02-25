package ads

import (
	"context"
	"github.com/AdPalfish/marketing-api-go-sdk/pkg/config"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

// AuthMiddleware ...
type AuthMiddleware struct {
	tads *SDKClient
}

// Handle ...
func (a *AuthMiddleware) Handle(
	req *http.Request,
	next func(req *http.Request) (rsp *http.Response, err error),
) (rsp *http.Response, err error) {
	nonce := uuid.NewV4().String()
	apiKey := config.APIKey{
		AccessToken: a.tads.GetAccessToken(),
		Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
		Nonce:       nonce[0:8] + nonce[9:13] + nonce[14:18],
	}
	query := req.URL.Query()
	query.Set("access_token", apiKey.AccessToken)
	query.Set("timestamp", apiKey.Timestamp)
	query.Set("nonce", apiKey.Nonce)
	req.URL.RawQuery = query.Encode()
	ctx := context.WithValue(req.Context(), config.ContextAPIKey, apiKey)
	req = req.WithContext(ctx)
	rsp, err = next(req)
	return rsp, err
}
