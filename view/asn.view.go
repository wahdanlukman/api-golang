package view

import (
	"api-golang/helper"
	"api-golang/myconfig"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pegawai struct {
	ID     int64
	Nik    string
	Nip    string
	Nama   string
	Alamat string
	Profil string
}

//ListAllAsn unutk get data asn
//LIST
func ListAllAsn(w http.ResponseWriter, req *http.Request) {
	db, err := myconfig.GetMysqlConnect()
	defer db.Close()
	rows, err := db.Query("SELECT id, Nik, Nip, Nama, Alamat, Profil FROM sim_asn ORDER BY id DESC")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	asns := make([]pegawai, 0)
	for rows.Next() {
		pegawais := pegawai{}
		rows.Scan(&pegawais.ID, &pegawais.Nik, &pegawais.Nip, &pegawais.Nama, &pegawais.Alamat, &pegawais.Profil)
		asns = append(asns, pegawais)
	}
	tpl := template.Must(template.ParseGlob("templates/*"))
	tpl.ExecuteTemplate(w, "index.html", asns)
}

//AddAsn form untuk add data asn baru
func AddAsn(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*"))
	err := tpl.ExecuteTemplate(w, "asnForm.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//EditAsn form untuk edit data asn
func EditAsn(w http.ResponseWriter, req *http.Request) {
	db, err := myconfig.GetMysqlConnect()
	defer db.Close()

	id := req.FormValue("id")
	rows, err := db.Query("SELECT * FROM sim_asn WHERE id = " + id)

	helper.CheckErr(err)
	pegawais := pegawai{}
	for rows.Next() {
		rows.Scan(&pegawais.ID, &pegawais.Nik, &pegawais.Nip, &pegawais.Nama, &pegawais.Alamat, &pegawais.Profil)
	}
	tpl := template.Must(template.ParseGlob("templates/*"))
	tpl.ExecuteTemplate(w, "editAsn.html", pegawais)
	fmt.Println(id)
}
