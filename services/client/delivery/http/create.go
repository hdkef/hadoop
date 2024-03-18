package http

import (
	"fmt"
	"net/http"

	"github.com/hdkef/hadoop/pkg/logger"
	"github.com/hdkef/hadoop/services/client/entity"
)

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

	// if not POST reject
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
		return
	}

	// set header SSE
	w.Header().Set("Content-Type", "text/event-stream")

	dto := &entity.CreateDto{}
	err := dto.NewFromHttp(r)
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	progressCh := make(chan entity.CreateStreamRes)

	go h.writeUC.Create(r.Context(), dto, progressCh)

	for val := range progressCh {

		if val.IsError() {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(val.GetError().Error()))
			logger.LogError(val.GetError())
			return
		}

		_, err = fmt.Fprintf(w, "%d %%", val.GetProgress())
		if err != nil {
			logger.LogError(err)
		}
		w.(http.Flusher).Flush()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
