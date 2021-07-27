# Golang-rapid




Golang-rapid is a small and simple framework that helps you building WebApps in Go fast and easy.
(it is inspired by the Python-Framework Django)


***Before use make sure that Go is installed on your computer***





### Build from source:

1.) Download the Code as a ZIP-File. Unpack it, and open the folder in a shell. Then type

```
go build
```

2.) Move the compiled binary to your default bin-folder so that you can use it just by typing ```rapid``` in the shell.


### Creating a new app:

For creating a bew app use the following command:

```
rapid new YOUR_APP_NAME
```

Golang-rapid will create a new project-folder with your app’s source files:

* router.go   -   here you can set your http-paths and link them to a function that handles the request
* handlers.go -   here you can define the functions that handle the requests
* main.go     -   this file contains the main-function where the server is being started and where you can set the port your app will listen to


### Building the app:

To compile your webapp use the following command:

```
rapid make
```

Golang-rapidapp will create a new subfolder that contains the compiled executable and all the other files and subfolders of your project-folder except the .go-files.


### Cross-compiling:


If you want to cross-compile your app first use ```go build``` to set your preferred environment variables, e.g.
```
env GOOS=linux GOARCH=arm go build
```

An executable will be putted in your project-folder.
Then use 
```
rapid make
```

Golang-rapid will find your cross-compiled binary automatically and copy it to the new app folder without overwriting the existing cross-compiled executable.



### Info:
Golang-rapid uses the following repo to manage the routing of requests: [github.com/bouk/httprouter](https://github.com/bouk/httprouter/)
