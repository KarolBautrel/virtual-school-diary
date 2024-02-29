package main

import (
	"log"
)

func prepareClassDomain(classDomainInitialized chan bool) {
	//gonna init class domain here
	classDomainInitialized <- true
	close(classDomainInitialized)
}

func prepareStudentDomain(studentDomainInitialized chan bool) {
	//gonna init student domain here
	studentDomainInitialized <- true
	close(studentDomainInitialized)
}

func main() {

	studentChannel := make(chan bool)
	classChannel := make(chan bool)

	go prepareClassDomain(classChannel)
	go prepareStudentDomain(studentChannel)
	_, studentChannelClosed := <-classChannel
	_, classChannelClosed := <-studentChannel
	if !studentChannelClosed || !classChannelClosed {
		log.Fatalf("Initialization failed")
	}

}
