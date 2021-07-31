package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(manual)
		return
	}

	if os.Args[1] == "new" {

		if os.Args[2] != "" && !Checkspace(os.Args[2]) {

			appname := os.Args[2]

			if err := CreateDir(appname); err != nil {
				log.Fatal(err)
			}

			if err := InitGomod(appname); err != nil {
				log.Fatal(err)
			}

			if err := GetRouterModule(); err != nil {
				log.Fatal(err)
			}

			if err := WriteFiles(appname); err != nil {
				log.Fatal(err)
			}

			if err := EditGomod(); err != nil {
				log.Fatal(err)
			}

		}

		return

	}

	if os.Args[1] == "make" {
		appname, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		appname = appname[strings.LastIndex(appname, "/")+1:]

		if err := Make(appname); err != nil {
			log.Fatal(err)
		}

		return

	}

	if os.Args[1] != "new" && os.Args[1] != "make" {
		fmt.Println(manual)
		return
	}

}
