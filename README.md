#rapidapp


rapidapp is a small and simple framework that helps you building WebApps in Go fast and easy.

First make sure that Go is installed on your computer. Then donwload these files and make a binary using "go build". copy it to a bin-folder so that you can use it as a command in the shell.

rapidapp uses the followinf repo to manage the routing of requests: github.com/bouk/httprouter



RAPIDAPP USAGE:


Command: "rapidapp new APPNAME"       for creating a folder with your app’s source files

This command will generate a router.go-file where you can set your paths an link them to a handler-function, a handlers.go-file where you can definde your handler-functions and a main.go-file where you can set the port that the app will listen to and the server is being started.

Command: "rapidapp make"              for compiling your app and create a new folder with the binary and all the related directories


If you want to cross-compile your app use the „go build“-command and set your preferred environment variables, e.g. "env GOOS=linux GOARCH=arm go build"
Then use the "rapidapp make"-command. rapidapp will find your cross-compiled binary automatically and copy it to the new app folder.


