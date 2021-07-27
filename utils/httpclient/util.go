package httpclient

import (
	"net/url"

	"github.com/pkg/errors"
)

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
