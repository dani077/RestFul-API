package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sellingGorilla/repositories"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (initial *App) Run(strAddress string) {
	log.Fatal(http.ListenAndServe(":8898", initial.Router))
}
func (intial *App) Initializer(username, password, dbname string) {
	var err error
	intial.DB, err = sql.Open("mysql", "root:root@/DBSelling")
	if err != nil {
		log.Fatal(err)
	}
	intial.Router = mux.NewRouter()
	intial.InitialRoute()
}

func (initial *App) InitialRoute() {
	initial.Router.HandleFunc("/item", initial.gettblItem).Methods("Get")
	initial.Router.HandleFunc("/officer", initial.GetOfficer).Methods("Get")
	initial.Router.HandleFunc("/selling", initial.GetSelling).Methods("Get")
	initial.Router.HandleFunc("/detail", initial.GetDetail).Methods("Get")
	initial.Router.HandleFunc("/transaksi", initial.GetTrans).Methods("Get")
	initial.Router.HandleFunc("/update/{tblItemID:[0-9]+}", initial.UpdatetblItem).Methods("PUT")
	initial.Router.HandleFunc("/delete/{tblItemID:[0-9]+}", initial.deleteItem).Methods("DELETE")
	initial.Router.HandleFunc("/insert", initial.insertItem).Methods("POST")
}

/*func (initial *App) GetItem(w http.ResponseWriter, r *http.Request) {
	i := repositories.Item{}
	if err := i.GetItem(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, i)
}*/
func (init *App) gettblItem(w http.ResponseWriter, r *http.Request) {
	data, err := repositories.GetItem(init.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}
func (initial *App) GetOfficer(w http.ResponseWriter, r *http.Request) {
	o := repositories.Officer{}
	if err := o.GetOfficer(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, o)
}
func (initial *App) GetSelling(w http.ResponseWriter, r *http.Request) {
	s := repositories.Selling{}
	if err := s.GetSelling(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, s)
}
func (initial *App) GetDetail(w http.ResponseWriter, r *http.Request) {
	d := repositories.Detail{}
	if err := d.GetDetail(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, d)
}
func (initial *App) GetTrans(w http.ResponseWriter, r *http.Request) {
	t := repositories.Transaksi{}
	if err := t.GetTrans(initial.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No Record Found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, t)
}

func (a *App) UpdatetblItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["tblItemID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	var i repositories.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	i.TblItemID = id

	if err := i.UpdateItem(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (a *App) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["tblItemID"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Item ID")
		return
	}

	p := repositories.Item{TblItemID: id}
	if err := p.DeleteItem(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
func (a *App) insertItem(w http.ResponseWriter, r *http.Request) {

	var i repositories.Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&i); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := i.IItem(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, i)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	respond, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(respond)

}
