package main

import (
	"log"

	"github.com/ShavqatKavrakov/Lesson16_Export/pkg/wallet"
)

func main() {
	srv := &wallet.Service{}
	_, err := srv.RegisterAccount("918925874")
	if err != nil {
		log.Print(err)
	}
	_, err = srv.Deposit(1, 10_00000)
	if err != nil {
		log.Print(err)
	}
	_, err = srv.RegisterAccount("987026424")
	if err != nil {
		log.Print(err)
	}
	_, err = srv.Deposit(2, 20_00000)
	if err != nil {
		log.Print(err)
	}
	path := "data/massege.txt"
	err = srv.ExportToFile(path)
	if err != nil {
		log.Print(err)
	}
	err = srv.ImportToFile(path)
	if err != nil {
		log.Print(err)
	}
}
