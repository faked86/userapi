package responses

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func RenderInternalError(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error.",
		ErrorText:      err.Error(),
	}

	RenderErrResponse(w, r, resp)
}

func RenderInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}

	RenderErrResponse(w, r, resp)
}

func RenderNotFound(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     "Not found.",
		ErrorText:      err.Error(),
	}

	RenderErrResponse(w, r, resp)
}

func RenderErrResponse(w http.ResponseWriter, r *http.Request, resp ErrResponse) {
	if e := render.Render(w, r, &resp); e != nil {
		HandleRenderError(w, r, e)
	}
}

func HandleRenderError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	txt := fmt.Sprintf("Render error: %e.", err)
	render.PlainText(w, r, txt)
}
