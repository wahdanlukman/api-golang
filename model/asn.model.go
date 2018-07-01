package model

import (
	"api-golang/myconfig"
	"io"
	"net/http"
	"os"
	"path"

	"api-golang/helper"
)

type pegawai struct {
	ID     int64
	Nik    string
	Nip    string
	Nama   string
	Alamat string
	Profil string
}

//AddAsn adalah proses data dari form add asn
//ADD
func AddAsn(w http.ResponseWriter, req *http.Request) {
	db, err := myconfig.GetMysqlConnect()
	defer db.Close()

	//Parsing form bertype multipart
	req.ParseMultipartForm(32 << 20)

	if req.Method == http.MethodPost {
		pegawai := pegawai{}
		pegawai.Nik = req.FormValue("nik")
		pegawai.Nip = req.FormValue("nip")
		pegawai.Nama = req.FormValue("nama")
		pegawai.Alamat = req.FormValue("alamat")
		pegawai.Profil = req.FormValue("alamat")
		helper.CheckErr(err)

		file, handler, err := req.FormFile("fileFoto")
		defer file.Close()

		ext := path.Ext(handler.Filename)

		newNameFile := pegawai.Nik + "-" + pegawai.Nama + ext

		f, err := os.OpenFile("./test/"+newNameFile, os.O_WRONLY|os.O_CREATE, 0666)
		helper.CheckErr(err)

		defer file.Close()
		io.Copy(f, file)

		_, err = db.Exec("INSERT INTO sim_asn(nik, nip, nama, alamat, profil) VALUES (?, ?, ?, ?, ?)",
			pegawai.Nik,
			pegawai.Nip,
			pegawai.Nama,
			pegawai.Alamat,
			newNameFile,
		)
		helper.CheckErr(err)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Harus menggunakan method post", http.StatusMethodNotAllowed)
}

//UpdateAsn untuk proses data dari form update
//UPDATE
func UpdateAsn(w http.ResponseWriter, req *http.Request) {
	db, err := myconfig.GetMysqlConnect()
	if req.Method == http.MethodPost {
		_, err = db.Exec("UPDATE sim_asn SET nik = ?, nip = ?, nama = ?, alamat = ? WHERE id = ?",
			req.FormValue("nik"),
			req.FormValue("nip"),
			req.FormValue("nama"),
			req.FormValue("alamat"),
			req.FormValue("id"),
		)
	}
	helper.CheckErr(err)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

//DeleteAsn untuk hapus data
//DELETE
func DeleteAsn(w http.ResponseWriter, req *http.Request) {
	db, err := myconfig.GetMysqlConnect()
	id := req.FormValue("id")
	db.Exec("DELETE FROM sim_asn WHERE id = " + id)

	helper.CheckErr(err)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
