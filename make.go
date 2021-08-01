package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var BinName string

func CheckDir() ([]string, []string, error) {
	var filestocopy []string
	var folderstocopy []string

	files, err := ioutil.ReadDir("./")

	if err != nil {
		return filestocopy, folderstocopy, err
	}

	for _, x := range files {

		checkpackage := true

		if x.IsDir() {
			f, _ := ioutil.ReadDir(x.Name())
			for _, y := range f {
				if !strings.Contains(y.Name()[len(y.Name())-3:len(y.Name())], ".go") {
					checkpackage = false
				}
			}
			if !checkpackage {
				folderstocopy = append(folderstocopy, x.Name())
			}

		} else if !strings.Contains(x.Name(), ".go") {
			if x.Name() != "go.mod" && x.Name() != "go.sum" {
				filestocopy = append(filestocopy, x.Name())
			}
		}
	}

	return filestocopy, folderstocopy, nil
}

func CheckBin(appname string) (bool, error) {

	files, err := ioutil.ReadDir("./")

	if err != nil {
		return false, err
	}

	check := false

	for _, x := range files {
		if !strings.Contains(x.Name(), ".") || strings.Contains(x.Name(), ".exe") {
			if !x.IsDir() {
				if x.Name() == appname {
					BinName = x.Name()
					check = true
				}
			}
		}
	}

	return check, nil
}

func Make(appname string) error {

	BinName = appname

	filestocopy, folderstocopy, err := CheckDir()

	if err != nil {
		return err
	}

	err = os.Mkdir("."+s+"__"+appname, 0755)

	if err != nil {
		return err
	}

	for _, x := range filestocopy {
		content, err := ioutil.ReadFile(x)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile("."+s+"__"+appname+s+x, content, 0755)
		if err != nil {
			return err
		}
	}

	for _, x := range folderstocopy {
		err = CopyDir(x, "."+s+"__"+appname+s+x)
		if err != nil {
			return err
		}
	}

	checkbin, err := CheckBin(appname)

	if err != nil {
		return err
	}

	if !checkbin {

		cmd := exec.Command("go", "build")

		err = cmd.Run()

		if err != nil {
			return err
		}

	}

	//Check if binary is on Windows

	files, err := ioutil.ReadDir("." + s)

	if err != nil {
		return err
	}

	for _, x := range files {
		if x.Name() == BinName+".exe" {
			BinName = BinName + ".exe"
		}
	}

	bin, err := ioutil.ReadFile(BinName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("."+s+"__"+appname+s+BinName, bin, 0755)

	if err != nil {
		return err
	}

	err = os.Remove(BinName)
	if err != nil {
		return err
	}

	err = os.Rename("."+s+"__"+appname, "."+s+appname)
	if err != nil {
		return err
	}

	return nil

}
