package main

import (
	"fmt"
	"log"
	"net/http"
	"virtual-diary/internal/class"
	"virtual-diary/internal/db"
	"virtual-diary/internal/student"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")
	dbConn, err := db.DBConnector()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	router := mux.NewRouter()
	classRepository := class.NewClassRepository(dbConn)
	classService := class.NewClassService(dbConn, classRepository)
	class.RegisterRoutes(router, classService)
	studentRepo := student.NewStudentRepo(dbConn)
	studentService := student.NewStudentService(dbConn, studentRepo)
	student.RegisterRoutes(router, studentService)

	log.Fatal(http.ListenAndServe(":8080", router))

}
