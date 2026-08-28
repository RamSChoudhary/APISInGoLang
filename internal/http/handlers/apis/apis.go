package apis

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/GoLang/APIs/internal/http/types"
	"github.com/GoLang/APIs/internal/http/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {

			response.WriteJson(w, http.StatusBadRequest, err.Error())

			return
		}

		slog.Info("creating a student")

		//w.Write([]byte("welcome to go lang apis"))

		validator.New().Struct(student)

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
