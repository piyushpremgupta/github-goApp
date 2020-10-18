package github

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
)


func GetShaBranch(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")
	switch r.Method{
		case "GET":
			Auth:=r.Header.Get("Authorization")
			githubApiBaseUrl:="https://api.github.com"
			ReqPathParam:=strings.TrimPrefix(r.URL.Path,"/GetShaBranch/")
			Org:=strings.Split(ReqPathParam,"/")[0]
			Repo:=strings.Split(ReqPathParam,"/")[1]
			RequestUrl:=fmt.Sprintf("%s/repos/%s/%s/git/ref/heads/main",githubApiBaseUrl,Org,Repo)
			
			Response:=MakeRequest(r.Method,RequestUrl,Auth,nil)
			value,ok:=Response.(map[string]interface{})
			
			if ok{
				val,done:=value["object"].(map[string]interface{})
				
				if done {
					//fmt.Println(val)
					json.NewEncoder(w).Encode(map[string]interface{}{"sha":val["sha"]})
				}
			}
		default:
			json.NewEncoder(w).Encode(map[string]string{"Error":"Method Not Supported"})

	}
}