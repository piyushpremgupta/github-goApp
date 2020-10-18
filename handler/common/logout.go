package common
import "net/http"

func LogoutHandler(w http.ResponseWriter , r *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		MaxAge: -1,
	})
	http.Redirect(w,r,"/login",303)
	w.Write([]byte(""))
}