package httpclient

import "time"

// Option 自定义设置http请求
type Option func(*option)

type option struct {
	ttl    time.Duration
	header map[string][]string
}

// WithTTL 本次http请求最长执行时间
func WithTTL(ttl time.Duration) Option {
	return func(opt *option) {
		opt.ttl = ttl
	}
}

// WithHeader 设置http header，可以调用多次设置多对key-value
func WithHeader(key, value string) Option {
	return func(opt *option) {
		opt.header[key] = []string{value}
	}
}

func NewOpt() *option {
	return &option{
		header: make(map[string][]string),
	}
}
