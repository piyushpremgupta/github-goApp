package github

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
)

var data interface{}
func GetShaFile(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")
	switch r.Method{
		case "GET":
			
			Auth:=r.Header.Get("Authorization")
			githubApiBaseUrl:="https://api.github.com"
			ReqPathParam:=strings.TrimPrefix(r.URL.Path,"/GetShaFile/")
			Org:=strings.Split(ReqPathParam,"/")[0]
			Repo:=strings.Split(ReqPathParam,"/")[1]
			FilePath:=strings.Split(ReqPathParam,"/")[2]
			TargetBranch:=strings.Split(ReqPathParam,"/")[3]

			RequestUrl:=fmt.Sprintf("%s/repos/%s/%s/contents/%s?ref=%s",githubApiBaseUrl,Org,Repo,FilePath,TargetBranch)
			
			Response:=MakeRequest(r.Method,RequestUrl,Auth,nil)
			value,ok:=Response.(map[string]interface{})
			
			if ok{
				json.NewEncoder(w).Encode(map[string]interface{}{"sha":value["sha"]})
			}
		default:
			json.NewEncoder(w).Encode(map[string]string{"Error":"Method Not Supported"})

	}
}