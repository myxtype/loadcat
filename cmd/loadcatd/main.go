// Copyright 2015 The Loadcat Authors. All rights reserved.

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/myxtype/loadcat/api"
	"github.com/myxtype/loadcat/cfg"
	"github.com/myxtype/loadcat/data"
	"github.com/myxtype/loadcat/feline"
	_ "github.com/myxtype/loadcat/feline/nginx"
	"github.com/myxtype/loadcat/ui"
)

func main() {
	fconfig := flag.String("config", "loadcat.conf", "")
	flag.Parse()
	err := cfg.LoadFile(*fconfig)
	if err != nil {
		log.Fatal(err)
	}

	err = feline.SetBase(filepath.Join(cfg.Current.Core.Dir, "out"))
	if err != nil {
		log.Fatal(err)
	}

	err = data.OpenDB(filepath.Join(cfg.Current.Core.Dir, "loadcat.db"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := data.DB.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	err = data.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api", api.Router)
	http.Handle("/", ui.Router)

	go func() {
		err = http.ListenAndServe(cfg.Current.Core.Address, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	log.Printf("Received %s; exiting..", <-c)
}
