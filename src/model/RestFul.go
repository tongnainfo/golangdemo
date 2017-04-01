package model

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"github.com/emicklei/go-restful"
)

func Restful() {

	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(hello))
	restful.Add(ws)
	fmt.Print("Server starting on port 808\n")
	http.ListenAndServe(":808", nil)
}
var i=0
func hello(req *restful.Request, resp *restful.Response) {
	name := req.QueryParameter("name");
	fmt.Println(i,"---",name)
	i++
	article := Article{"A Royal Baby", "A slow news week"}
	b, _ := json.Marshal(article)
	io.WriteString(resp, string(b))
}

type Article struct {
	Name string
	Body string
}
