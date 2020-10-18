package github

import (
	_"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
)



func MakeRequest(Method,RequestUrl,Auth string ,Payload io.Reader) interface{}{
	Client:=&http.Client{}
	req,err:=http.NewRequest(Method,RequestUrl,Payload)
	if err != nil{
		return map[string]interface{}{"Error":err}
	}
	req.Header.Add("Content-type","application/json")
	req.Header.Add("Authorization",Auth)
	resp,err2:=Client.Do(req)
	
	if err2 != nil{
		
		return map[string]interface{}{"Error":err2}
	}
	defer resp.Body.Close()

	RespBytes,err3:= ioutil.ReadAll(resp.Body)
	if err3 != nil{
		
		return map[string]interface{}{"Error":err3}
	}

	var data interface{}

	err4:= json.Unmarshal(RespBytes,&data)
	if err4 != nil{
		
		return map[string]interface{}{"Error":err4}
	}
	
	
	switch RespValue:=data.(type){
		case map[string]interface{}:
			return RespValue
		case []interface{}:
			return RespValue
		default:
			return RespValue
	}
	
}