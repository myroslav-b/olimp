package server

import (
	"fmt"
	"log"
	"net/http"
	"olimp/cmd/olimp/catalogs"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type iStoreRequester interface {
	QueryInstitutions(instType, regCode string, valFields, reqFields map[string]string) ([]map[string]string, error)
}

type TServer struct {
	Addr   string
	Engine iStoreRequester
}

type tJSON map[string]interface{}

func init() {

}

func (s TServer) Run() error {
	log.Printf("[INFO] run server ")

	err := http.ListenAndServe(s.Addr, s.routes())
	if err != nil {
		return errors.Wrap(err, " failed")
	}
	return nil
}

func (s TServer) routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID, middleware.RealIP)
	router.Use(middleware.Throttle(1000), middleware.Timeout(60*time.Second))

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/register/{inst}/{reg}", s.getInstList)
		r.Get("/catalog/{name}", s.getCatalog)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("Ok"))
	})

	return router
}

func (s TServer) getInstList(w http.ResponseWriter, r *http.Request) {

	fmt.Println(">>>    ", r.Context())
	inst, reg := chi.URLParam(r, "inst"), chi.URLParam(r, "reg")

	if inst == "" || reg == "" {
		log.Print("[WARN] inst or reg is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, tJSON{"error": "inst or reg is empty"})
		return
	}

	request := func() *struct {
		instType, regCode    string
		valFields, reqFields map[string]string
	} {
		s := struct {
			instType, regCode    string
			valFields, reqFields map[string]string
		}{inst, reg, make(map[string]string), make(map[string]string)}

		for fieldInst := range catalogs.MapFieldInst() {
			req := r.URL.Query().Get("r." + fieldInst)
			if req != "" {
				s.reqFields[fieldInst] = req
			}
			val := r.URL.Query().Get("v." + fieldInst)
			if val != "" {
				s.valFields[fieldInst] = val
			}
		}
		return &s
	}()

	//fmt.Println(*request)

	listInsts, err := func(source iStoreRequester) ([]map[string]string, error) {
		li, e := source.QueryInstitutions(request.instType, request.regCode, request.valFields, request.reqFields)
		return li, e
	}(s.Engine)
	if err != nil {
		log.Print("[WARN] failed to get list of institutions ")
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, tJSON{"error": "Failed to get list of institutions "})
	}
	//p := r.URL.Query().Get("v.p")
	render.Status(r, http.StatusOK)
	//render.JSON(w, r, tJSON{"inst": inst, "reg": reg, "p": request.reqFields})
	render.JSON(w, r, listInsts)
	return
}

func (s TServer) getCatalog(w http.ResponseWriter, r *http.Request) {

	name := chi.URLParam(r, "name")
	if name == "" {
		log.Print("[WARN] catalog name is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, tJSON{"error": "catalog name is empty"})
		return
	}

	switch name {
	case "fi":
		render.Status(r, http.StatusOK)
		render.JSON(w, r, catalogs.MapFieldInst())
	case "it":
		render.Status(r, http.StatusOK)
		render.JSON(w, r, catalogs.MapInstType())
	case "rc":
		render.Status(r, http.StatusOK)
		render.JSON(w, r, catalogs.MapRegCode())
	case "tr":
		render.Status(r, http.StatusOK)
		render.JSON(w, r, catalogs.MapTypeRequest())
	default:
		log.Print("[WARN] catalog name is wrong")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, tJSON{"error": "catalog name is wrong"})
	}

	return
}
