package main

import (
	"log"
	"os"
)

func main() {
	f, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	log.SetOutput(f)
	defer f.Close()

	log.Println("test")
	log.Println("test")
	log.Println("test")

	//f, err := os.OpenFile("/tmp/orders.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer f.Close()
	//wrt := io.MultiWriter(os.Stdout, f)
	//log.SetOutput(wrt)
	//log.Println(" Orders API Called")
}
