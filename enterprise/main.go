package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"enterprise/controllers"
	"enterprise/models"

	log "github.com/inconshreveable/log15"
)

func main() {
	var (
		addr = envString("PORT", "8080")
	)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Put Controller Here
	company := new(controllers.CompanyController)
	company.Db = db

	//Put All Routes In Here
	router.GET("/H2O/company", company.Get)
	router.POST("/H2O/company/create", company.CreateCompany)
	router.POST("/H2O  /company/login", company.LoginCompany)

	router.Run(":" + addr)
	srvlog := log.New("module", "app/server")
	srvlog.Info("connection open")
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

/*
{

	{
    {
	 "Name":"harikrishna_enterprise",
    "Email": "sl.hari@gmail.com",
    "Password": "hfdu87347",
    "Telephone": "754581236",
    "Code":"127"

}

}
   {
    "created": {
        "Id": "000000000000000000000000",
        "Name": "harikrishna_enterprise",
        "Email": "sl.hari@gmail.com",
        "Password": "$2a$10$6UI7QBb3jp29sOo.t8dIbe2sNbKqInsh5IgVUVzDiJPU2wjzRkBXG",
        "Telephone": "754581236",
        "Code": "127",
        "Is_active": 0,
        "Is_verified": 0,
        "Last_login": null,
        "Remember_login": "",
        "CreatedAt": "2024-03-05T22:41:05.4339377+05:30",
        "UpdatedAt": "2024-03-05T22:41:05.4339377+05:30",
        "DeletedAt": null
    },
    "result": "ok!"
}

}*/
/*
create:
{

"Name":"erp_company",
   "Email": "slshjhdjs@gmail.com",
   "Password": "45323723",
   "Telephone": "516557",
    "Code":"525"

}
{
    "created": {
        "Id": 0,
        "Name": "erp_company",
        "Email": "slshjhdjs@gmail.com",
        "Password": "$2a$10$ZVr7e47NC8ySa3AQSprENevMHgkGW99TfuxhsqieFy.HMecDvQqVu",
        "Telephone": "516557",
        "Code": "525",
        "Is_active": 0,
        "Is_verified": 0,
        "Last_login": null,
        "Remember_login": "",
        "CreatedAt": "2024-03-12T13:56:55.0964284+05:30",
        "UpdatedAt": "2024-03-12T13:56:55.0964284+05:30",
        "DeletedAt": null
    },
    "result": "ok!"
}
*/
/*
Login:
{

   "Email": "slshjhdjs@gmail.com",
   "Password": "45323723"


}
{
    "login": null,
    "result": "email not exists"
}
*/
/*
Result:
{

   "Email": "slshjhdjs@gmail.com",
   "Password": "45323723"


}
{
    "result": []
}
*/
