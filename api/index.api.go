package api

import (
	"api-golang/helper"
	"api-golang/myconfig"
	"encoding/json"
	"log"
	"net/http"
)

type pegawai struct {
	ID     int64  `json:"id"`
	Nik    string `json:"nik"`
	Nip    string `json:"nip"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Profil string `json:"profil"`
}

//ListAllAsnAPI unutk get data asn
//LIST
func ListAllAsnAPI(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == http.MethodGet {
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
		result, err := json.Marshal(asns)
		helper.CheckErr(err)

		w.Write(result)
		return
	}
}

//GetAsnAPI unutk get data asn berdasarkan id
//LIST
func GetAsnAPI(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method == http.MethodGet {
		id := req.FormValue("id")
		db, err := myconfig.GetMysqlConnect()
		defer db.Close()
		rows, err := db.Query("SELECT id, Nik, Nip, Nama, Alamat, Profil FROM sim_asn WHERE id = " + id)
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
		result, err := json.Marshal(asns)
		helper.CheckErr(err)

		w.Write(result)
		return
	}
}
