package httpclient

import (
	"bytes"
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	// _StatusReadRespErr read resp body err, should re-call doHTTP again.
	_StatusReadRespErr = -204
	// _StatusDoReqErr do req err, should re-call doHTTP again.
	_StatusDoReqErr = -500
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives:  true,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
	},
}

// addFormValuesIntoURL 追加参数到URL
func addFormValuesIntoURL(rawURL string, form url.Values) (string, error) {
	if rawURL == "" {
		return "", errors.New("rawURL required")
	}

	if len(form) == 0 {
		return "", errors.New("form required")
	}

	target, err := url.Parse(rawURL)

	if err != nil {
		return "", errors.Wrapf(err, "parse rawURL `%s` err", rawURL)
	}

	urlValues := target.Query()
	for key, values := range form {
		for _, val := range values {
			urlValues.Add(key, val)
		}
	}

	target.RawQuery = urlValues.Encode()
	return target.String(), nil
}

func doHTTP(ctx context.Context, method, url string, payload []byte, opt *option) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, -1, errors.Wrapf(err, "new request [%s %s] err", method, url)
	}

	for key, value := range opt.header {
		req.Header.Set(key, value[0])
	}

	resp, err := defaultClient.Do(req)
	if err != nil {
		return nil, _StatusReadRespErr, errors.Wrapf(err, "do request [%s %s] err", method, url)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, _StatusReadRespErr, errors.Wrapf(err, "read resp body from [%s %s] err", method, url)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.Errorf("do [%s %s] return code: %d message: %s", method, url, resp.StatusCode, string(body))
	}

	return body, http.StatusOK, nil
}
