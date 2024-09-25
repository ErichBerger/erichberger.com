package templates

import (
	"io"
	"net/http"

	"github.com/a-h/templ"
)

func HXRender(w io.Writer, r *http.Request, solo templ.Component, full templ.Component) error {
	hx := r.Header.Get("HX-Request")
	var err error
	switch hx {
	case "":
		err = full.Render(r.Context(), w)
	case "true":
		err = solo.Render(r.Context(), w)
	}

	if err != nil {
		return err
	}
	return nil
}
