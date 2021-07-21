package source

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"schneider.vip/problem"
)

type routerHandler struct {
	repo Repository
}

func NewRouter(repo Repository) *mux.Router {
	handler := routerHandler{repo}

	r := mux.NewRouter()
	r.HandleFunc("/fdns", handler.Get).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	return r
}

func (h routerHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Path: " + r.URL.Path)
	defer r.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Info().Msg("Body: " + bodyString)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func (h routerHandler) Get(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		log.Error().Msg("error missing query parameter")
		problem.Of(http.StatusBadRequest).Append(
			problem.Title("Invalid Payload"),
			problem.Detail("There's validation error on your content, check the documentation about this contract."),
			problem.Type("http://ourDocLink"), //TODO: refactor
			problem.Custom("invalid_params", []interface{}{"domain", "is required"}),
		).WriteTo(w)
	}

	fdns, err := h.repo.Select(domain)
	if err != nil {
		log.Error().Err(err).Msg("error querying database")
		problem.Of(http.StatusInternalServerError).Append(
			problem.Title(http.StatusText(http.StatusInternalServerError)),
			problem.Detail("The server encountered an unexpected condition which prevented it from fulfilling the request."),
			problem.Type("contact us at suporte@xpto"), //TODO: refactor
		).WriteTo(w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fdns)
}
