package main
import(
	"net/http"
	"text/template"
)
var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("form.html","processor.html"))
}
func main(){
http.HandleFunc("/", indexfunc)
http.HandleFunc("/process",processor)
http.ListenAndServe(":9000",nil)
}

func indexfunc(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w,nil)
}
func processor(w http.ResponseWriter,r *http.Request){
	if r.Method!="POST"{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	fname:=r.FormValue("first_name")
	lname:=r.FormValue("last_name")
	company:=r.FormValue("company")
	email:=r.FormValue("email")
	phone:=r.FormValue("phone")
d:=struct{
	First string
	Last string
	Comp string
	Mail string
	Ph string
}{fname,lname,company,email,phone}
tpl.ExecuteTemplate(w,"processor.html",d)
}
