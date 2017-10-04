// 当我们需要传入不同的参数，并期待得到不同结果时，就要用到table test。一个table test除了维护
//一张不同参数和结果的表外，和基本的单元测试很像。
package listing08

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// 继续使用http.Get作为例子。
func TestDownload(t *testing.T) {
	//下面是一张不同参数和记过的表。
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.baidu.com",
			http.StatusOK,
		},
		{
			"http://www.csgo.com.cn/esports/asdf.html",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url.",
					checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status. %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
