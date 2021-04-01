package main

import (
	"crud-mysql-gorilla-mux/app/config"
	"crud-mysql-gorilla-mux/app/controller"
	"crud-mysql-gorilla-mux/app/models"
	"crud-mysql-gorilla-mux/app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	conect := config.InitMysql()

	if err := conect.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Db sudh terkoneksi")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", GetHome).Methods("GET")
	router.HandleFunc("/api/v1/mahasiswa", GetAllDataMhs).Methods("GET")
	router.HandleFunc("/api/v1/mahasiswa", GetCreateDataMhs).Methods("POST")
	router.HandleFunc("/api/v1/mahasiswa/id={id}", GetIdDataMhs).Methods("GET")
	router.HandleFunc("/api/v1/mahasiswa/id={id}", GetUpdateDataMhs).Methods("PUT")
	router.HandleFunc("/api/v1/mahasiswa/id={id}", GetDeleteDataMhs).Methods("DELETE")

	log.Println("Server berjalan di 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func HttpInfo(r *http.Request) {
	fmt.Printf("%s/\t %s\t %s%s\t %s\n", r.Method, r.Proto, r.Host, r.URL, utils.GetDateTime())
}

func GetHome(w http.ResponseWriter, r *http.Request) {

	utils.ResponseJson(w, map[string]interface{}{
		"data": "ini adalah home",
	})
}

func GetAllDataMhs(w http.ResponseWriter, r *http.Request) {
	HttpInfo(r)
	resData, err := controller.GetAllData()
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.ResponseJson(w, map[string]interface{}{
		"data": resData,
	})
}

func GetCreateDataMhs(w http.ResponseWriter, r *http.Request) {
	HttpInfo(r)
	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		log.Fatal(err)
		return
	}

	if err := controller.CreateData(mhs); err != nil {
		log.Fatal(err)
		return
	}

	utils.ResponseJson(w, map[string]interface{}{
		"pesan": "data berhasil dibuat",
	})

}

func GetIdDataMhs(w http.ResponseWriter, r *http.Request) {
	HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	resData, err := controller.GetIdData(id)
	if err != nil {
		log.Println("Id tidak sesuai")
		return
	}
	utils.ResponseJson(w, map[string]interface{}{
		"data": resData,
	})
	return
}

func GetUpdateDataMhs(w http.ResponseWriter, r *http.Request) {
	HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
		log.Fatal(err)
		return
	}
	mhs.ID = id
	if err := controller.GetUpdateData(mhs); err != nil {
		log.Fatal(err)
		return
	}

	utils.ResponseJson(w, map[string]interface{}{
		"data": "Data Berhasil diubah",
	})
}

func GetDeleteDataMhs(w http.ResponseWriter, r *http.Request) {
	HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := controller.GetDeleteData(id); err != nil {
		log.Fatal(err)
		return
	}

	utils.ResponseJson(w, map[string]interface{}{
		"Pesan": "data berhasil dihapus",
	})
}
