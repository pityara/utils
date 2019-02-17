package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpJSONRequest(w http.ResponseWriter, str interface{}) {
	res, err := json.Marshal(str);
	if err != nil {
		log.Fatal("Error occurs while trying marshall json")
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println(res)

	w.Write(res)
}
