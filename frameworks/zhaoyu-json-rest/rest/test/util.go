package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//MakeSimpleRequest 返回一个自定义的http.request对象，方便进一步添加headers，
//query string等参数。
func MakeSimpleRequest(method string, urlStr string, payload interface{}) *http.Request {
	var s string

	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		s = fmt.Sprintf("%s", b)
	}

	r, err := http.NewRequest(method, urlStr, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	r.Header.Set("Accept-Encoding", "gzip")
	if payload != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	return r
}

//比较responseWriter返回的code。
func CodeIs(t *testing.T, r *httptest.ResponseRecorder, expectedCode int) {
	if r.Code != expectedCode {
		t.Errorf("Code %d expected, got: %d", expectedCode, r.Code)
	}
}

// HeaderIs 比较header 中已知key的第一个值
func HeaderIs(t *testing.T, r *httptest.ResponseRecorder, headerKey, expectedValue string) {
	value := r.HeaderMap.Get(headerKey)
	if value != expectedValue {
		t.Errorf(
			"%s: %s expected, got: %s",
			headerKey,
			expectedValue,
			value,
		)
	}
}

//测试返回是否是json,utf8
func ContentTypeIsJson(t *testing.T, r *httptest.ResponseRecorder) {

	mediaType, params, _ := mime.ParseMediaType(r.HeaderMap.Get("Content-Type"))
	charset := params["charset"]

	if mediaType != "application/json" {
		t.Errorf(
			"Content-Type media type: application/json expected, got: %s",
			mediaType,
		)
	}

	if charset != "" && strings.ToUpper(charset) != "UTF-8" {
		t.Errorf(
			"Content-Type charset: must be empty or UTF-8, got: %s",
			charset,
		)
	}
}

//测试返回是不是gzip
func ContentEncodingIsGzip(t *testing.T, r *httptest.ResponseRecorder) {
	HeaderIs(t, r, "Content-Encoding", "gzip")
}

//测试返回body
func BodyIs(t *testing.T, r *httptest.ResponseRecorder, expectedBody string) {
	body := r.Body.String()
	if body != expectedBody {
		t.Errorf("Body '%s' expected, got: '%s'", expectedBody, body)
	}
}

//测试body的decode类型
func DecodeJsonPayload(r *httptest.ResponseRecorder, v interface{}) error {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, v)
	if err != nil {
		return err
	}
	return nil
}

//封装了ResponseRecorder 和testing.T
type Recorded struct {
	T        *testing.T
	Recorder *httptest.ResponseRecorder
}

//通过给出的handler执行一个HTTP请求
func RunRequest(t *testing.T, handler http.Handler, request *http.Request) *Recorded {
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return &Recorded{t, recorder}
}

func (rd *Recorded) CodeIs(expectedCode int) {
	CodeIs(rd.T, rd.Recorder, expectedCode)
}

func (rd *Recorded) HeaderIs(headerKey, expectedValue string) {
	HeaderIs(rd.T, rd.Recorder, headerKey, expectedValue)
}

func (rd *Recorded) ContentTypeIsJson() {
	ContentTypeIsJson(rd.T, rd.Recorder)
}

func (rd *Recorded) ContentEncodingIsGzip() {
	rd.HeaderIs("Content-Encoding", "gzip")
}

func (rd *Recorded) BodyIs(expectedBody string) {
	BodyIs(rd.T, rd.Recorder, expectedBody)
}

func (rd *Recorded) DecodeJsonPayload(v interface{}) error {
	return DecodeJsonPayload(rd.Recorder, v)
}
