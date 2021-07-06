// 创建人：lg
// 创建时间：2021/4/7
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	SiteId int64 `json:"site_id,string"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			s Site
			siteIdStr string
		)
		siteIdStr = r.URL.Query().Get("site_id")
		print(siteIdStr)
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(s.SiteId)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
