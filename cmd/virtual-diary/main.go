package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"virtual-diary/internal/auth"
	"virtual-diary/internal/class"
	"virtual-diary/internal/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func prepareClassDomain(router *mux.Router, dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	classRepository := class.NewClassRepository(dbConn)
	classReadService := class.NewReadClassService(classRepository)
	classWriteService := class.NewWriteClassService(classRepository)
	class.RegisterRoutes(router, classReadService, classWriteService)

}

func prepareAuthDomain(router *mux.Router, dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	authRepository := auth.NewAuthRepository(dbConn)
	authReadService := auth.NewReadService(authRepository)
	auth.RegisterRoutes(router, authReadService)
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

	domains := []func(*mux.Router, *gorm.DB, *sync.WaitGroup){
		prepareClassDomain,
		prepareStudentDomain,
		prepareAuthDomain,
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(domains))
	for _, domain := range domains {
		go domain(router, dbConn, wg)
	}
	wg.Wait()
	///Listen to server after
	log.Fatal(http.ListenAndServe(":8080", router))

}
