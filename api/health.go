package api

import (
	"errors"
	"fmt"
	"github.com/InVisionApp/rye"
	"net/http"
)

func (a *Api) healthHandler(rw http.ResponseWriter, r *http.Request) *rye.Response {
	if a.Deps.Postgres.IsAlive() {
		fmt.Fprint(rw, "OK")
		return nil
	}
	return &rye.Response{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New("postgres is not alive"),
	}
}
