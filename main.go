package main

import (
	"fmt"
	"log"
	"net/http"

	//External Packages
	_ "github.com/go-sql-driver/mysql"

	//Internal Packages

	"api-golang/api"
	"api-golang/model"
	"api-golang/myconfig"
	"api-golang/view"
)

func main() {
	db, err := myconfig.GetMysqlConnect()
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("GOLANG API")
	http.HandleFunc("/", view.ListAllAsn)
	http.HandleFunc("/asnForm", view.AddAsn)
	http.HandleFunc("/prosesAsn", model.AddAsn)
	http.HandleFunc("/editAsn", view.EditAsn)
	http.HandleFunc("/deleteAsn", model.DeleteAsn)
	http.HandleFunc("/updateAsn", model.UpdateAsn)

	// API ROUTES
	http.HandleFunc("/api/", api.ListAllAsnAPI)
	http.HandleFunc("/api/asn/", api.GetAsnAPI)

	log.Println("Serevr running at port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
