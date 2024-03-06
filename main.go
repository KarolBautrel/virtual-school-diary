package main

import (
	"fmt"
	"sync"
	"virtual-diary/internal/class"
	"virtual-diary/internal/db"

	"gorm.io/gorm"
)

func prepareClassDomain(dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	classRepository := class.NewClassRepository(dbConn)
	classReadService := class.NewReadClassService(classRepository)
	fmt.Println(classReadService)
	// Prepare write service and routter for class
	fmt.Println("Initializing Domain")
}

func prepareStudentDomain(dbConn *gorm.DB, wg *sync.WaitGroup) {
	defer wg.Done()
	// Initialized domain of student
}

/// Initialize more domains like gradebook, teachers etc...

func main() {
	dbConn, err := db.DBConnector()
	if err != nil {
		fmt.Errorf("ERROR")
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go prepareClassDomain(dbConn, wg)
	go prepareStudentDomain(dbConn, wg)
	wg.Wait()
	///Listen to server after

}
