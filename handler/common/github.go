package common


import(
	"fmt"
	"net/http"
	_"encoding/json"
	"html/template"
	)
func GithubHandler(w http.ResponseWriter , r *http.Request){
	tmpl, err := template.ParseFiles("template/common/github.html","template/layout/header.html","template/layout/footer.html")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(tmpl.ExecuteTemplate(w, "github", ""))	
}