package customerrors

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func HandleRenderError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	txt := fmt.Sprintf("Render error: %e.", err)
	render.PlainText(w, r, txt)
}

func RenderInternalError(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error.",
		ErrorText:      err.Error(),
	}

	if e := render.Render(w, r, &resp); e != nil {
		HandleRenderError(w, r, e)
	}
}

func RenderInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}

	if e := render.Render(w, r, &resp); e != nil {
		HandleRenderError(w, r, e)
	}
}
