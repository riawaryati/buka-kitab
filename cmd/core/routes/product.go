package routes

import (
	"net/http"

	"github.com/buka-kitab/backend/domain/general"
	"github.com/buka-kitab/backend/handlers/core"
	"github.com/gorilla/mux"
)

func getProductData(router, routerJWT *mux.Router, conf *general.SectionService, handler core.Handler) {
	routerJWT.HandleFunc("/author", handler.Product.Author.GetListAuthor).Methods(http.MethodGet)
	routerJWT.HandleFunc("/category", handler.Product.Category.GetListCategory).Methods(http.MethodGet)
	routerJWT.HandleFunc("/form", handler.Product.Form.GetListForm).Methods(http.MethodGet)
	routerJWT.HandleFunc("/language", handler.Product.Language.GetListLanguage).Methods(http.MethodGet)
	routerJWT.HandleFunc("/publisher", handler.Product.Publisher.GetListPublisher).Methods(http.MethodGet)
}
