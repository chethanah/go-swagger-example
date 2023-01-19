package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// @title           API calculator
// @version         1.0
// @description     Calculate sum and prod for two numbers

// @contact.name   Chethan Suresh
// @contact.email  chethan.suresh@sony.com

// @host      43.88.80.127:8080
// @BasePath  /

type InputNumbers struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type Sum struct {
	SumRet int `json:"sum"`
}

type Prod struct {
	MulRet int `json:"product"`
}

// Add godoc
//
// @Summary Add two numbers.
// @Description Takes two numbers and finds the sum.
// @Tags calc
// @Produce json
// @Param num1 body InputNumbers true "addends"
// @Success 200 {object} Sum
// @Router /add [post]
func Add(w http.ResponseWriter, r *http.Request) {
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

// Mul godoc
//
// @Summary Products of two numbers
// @Description Takes two numbers and finds the product
// @Tags calc
// @Produce json
// @Param num1 body InputNumbers true "factors"
// @Success 200 {object} Prod
// @Router /mul [post]
func Mul(w http.ResponseWriter, r *http.Request) {

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
		retdata := &Prod{
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
	http.HandleFunc("/add", Add)
	http.HandleFunc("/mul", Mul)

	fmt.Printf("Starting calc API server on port :8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
