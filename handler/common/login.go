package common
import(
"fmt"
"net/http"
_"encoding/json"
"html/template"
)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/common/login.html","template/layout/header.html","template/layout/footer.html")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(tmpl.ExecuteTemplate(w, "login", ""))	
}