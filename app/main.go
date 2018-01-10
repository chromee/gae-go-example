package main

import (
	"fmt"
	"net/http"
	"log"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
	_ "go.mercari.io/datastore/aedatastore"
	"google.golang.org/appengine"
)


const (
	ProjectID = "morning-tide"
)

type TestData struct {
	ID int64 `json:"id" datastore:"-" boom:"id"`
	Name string `json:"name"`
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	opts := datastore.WithProjectID(ProjectID)
	ds, err := datastore.FromContext(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to create datastore client: %v", err)
		return
	}

	b := boom.FromClient(ctx, ds)
	data := &TestData{Name: "hello-local-datastore"}
	key, err := b.Put(data)
	if err != nil {
		log.Fatalf("Failed get from datastore: %v", err)
		return
	}
	fmt.Fprintf(w, "key: %s", key)
}
