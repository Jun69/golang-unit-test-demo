package gock_demo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetResultByClient(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的mock

	// mock 请求外部指定url api时传参x=1返回100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})
	api :=&APIClient{}
	// 调用依赖外部函数的
	res := api.GetResultByAPI(1, 1)
	assert.Equal(t, 101,res )
}
func TestGetResultByAPI(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的mock

	// mock 请求外部api时传参x=1返回100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})

	// 调用我们的业务函数
	res := GetResultByAPI(1, 1)
	// 校验返回结果是否符合预期
	assert.Equal(t, res, 101)

	// mock 请求外部api时传参x=2返回200
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 2}).
		Reply(200).
		JSON(map[string]int{"value": 200})
	// 调用我们的业务函数
	res = GetResultByAPI(2, 2)
	// 校验返回结果是否符合预期
	assert.Equal(t, res, 202)

	assert.True(t, gock.IsDone()) // 断言mock被触发
}

func TestGetResultByAPIByHttpTest(t *testing.T)  {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method== http.MethodPost{
			res.WriteHeader(http.StatusOK)
			res.Write([]byte("body"))
		}else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer func() { testServer.Close() }()

	req, err := http.NewRequest(http.MethodPost, testServer.URL, nil)
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode, "status code should match the expected response")
}
