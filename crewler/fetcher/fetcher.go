package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

var headers = []string{
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; it; rv:1.8.1.11) Gecko/20071127 Firefox/2.0.0.11",
	"Opera/9.25 (Windows NT 5.1; U; en)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Kubuntu)",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.12) Gecko/20070731 Ubuntu/dapper-security Firefox/1.5.0.12",
	"Lynx/2.8.5rel.1 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/1.2.9",
	"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.04 Chromium/16.0.912.77 Chrome/16.0.912.77 Safari/535.7",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:10.0) Gecko/20100101 Firefox/10.0 ",
}

func Fetch(url string) ([]byte, error) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	request.Header.Add("cookie", "sid=3ddf41c6-ddf1-4c9a-8fd6-c7a10961f4f0; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1627461021; ec=wWE2bwH0-1627464650265-1340f4b2f13d3-1903611607; FSSBBIl1UgzbN7NO=5JhFapgQDa4lymWr0Z8Q_MA1G9ecEBQxvmmkHO_EdwjPCQnXb1LXQL6NdkoWehvbogN34kbR.bsF9jDOQAtOBMG; _exid=RKFN%2BCy3XKEYIDoXFLUQbkNY1MAL5O51G96N7W53oT6r0NRtT7Z92ZMG%2F6wz44qseIdGmPbIUhdd3xYiv3BFiw%3D%3D; _efmdata=FnTY%2FRXwqliJlLAYZK5v5OGx3Shm5VgoFrs12c6gtBajQmZSwLLNItkS%2FQj2c3k4basV3vB6%2BXiX19mYmlXujXecR9YkKEEgSvt8RZL0ZH4%3D; FSSBBIl1UgzbN7NP=53As_0Kll570qqqm_56HcGaOe1EHyFbgYk2cgogb5YMuyvLOSQarq.2zypTvzvIE_2OEmD3eQxaMDKAtD7vZPnViFpPQW36RGJGLLwgLWoWWpvNIDpL3TjtWQ5lFd5GcJ8RE9J4Q1jqpX8HHboze10LQrVt1Aq5.BbpnyUSPc1a0gWMjqL68bc9IS9sqG7tafpmvBqixmSsgpDmCQdwC6dmNZUgbZkt_S.ta9WFoa9S0vyL3IIVPs9IFQi6Ho24GPL; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1627716302")

	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	for i := 0; i < 3; i++ {
		if resp.StatusCode == http.StatusAccepted {
			log.Printf("wrong code retry %d", i)
			time.Sleep(time.Duration(2) * time.Second)
			resp, err = (&http.Client{}).Do(request)
			if err != nil {
				return nil, err
			}
		} else if resp.StatusCode == http.StatusOK {
			break
		} else {
			return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		}

	}

	//if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
	//	return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	//}
	bufReader := bufio.NewReader(resp.Body) //  bufio — 缓存IO

	e := determineEncoding(bufReader)
	utf8Reader := transform.NewReader(bufReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 自动发现编码类型
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html
	bytes, err := r.Peek(1024) // Peek只是“窥探”一下 Reader 中没有读取的 n 个字节。好比栈数据结构中的取栈顶元素，但不出栈。
	if err != nil {
		log.Printf("fetcher error:%+v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// re 匹配
func submatch(html []byte) [][][]byte {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]*>([^<]+)</a>`)
	stringSubmatch := re.FindAllSubmatch(html, -1)
	for _, strings := range stringSubmatch {
		fmt.Printf("cit: %s , url: %s \n", strings[2], strings[1])

	}
	fmt.Println(len(stringSubmatch))
	return stringSubmatch
}
