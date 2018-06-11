package handler

import (
	"fmt"
	"html/template"
	"net/http"

	us "github.com/asepnur/iskandar/src/module/users"
	t "github.com/asepnur/iskandar/src/webserver/template"
	"github.com/julienschmidt/httprouter"
)

var (
	emptyTime  = "0001-01-01 00:00:00"
	layoutTime = "2006-01-02 15:04:05"
)

// ViewHTML ::
func ViewHTML(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("src/webserver/template/users.html")
	el, err := us.GetVisitor()
	us.IncreaseVisitor(fmt.Sprintf("%d", el))
	data := map[string]interface{}{
		"Visitor": el,
	}
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(w, data)
	return
}

// TestingHTML func
func TestingHTML(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t.RenderJSONResponse(w, new(t.Response).
		SetCode(http.StatusOK).
		SetData("halo"))
	return
}

// SelectUserHandler ..
func SelectUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error

	q := r.FormValue("name")
	resp := []User{}
	params := UserParams{
		Q: q,
	}
	resp, code, err := params.SelectUserController()
	if err != nil {
		t.RenderJSONResponse(w, new(t.Response).
			SetCode(code))
		return
	}
	t.RenderJSONResponse(w, new(t.Response).
		SetCode(code).
		SetData(resp))
	return

}
