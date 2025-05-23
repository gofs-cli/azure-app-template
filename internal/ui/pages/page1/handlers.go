package page1

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/gofs-cli/azure-app-template/internal/ui"
	"github.com/gofs-cli/azure-app-template/internal/ui/components/header"
)

func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ui.IndexPage(layout(header.Header(), body()))).ServeHTTP(w, r)
	})
}
