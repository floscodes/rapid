package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

//String s is the path-separator. The value depends on the OS. (On Windows it's vlaue would be "\")
var s string

func main() {

	//Check if rapid is running on Windows. If so, then set path-separator s to "\"

	switch os := runtime.GOOS; os {
	case "windows":
		s = `\`
	default:
		s = "/"
	}

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

		appname = appname[strings.LastIndex(appname, s)+1:]

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
