#rapidapp

First make sure that Go is installed on your computer. Then donwload these files and make a binary using "go build". copy it to a bin-folder so that you can use it as a command in the shell.



RAPIDAPP USAGE:


Command: "rapidapp new APPNAME"       for creating a folder with your app’s source files

Command: "rapidapp make"              for compiling your app and create a new folder with the binary and all the related data


If you want to cross-compile your app use the „go build“-command and set your preferred environment variables, e.g. "env GOOS=linux GOARCH=arm go build"
Then use the "rapidapp make"-command. rapidapp will find your cross-compiled binary automatically and copy it to the new app folder.


