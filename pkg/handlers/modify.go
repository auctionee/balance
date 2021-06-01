package handlers

import (
	"encoding/json"
	"github.com/auctionee/balance/pkg/data"
	"github.com/auctionee/balance/pkg/db"
	"github.com/auctionee/balance/pkg/helpers"
	"github.com/auctionee/balance/pkg/logger"
	"io/ioutil"
	"net/http"
)

func ModifyHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger(r.Context())
	mod := data.Modify{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Println("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, mod)
	if err != nil {
		l.Println("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ok := helpers.ValidateToken(mod.Token); !ok {
		l.Println("bad token from", r.RemoteAddr)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer r.Body.Close()
	newBalance, err := db.Modify(mod)
	resp, err := json.Marshal(newBalance)
	if err != nil {
		l.Println("can't marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		l.Println("problem in /modify")
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
