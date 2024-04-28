package main

import (
	"fmt"
	"github.com/Kumengda/pageParser/parser"
	"io"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://127.0.0.1:8765/vul/rce/rce_eval.php")
	body, _ := io.ReadAll(resp.Body)

	tagExtract := parser.NewTagExtract()
	tagExtract.InitTags(parser.DefaultTagRules)
	tagExtract.AddTag("a", "class")
	formExtract := parser.NewFormExtract()

	tag := tagExtract.Extract(string(body))
	form := formExtract.Extract(string(body))

	for _, t := range tag {
		fmt.Println(t)
	}
	for _, f := range form {
		fmt.Println("--------form--------")
		fmt.Println("method:" + f.Method)
		fmt.Println("action" + f.Action)
		fmt.Println("encType" + f.Enctype)
		fmt.Printf("formData:%v\n", f.FormData)
		fmt.Println("--------form--------")

	}
}
