package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"virtual-diary/internal/class"
	"virtual-diary/internal/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func prepareClassDomain(router *mux.Router, dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	classRepository := class.NewClassRepository(dbConn)
	classReadService := class.NewReadClassService(classRepository)
	class.RegisterRoutes(router, classReadService)

}

func prepareStudentDomain(router *mux.Router, dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	// Initialized domain of student
}

/// Initialize more domains like gradebook, teachers etc...

func main() {
	dbConn, err := db.DBConnector()
	if err != nil {
		fmt.Errorf("ERROR")
	}
	router := mux.NewRouter()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go prepareClassDomain(router, dbConn, wg)
	go prepareStudentDomain(router, dbConn, wg)
	wg.Wait()
	///Listen to server after
	log.Fatal(http.ListenAndServe(":8080", router))

}
