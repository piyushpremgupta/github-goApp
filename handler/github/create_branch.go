package github

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"bytes"
)


func CreateBranch(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-type","application/json")
	switch r.Method{
	case "POST":
			var ShaId interface{} 
			Auth:=r.Header.Get("Authorization")
			githubApiBaseUrl:="https://api.github.com"
			ReqPathParam:=strings.TrimPrefix(r.URL.Path,"/CreateBranch/")
			Org:=strings.Split(ReqPathParam,"/")[0]
			Repo:=strings.Split(ReqPathParam,"/")[1]
			TargetBranch:=strings.Split(ReqPathParam,"/")[2]
			ShaBranchResp:=MakeRequest("GET","http://localhost:5001/GetShaBranch/AccuricsAssignment/Accurics",Auth,nil)
			value,ok:= ShaBranchResp.(map[string]interface{})
			if ok{
				ShaId=value["sha"]
			}
			RequestUrl:=fmt.Sprintf("%s/repos/%s/%s/git/refs",githubApiBaseUrl,Org,Repo)
			
			Payload:= map[string]interface{}{"ref":"refs/heads/"+TargetBranch,"sha":ShaId}
			PayloadBytes,err:=json.Marshal(Payload)
			if err != nil {
				json.NewEncoder(w).Encode(map[string]interface{}{"Error":err})
			}
			
			Response:=MakeRequest(r.Method,RequestUrl,Auth,bytes.NewBuffer(PayloadBytes))
			val,ok:=Response.(map[string]interface{})
			
			if ok{
				output,done:=val["object"].(map[string]interface{})
				
				if done {
					//fmt.Println(val)
					json.NewEncoder(w).Encode(map[string]interface{}{"sha":output["sha"]})
				}else{
					json.NewEncoder(w).Encode(map[string]interface{}{"Error":"Could not create the branch"})
				}
			}
			
	default:
		json.NewEncoder(w).Encode(map[string]string{"Error":"Method Not Supported"})

}
}