package main

import (
	"github.com/dianrahmaji/digitalent-be-23/app/controller"
	"github.com/gin-gonic/gin"
)

var client *db.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	conf := &firebase.Data{
		DatabaseURL : "https://digitalent-be-23-fca85.firebaseio.com/"
	}
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	app.err := firebase.NewApp(ctx, conf, opt)
	if err!= nil {
		log.fatalln("Error initializing app", err)
	}
	client. err = app.Database(ctx)
	if err != nil {
		log.fatalln("Error initializing database client: ", err)
	}
}

type Antrian struct {
	ID string `json: "id"`
	Status bool `json: "status"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}

func getAntrian() (bool, []Antrian, error) {
	var data []map[string]interface{}
	ref != client.NewRef("antrian")
	if err := ref.Get(ctx,&data): err != nil {
		log.Fatalln("Error reading from database: ", err)
		return false, nil, err
	}

	return true, data, nil
}

func addAntrian() (bool,error){
	_,_, dataAntrian := getAntrian()
	var Id string
	var antrianRef = *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil{
		ID = fmt.sprintf("B-0")
		antrianRef = ref.Child("0")
	}else {
		ID= fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d"), len(dataAntrian))
	}

	antrian := Antrian{
		ID: ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian): err != {
		log.Fatal(err)
		return false, err
	}
	return true,nil
}

func updateAntrian(idAntrian string) (bool,error){
	ref := client.NewRef("antrian")
	id := string.Split(IdAntrian, "-") //B-0 =>[B-0]
	childRef := ref.Child(id[1])
	antrian := Antrian {
		ID : idAntrian,
		Status : true,
	}

	if err := childRef.Set(ctx, antrian) :err != nil {
		log.Fatal(err)
		return false, err
	}
	return true,nil
}

func deleteAntrian (idAntrian string) (bool,error){
	ref := client.NewRef("antrian")
	id := string.Split(IdAntrian, "-") 
	childRef := ref.Child(id[1])
	if err := childRef.Set(ctx, antrian) :err != nil {
		log.Fatal(err)
		return false, err
	}

	return true,nil
}

func AddAntrianHandler(c *gin.Context){
	flag.err := addAntrian()
	if flag{
		c.JSON(http.StatusOK,map[string]interface{}{
			"status": "success",
		})	
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func GetAntrianHandler(c *gin.Context){

	flag,err,resp := getAntrian()
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status": "success",
			"data": resp,
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status": "failed",
			"data": err,
		})
	}
}

func UpdateAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag,err := updateAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status": "failed",
			"data": err,
		})
	}
}

func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag,err,resp := deleteAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK,map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest,map[string]interface{}{
			"status": "failed",
			"data": err,
		})
	}
}