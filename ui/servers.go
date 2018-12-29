// Copyright 2015 The Loadcat Authors. All rights reserved.

package ui

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/myxtype/loadcat/data"
	"github.com/myxtype/loadcat/feline"
)

func ServeServerNewForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	bal, err := data.GetBalancer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = TplServerNewForm.Execute(w, struct {
		Balancer *data.Balancer
	}{
		Balancer: bal,
	})
	if err != nil {
		panic(err)
	}
}

func HandleServerCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	bal, err := data.GetBalancer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	body := struct {
		Label    string `schema:"label"`
		Settings struct {
			Address string `schema:"address"`
		} `schema:"settings"`
	}{}
	err = schema.NewDecoder().Decode(&body, r.PostForm)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	srv := data.Server{}
	srv.BalancerId = bal.Id
	srv.Label = body.Label
	srv.Settings.Address = body.Settings.Address
	err = srv.Put()
	if err != nil {
		panic(err)
	}

	err = feline.Commit(bal)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/servers/"+srv.Id.Hex()+"/edit", http.StatusSeeOther)
}

func HandleServerDelForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	srv, err := data.GetServer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = TplServerDelForm.Execute(w, struct {
		Server         *data.Server
		Availabilities []data.Availability
	}{
		Server:         srv,
		Availabilities: data.Availabilities,
	})
	if err != nil {
		panic(err)
	}
}

func HandleServerDel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	srv, err := data.GetServer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = srv.Del()
	if err != nil {
		panic(err)
	}

	bal, err := srv.Balancer()
	if err != nil {
		panic(err)
	}
	err = feline.Commit(bal)
	if err != nil {
		//panic(err)
	}

	http.Redirect(w, r, "/balancers/"+srv.BalancerId.Hex(), http.StatusSeeOther)
}

func ServeServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	srv, err := data.GetServer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/servers/"+srv.Id.Hex()+"/edit", http.StatusSeeOther)
}

func ServeServerEditForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	srv, err := data.GetServer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = TplServerEditForm.Execute(w, struct {
		Server         *data.Server
		Availabilities []data.Availability
	}{
		Server:         srv,
		Availabilities: data.Availabilities,
	})
	if err != nil {
		panic(err)
	}
}

func HandleServerUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	srv, err := data.GetServer(bson.ObjectIdHex(vars["id"]))
	if err != nil {
		panic(err)
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	body := struct {
		Label    string `schema:"label"`
		Settings struct {
			Address      string `schema:"address"`
			Weight       int    `schema:"weight"`
			Availability string `schema:"availability"`
		} `schema:"settings"`
	}{}
	err = schema.NewDecoder().Decode(&body, r.PostForm)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	srv.Label = body.Label
	srv.Settings.Address = body.Settings.Address
	srv.Settings.Weight = body.Settings.Weight
	srv.Settings.Availability = data.Availability(body.Settings.Availability)
	err = srv.Put()
	if err != nil {
		panic(err)
	}

	bal, err := srv.Balancer()
	if err != nil {
		panic(err)
	}
	err = feline.Commit(bal)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/balancers/"+srv.BalancerId.Hex(), http.StatusSeeOther)
}

func init() {
	Router.NewRoute().
		Methods("GET").
		Path("/balancers/{id}/servers/new").
		Handler(http.HandlerFunc(ServeServerNewForm))
	Router.NewRoute().
		Methods("POST").
		Path("/balancers/{id}/servers/new").
		Handler(http.HandlerFunc(HandleServerCreate))
	Router.NewRoute().
		Methods("GET").
		Path("/servers/{id}").
		Handler(http.HandlerFunc(ServeServer))
	Router.NewRoute().
		Methods("GET").
		Path("/servers/{id}/edit").
		Handler(http.HandlerFunc(ServeServerEditForm))
	Router.NewRoute().
		Methods("POST").
		Path("/servers/{id}/edit").
		Handler(http.HandlerFunc(HandleServerUpdate))
	Router.NewRoute().
		Methods("GET").
		Path("/servers/{id}/del").
		Handler(http.HandlerFunc(HandleServerDelForm))
	Router.NewRoute().
		Methods("POST").
		Path("/servers/{id}/del").
		Handler(http.HandlerFunc(HandleServerDel))
}
