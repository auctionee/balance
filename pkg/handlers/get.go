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

func GetHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.GetLogger(r.Context())
	creds := data.Credentials{}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Println("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, creds)
	if err != nil {
		l.Println("can't read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if ok := helpers.ValidateToken(creds.Token); !ok {
		l.Println("bad token from", r.RemoteAddr)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer r.Body.Close()

	info := db.GetInfo()
	resp, err := json.Marshal(info)
	if err != nil {
		l.Println("can't marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		l.Println("problem in /get")
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
