package router

// "webapp/src/router/routers"

import (
	"webapp/src/router/routers"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routers.Configurar(r)
}
