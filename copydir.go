package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CopyDir(src string, dest string) error {

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	file, err := f.Stat()
	if err != nil {
		return err
	}
	if !file.IsDir() {
		return fmt.Errorf("Source " + file.Name() + " is not a directory!")
	}

	err = os.Mkdir(dest, 0755)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range files {

		if f.IsDir() {

			err = CopyDir(src+s+f.Name(), dest+s+f.Name())
			if err != nil {
				return err
			}

		}

		if !f.IsDir() {

			content, err := ioutil.ReadFile(src + s + f.Name())
			if err != nil {
				return err

			}

			if !strings.Contains(f.Name()[len(f.Name())-3:len(f.Name())], ".go") {

				err = ioutil.WriteFile(dest+s+f.Name(), content, 0755)
				if err != nil {
					return err

				}
			}

		}

	}

	return nil
}
