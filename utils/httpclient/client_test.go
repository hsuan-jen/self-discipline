package httpclient

import (
	httpURL "net/url"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	url := "http://www.weather.com.cn/data/sk/101190408.html"

	form := make(httpURL.Values)
	form["name"] = []string{"amiy"}

	body, err := Get(url, form, WithTTL(time.Second*5), WithHeader("Authorization", "demoGetAuthorization"))

	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}

func BenchmarkGet(b *testing.B) {
	url := "http://www.weather.com.cn/data/sk/101190408.html"

	for i := 0; i < b.N; i++ {
		Get(url, nil, WithTTL(time.Second*5), WithHeader("Authorization", "demoGetAuthorization"))
	}
}

func TestPostForm(t *testing.T) {
	url := "http://www.weather.com.cn/data/sk/101190408.html"

	form := make(httpURL.Values)
	form["name"] = []string{"amiy"}

	body, err := PostForm(url, form, WithTTL(time.Second*5))

	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}
