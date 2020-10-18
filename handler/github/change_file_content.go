package github

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"bytes"
	"encoding/base64"
)


func ChangeFileContent(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")

	switch r.Method{
		case "PUT":
				var ShaId interface{} 
				var data map[string]string
				err:=json.NewDecoder(r.Body).Decode(&data)
				if err!=nil{
					fmt.Println(err)
				}
				Auth:=r.Header.Get("Authorization")
				fmt.Println(Auth)
				githubApiBaseUrl:="https://api.github.com"
				ReqPathParam:=strings.TrimPrefix(r.URL.Path,"/ChangeFileContent/")
				Org:=strings.Split(ReqPathParam,"/")[0]
				Repo:=strings.Split(ReqPathParam,"/")[1]
				FilePath:=strings.Split(ReqPathParam,"/")[2]
				ShaFileReqUrl:=fmt.Sprintf("http://localhost:5001/GetShaFile/%s/%s/%s/%s",Org,Repo,FilePath,data["TargetBranch"])
				fmt.Println(ShaFileReqUrl)
				ShaFileResp:=MakeRequest("GET",ShaFileReqUrl,Auth,nil)
				fmt.Println(ShaFileResp)
				value,ok:= ShaFileResp.(map[string]interface{})
				
				if ok{
					ShaId=value["sha"]
				}
				RequestUrl:=fmt.Sprintf("%s/repos/%s/%s/contents/%s",githubApiBaseUrl,Org,Repo,FilePath)
				base64EncodedFileContent:=base64.StdEncoding.EncodeToString([]byte(data["FileContent"]))
				
				Payload:= map[string]interface{}{"message":"Changing the content","sha":ShaId,"content":base64EncodedFileContent,"branch":data["TargetBranch"]}
				PayloadBytes,err:=json.Marshal(Payload)
				if err != nil {
					json.NewEncoder(w).Encode(map[string]interface{}{"Error":err})
				}
				
				Response:=MakeRequest(r.Method,RequestUrl,Auth,bytes.NewBuffer(PayloadBytes))
				fmt.Println(Response)
				val,ok:=Response.(map[string]interface{})
				
				if ok{
					output,done:=val["commit"].(map[string]interface{})
					
					if done {
						//fmt.Println(val)
						json.NewEncoder(w).Encode(map[string]interface{}{"sha":output["sha"]})
					}else{
						json.NewEncoder(w).Encode(map[string]interface{}{"Error":"Could not update the file"})
					}
				}
					
					
		default:
			json.NewEncoder(w).Encode(map[string]string{"Error":"Method Not Supported"})

	}
}