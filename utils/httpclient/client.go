package httpclient

import (
	"time"
)

const (
	//一次http请求最长执行10秒
	DefaultTTL = time.Second * 10
)

/* func Get(url string, form url.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodGet, url, form)
}

func withoutBody(method, url string, form url.Values, option ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}

	if len(form) > 0 {
		if url, err := addFormValuesIntoURL(url, form) {

		}
	}

	//return , nil
} */
