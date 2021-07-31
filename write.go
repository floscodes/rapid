package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func GetRouterModule() error {
	cmd := exec.Command("go", "get", "github.com/bouk/httprouter")
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}

func CreateDir(appname string) error {
	err := os.Mkdir(appname, 0755)
	return err
}

func InitGomod(appname string) error {

	if err := os.Chdir(appname); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", appname)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func WriteFiles(appname string) error {

	if err := ioutil.WriteFile("main.go", []byte(maingo), 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile("router.go", []byte(routergo), 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile("handlers.go", []byte(handlersgo), 0755); err != nil {
		return err
	}

	return nil

}

func EditGomod() error {
	c, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return err
	}
	Content := string(c)

	Content = Content[:strings.Index(Content, "// indirect")]

	err = ioutil.WriteFile("go.mod", []byte(Content), 0755)

	if err != nil {
		return err
	}

	return nil
}
