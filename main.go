package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

type InputNumbers struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Sum struct {
	SumRet int `json:"sum"`
}

type Mul struct {
	MulRet int `json:"product"`
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Addition API")
	case "POST":
		d := json.NewDecoder(r.Body)
		req := &InputNumbers{}
		err := d.Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sum := req.Num1 + req.Num2
		retdata := &Sum{
			SumRet: sum,
		}

		retdatajson, err := json.MarshalIndent(retdata, "", " ")
		if err != nil {
			log.Println("Unable to marshal output")
		}
		w.Write(retdatajson)

	default:
		fmt.Fprintf(w, "/add - Sorry, only GET and POST methods are supported.")
	}
}

func mul(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Multiply API")
	case "POST":
		d := json.NewDecoder(r.Body)
		req := &InputNumbers{}
		err := d.Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		mul := req.Num1 * req.Num2
		retdata := &Mul{
			MulRet: mul,
		}

		retdatajson, err := json.MarshalIndent(retdata, "", " ")
		if err != nil {
			log.Println("Unable to marshal output")
		}
		w.Write(retdatajson)

	default:
		fmt.Fprintf(w, "mul/ - Sorry, only GET and POST methods are supported.")
	}

}

func main() {

	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			accounts.POST("", c.AddAccount)
			accounts.DELETE(":id", c.DeleteAccount)
			accounts.PATCH(":id", c.UpdateAccount)
			accounts.POST(":id/images", c.UploadAccountImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

	http.HandleFunc("/add", add)
	http.HandleFunc("/mul", mul)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
