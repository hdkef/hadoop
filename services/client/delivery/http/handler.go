package http

import (
	"net/http"

	"github.com/hdkef/hadoop/services/client/config"
	"github.com/hdkef/hadoop/services/client/usecase"
)

type handler struct {
	writeUC usecase.WriteUsecase
}

func NewHTTPHandler(cfg *config.Config, w usecase.WriteUsecase) *http.ServeMux {

	if cfg == nil {
		panic("nil config")
	}

	mux := http.NewServeMux()

	h := &handler{
		writeUC: w,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/create", h.Create)

	return mux
}
