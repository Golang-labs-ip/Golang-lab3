package tablets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Golang-labs-ip/Golang-lab3/server/tools"
)

// HTTPHandlerFunc is Balancers HTTP handler.
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of Balancers HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListTablets(store, rw)
		} else if r.Method == "PATCH" {
			handleTabletUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListTablets(store *Store, rw http.ResponseWriter) {
	res, err := store.ListTablets()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleTabletUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var tablet Tablet
	if err := json.NewDecoder(r.Body).Decode(&tablet); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateTablet(tablet.id)
	if err == nil {
		tools.WriteJsonOk(rw, &tablet)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}