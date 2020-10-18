package github

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"bytes"
)


func CreatePullRequest(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")
	switch r.Method{
	case "POST":

				var data map[string]string
				json.NewDecoder(r.Body).Decode(&data)
				Auth:=r.Header.Get("Authorization")
				githubApiBaseUrl:="https://api.github.com"
				ReqPathParam:=strings.TrimPrefix(r.URL.Path,"/CreatePullRequest/")
				Org:=strings.Split(ReqPathParam,"/")[0]
				Repo:=strings.Split(ReqPathParam,"/")[1]
				
				RequestUrl:=fmt.Sprintf("%s/repos/%s/%s/pulls",githubApiBaseUrl,Org,Repo)
				
				
				Payload:= map[string]interface{}{"title":data["Title"],"head":data["Head"],"base":data["Base"]}
				PayloadBytes,err:=json.Marshal(Payload)
				if err != nil {
					json.NewEncoder(w).Encode(map[string]interface{}{"Error":err})
				}
				
				Response:=MakeRequest(r.Method,RequestUrl,Auth,bytes.NewBuffer(PayloadBytes))
				fmt.Println(Response)
				val,ok:=Response.(map[string]interface{})
				
				if ok{
					json.NewEncoder(w).Encode(map[string]interface{}{"PullUrl":val["url"]})
					}else{
						json.NewEncoder(w).Encode(map[string]interface{}{"Error":"Could not create pull request"})
					}
									
	default:
		json.NewEncoder(w).Encode(map[string]string{"Error":"Method Not Supported"})

}
}