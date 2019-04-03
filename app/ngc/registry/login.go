package registry

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	resty "gopkg.in/resty.v1"
)

var (
	host = "api.ngc.nvidia.com"
)

func init() {
	resty.
		SetDebug(false).
		SetHostURL(fmt.Sprintf("https://%s", host)).
		SetRESTMode().
		SetRetryCount(3).
		SetRetryWaitTime(100 * time.Millisecond).
		SetRetryMaxWaitTime(1 * time.Second)
}

// Debug makes HTTP requests debug mode
func Debug(debug bool) {
	resty.SetDebug(debug)
}

// SetAPIHost change its host name
func SetAPIHost(hostname string) {
	resty.SetHostURL(fmt.Sprintf("https://%s", hostname))
	host = hostname
}

// Login to NGC web console
func Login(email, password string) (token *string, err error) {
	resp, err := resty.R().
		SetBasicAuth(email, password).
		SetHeaders(map[string]string{
			"Referer":      "https://ngc.nvidia.com",
			"Content-Type": "application/json",
		}).
		SetResult(&loginResult{}).
		Post("/login")
	if err != nil {
		return nil, err
	}
	ret, ok := resp.Result().(*loginResult)
	if !ok || !strings.EqualFold(ret.RequestStatus.StatusCode, "SUCCESS") {
		return nil, fmt.Errorf("Status failed or parse error")
	}
	session := ret.UserSession.SessionToken

	for _, cookie := range resp.Cookies() {
		if strings.EqualFold(cookie.Name, "SessionToken") {
			bytes, err := base64.URLEncoding.DecodeString(cookie.Value)
			if err != nil {
				break
			}
			if strings.EqualFold(session, string(bytes)) {
				_, _, err = ParseJWT(session)
				if err != nil {
					return nil, err
				}
				return swag.String(session), nil
			}
		}
	}
	return nil, fmt.Errorf("Cookies does not have any SessionToken")
}

type requestStatus struct {
	StatusCode string `json:"statusCode"`
	RequestID  string `json:"requestId"`
}

type userSession struct {
	SessionToken string `json:"sessionToken"`
}

type loginResult struct {
	RequestStatus requestStatus `json:"requestStatus"`
	UserSession   userSession   `json:"userSession"`
}
