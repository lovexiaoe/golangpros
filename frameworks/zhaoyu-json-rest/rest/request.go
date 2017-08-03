package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	// 当 JSON payload 是空时返回，payload指有效负载，一个在网络中传输的信息包包括一些
	//辅助信息（数据量大小，校验位等）和有效负载（信息的原始数据）
	ErrJsonPayloadEmpty = errors.New("JSON payload is empty")
)

// Request inherits from http.Request, and provides additional methods.
type Request struct {
	*http.Request

	// Map of parameters that have been matched in the URL Path.
	PathParams map[string]string

	// Environment used by middlewares to communicate.
	Env map[string]interface{}
}

// PathParam provides a convenient access to the PathParams map.
func (r *Request) PathParam(name string) string {
	return r.PathParams[name]
}

// DecodeJsonPayload 读取 request 的 body 并使用 json.Unmarshal编码。
func (r *Request) DecodeJsonPayload(v interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return err
	}
	if len(content) == 0 {
		return ErrJsonPayloadEmpty
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return err
	}
	return nil
}

// BaseUrl 返回一个新的 URL 对象，包括了从 request中取到的 Host 和 Scheme .
// (host中没有最后的斜杠)
func (r *Request) BaseUrl() *url.URL {
	scheme := r.URL.Scheme
	if scheme == "" {
		scheme = "http"
	}

	host := r.Host
	if len(host) > 0 && host[len(host)-1] == '/' {
		host = host[:len(host)-1]
	}

	return &url.URL{
		Scheme: scheme,
		Host:   host,
	}
}

//URLFor返回设置了Path和query string的URL对象，其中的query string使用queryParams构建
func (r *Request) UrlFor(path string, queryParams map[string][]string) *url.URL {
	baseUrl := r.BaseUrl()
	baseUrl.Path = path
	if queryParams != nil {
		query := url.Values{}
		for k, v := range queryParams {
			for _, vv := range v {
				query.Add(k, vv)
			}
		}
		baseUrl.RawQuery = query.Encode()
	}
	return baseUrl
}

// CorsInfo 包含了从 rest.Request取出的 CORS request 信息。
// CORS 跨域资源共享，Cross-Origin Resource Sharing
type CorsInfo struct {
	IsCors      bool
	IsPreflight bool
	Origin      string
	OriginUrl   *url.URL

	// The header value 为避免一些错误，转换成大写。
	AccessControlRequestMethod string

	// header values 使用 http.CanonicalHeaderKey方法规范化。
	AccessControlRequestHeaders []string
}

// 从Request中获取 CorsInfo 信息。
func (r *Request) GetCorsInfo() *CorsInfo {

	origin := r.Header.Get("Origin")

	var originUrl *url.URL
	var isCors bool

	//判断是否可以跨源
	if origin == "" {
		isCors = false
	} else if origin == "null" {
		isCors = true
	} else {
		var err error
		originUrl, err = url.ParseRequestURI(origin)
		isCors = err == nil && r.Host != originUrl.Host
	}

	//取得可跨源请求方法
	reqMethod := r.Header.Get("Access-Control-Request-Method")

	//取得可跨源请求头
	reqHeaders := []string{}
	rawReqHeaders := r.Header[http.CanonicalHeaderKey("Access-Control-Request-Headers")]
	for _, rawReqHeader := range rawReqHeaders {
		// net/http does not handle comma delimited headers for us
		for _, reqHeader := range strings.Split(rawReqHeader, ",") {
			reqHeaders = append(reqHeaders, http.CanonicalHeaderKey(strings.TrimSpace(reqHeader)))
		}
	}

	//判断isPreflight标志
	isPreflight := isCors && r.Method == "OPTIONS" && reqMethod != ""

	return &CorsInfo{
		IsCors:                      isCors,
		IsPreflight:                 isPreflight,
		Origin:                      origin,
		OriginUrl:                   originUrl,
		AccessControlRequestMethod:  strings.ToUpper(reqMethod),
		AccessControlRequestHeaders: reqHeaders,
	}
}
