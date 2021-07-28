package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	httpURL "net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	//一次http请求最长执行5秒
	DefaultTTL = time.Second * 5
)

func Get(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withoutBody(http.MethodGet, url, form, options...)
}

// PostForm post form 请求
func PostForm(url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	return withFormBody(http.MethodPost, url, form, options...)
}

// PostJSON post json 请求
func PostJSON(url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	return withJSONBody(http.MethodPost, url, raw, options...)
}

func withoutBody(method, url string, form httpURL.Values, option ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}

	if len(form) > 0 {
		if url, err = addFormValuesIntoURL(url, form); err != nil {
			return nil, err
		}
	}

	opt := NewOpt()
	for _, f := range option {
		f(opt)
	}

	opt.header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	ttl := opt.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	body, _, err = doHTTP(ctx, method, url, nil, opt)
	return
}

func withFormBody(mehtod, url string, form httpURL.Values, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(form) == 0 {
		return nil, errors.New("form required")
	}

	opt := NewOpt()
	for _, f := range options {
		f(opt)
	}
	opt.header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	ttl := opt.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	formValue := form.Encode()
	body, _, err = doHTTP(ctx, mehtod, url, []byte(formValue), opt)
	return
}

func withJSONBody(method, url string, raw json.RawMessage, options ...Option) (body []byte, err error) {
	if url == "" {
		return nil, errors.New("url required")
	}
	if len(raw) == 0 {
		return nil, errors.New("raw required")
	}

	opt := NewOpt()
	for _, f := range options {
		f(opt)
	}
	opt.header["Content-Type"] = []string{"application/json; charset=utf-8"}

	ttl := opt.ttl
	if ttl <= 0 {
		ttl = DefaultTTL
	}

	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()

	body, _, err = doHTTP(ctx, method, url, raw, opt)
	return
}
