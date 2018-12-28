package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"plugin"
)

type RPlugin interface {
	RunPlugin(string) string
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := r.FormValue("uid")
	plugin_name := ps.ByName("plugin")
	mod := "./plugin/" + plugin_name + "/" + plugin_name + ".so"
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	symPlugin, err := plug.Lookup("RPlugin")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	var rPlugin RPlugin
	rPlugin, ok := symPlugin.(RPlugin)
	if !ok {
		fmt.Fprintf(w, "unexpected type from module symbol")
	}
	res := rPlugin.RunPlugin(uid)
	fmt.Fprint(w, res)
}

func main() {
	router := httprouter.New()
	router.GET("/:plugin", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
