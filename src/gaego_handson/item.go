package gaego_handson

import (
	"encoding/json"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/pborman/uuid"
)

// Item
type Item struct {
	KeyStr    string    `json:"key" datastore:"-"`
	Title     string    `json:"title" datastore:",noindex"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ItemApi struct {
}

func SetUpItem(m *http.ServeMux) {
	api := ItemApi{}

	m.HandleFunc("/api/1/item", api.handler)
}

func (a *ItemApi) handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		a.doPost(w, r)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (a *ItemApi) doPost(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var param Item
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	item := Item{
		Title:     param.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	key, err := datastore.Put(c, datastore.NewKey(c, "Item", uuid.New(), 0, nil), &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	item.KeyStr = key.Encode()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
