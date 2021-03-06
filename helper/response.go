package coffeego

import (
	"net/http"

	"github.com/fajrulaulia/rewritejson"
)

// ReponseOnServerError Should be Exported
func ReponseOnServerError(w http.ResponseWriter) {
	w.WriteHeader(500)
	output := rewritejson.Response([]string{"code:number|500", "message:string|Sorry, Internal Server Error, Please contact Administrator"})
	w.Write(output)
}

// ReponseOnNotFound Should be Exported
func ReponseOnNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	output := rewritejson.Response([]string{"code:number|404", "message:string|Sorry, You request not found"})
	w.Write(output)
}
