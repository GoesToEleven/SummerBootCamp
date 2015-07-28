package stripe

import (
	"net/http"
	"html/template"
	"github.com/stripe/stripe-go"
)

var tpls *template.Template

func init(){
	tpls = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/payment", handlePayment)
}


func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/assets/imgs/minion.jpeg" {
		http.ServeFile(res, req, "minion.jpeg")
		return
	}
	tpls.ExecuteTemplate(res, "index.html", nil)
}

func handlePayment(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	stripeToken := req.FormValue()
}