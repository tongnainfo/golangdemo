package model

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func HttpDemo() {
	resp, erro := http.Get("http://www.baidu.com")
	if erro == nil {
		if (resp != nil && resp.Body != nil) {
			defer resp.Body.Close()
		} else {
			log.Println("ERROR " + " 返回为空 ")
		}
		if resp == nil || resp.Body == nil || erro != nil || resp.StatusCode != http.StatusOK {
			log.Println("ERROR ")
			log.Println(erro)
			return
		}

		var buf []byte
		buf, err := ioutil.ReadAll(resp.Body)
		if (err != nil) {
			return
		}
		fmt.Println(string(buf));
	}

}
