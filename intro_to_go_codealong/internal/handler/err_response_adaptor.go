package handler

import (
	"github.com/sirupsen/logrus"
	"intro_to_go_codealong/internal/itgerr"
	"net/http"
)

type ErrHandler func(w http.ResponseWriter, r *http.Request) error

func ErrRespAdaptor(next ErrHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err != nil {
			sts, msg := getErrDetails(err)
			w.WriteHeader(sts)
			w.Write([]byte(msg))

			logrus.WithError(err).Error("error in HTTP handler")
		}
	})
}

func getErrDetails(err error) (int, string) {
	k, m := itgerr.GetKind(err)

	switch k {
	case itgerr.KindNotFound:
		return 404, m
	case itgerr.KindInvalidInput:
		return 400, m
	}

	return 500, "internal server error"
}
