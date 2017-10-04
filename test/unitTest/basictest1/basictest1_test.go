// 这是一个基本的test用例，用来测试http Get方法.
package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// Test方法必须以Test开头，并且接受 testing.T的指针参数，不设置返回参数.
func TestDownload(t *testing.T) {
	url := "http://www.baidu.com"
	statusCode := 200
	//测试日志，(测试时加-v参数可以看到),每个测试应该用Given the need 说明测试的目的。
	t.Log("Given the need to test downloading content.")
	{
		//测试日志，测试时加-v参数可以看到，测试应该用when说明如何算测试通过。
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				//t.Fatal说明测试失败。并停止这个测试。
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark)
			} else {
				//t.Error说明测试出错，单并不会退出测试。
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
		}
	}
	//最后如果t.Error和t.Fatal都没有被调，说明测试通过。
}
