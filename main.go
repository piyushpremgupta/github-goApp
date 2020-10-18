package main 

import (
	"fmt"
	"net/http"
	"Accurics/handler/github"
	"Accurics/handler/common"
)




func main(){

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.Handle("/template/",http.StripPrefix("/template/",http.FileServer(http.Dir("template"))))
	http.HandleFunc("/GetShaFile/",github.GetShaFile)
	http.HandleFunc("/GetShaBranch/",github.GetShaBranch)
	http.HandleFunc("/ChangeFileContent/",github.ChangeFileContent)
	http.HandleFunc("/CreateBranch/",github.CreateBranch)
	http.HandleFunc("/CreatePullRequest/",github.CreatePullRequest)
	http.HandleFunc("/login",common.LoginHandler)
	http.HandleFunc("/logout",common.LogoutHandler)
	http.HandleFunc("/github",common.GithubHandler)
	http.HandleFunc("/auth",common.AuthHandler)
	

	fmt.Println("Server Started")
	http.ListenAndServe(":5001",nil)
	
}