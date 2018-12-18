package main
import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
)

func main() {
	url := "http://www.cmbc.com.cn/jrms/msdt/yjbg/index.htm"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

