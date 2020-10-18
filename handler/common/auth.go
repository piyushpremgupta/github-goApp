package common

import (
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/base64"
	github "Accurics/handler/github"
	"time"
	"bytes"
	"strings"
	"math/rand"
)

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
	"abcdefghijklmnopqrstuvwxyzåäö" +
	"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"
	return str
}
func basicAuth(username,password string) string{
	auth:= username+ ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))

}
func AuthHandler(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-type","application/json")
	var cred map[string]string

	json.NewDecoder(r.Body).Decode(&cred)

	Payload:=map[string]interface{}{"scopes":[]string{"public_repo"},"note":RandomString()}
	PayloadBytes,err:=json.Marshal(Payload)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"Error":err})
	}

	Response:=github.MakeRequest("POST","https://api.github.com/authorizations","Basic "+basicAuth(cred["UserName"],cred["Password"]),bytes.NewBuffer(PayloadBytes))
	fmt.Println(Response)
	value,ok:=Response.(map[string]interface{})
	if ok{
		val,present:=value["token"].(string)
		if present{
			
			//adding a cookie 
			expire:= time.Now().AddDate(0,0,1)
			cookie:= http.Cookie{
				Name: "Token",
				Value: val,
				Expires:expire,
			}
			http.SetCookie(w,&cookie)
			json.NewEncoder(w).Encode(map[string]string{"Authenticate":"True"})
		}else{
			json.NewEncoder(w).Encode(map[string]string{"Error":"Could not generate token"})
		}
	}

}

