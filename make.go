package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func CheckDir() ([]string, []string, error) {
	var filestocopy []string
	var folderstocopy []string

	files, err := ioutil.ReadDir("./")

	if err != nil {
		return filestocopy, folderstocopy, err
	}

	for _, x := range files {

		if x.IsDir() {
			folderstocopy = append(folderstocopy, x.Name())

		} else if !strings.Contains(x.Name(), ".go") {
			if x.Name() != "go.mod" && x.Name() != "go.sum" {
				filestocopy = append(filestocopy, x.Name())
			}
		}
	}

	return filestocopy, folderstocopy, nil
}

func CheckBin() (bool, error) {

	files, err := ioutil.ReadDir("./")

	if err != nil {
		return false, err
	}

	check := false

	for _, x := range files {
		if !strings.Contains(x.Name(), ".") {
			check = true
		}
	}

	return check, nil
}

func Make(appname string) error {

	filestocopy, folderstocopy, err := CheckDir()

	if err != nil {
		return err
	}

	err = os.Mkdir("__"+appname, 0755)

	if err != nil {
		return err
	}

	for _, x := range filestocopy {
		content, err := ioutil.ReadFile(x)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("./__"+appname+"/"+x, content, 0755)
		if err != nil {
			return err
		}
	}

	for _, x := range folderstocopy {
		err = CopyFolder(x, "./__"+appname+"/"+x)
		if err != nil {
			return err
		}
	}

	checkbin, err := CheckBin()

	if err != nil {
		return err
	}

	if checkbin {

		cmd := exec.Command("go", "build")

		err = cmd.Run()

		if err != nil {
			return err
		}

	}

	bin, err := ioutil.ReadFile(appname)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./__"+appname+"/"+appname, bin, 0755)

	if err != nil {
		return err
	}

	err = os.Remove(appname)
	if err != nil {
		return err
	}

	err = os.Rename("./__"+appname, "./"+appname)
	if err != nil {
		return err
	}

	return nil

}
