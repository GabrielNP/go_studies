# Go lang studies repo

## Set up

## How to run

## Personal Doc about Go

As Go team defined, Go lang has a default Go Workspace at $HOME path. So Go lang always waits to `go/` folder and inside it `src/`, `pkg/` and `bin/` folders at Go workspace path.

This default path can be changed, but all your Go applications must be on some defined `Go Workspace/src`.

See [quickstart](quickstart/) folder

### Commands

|Command|Description|
|------|-------|
|go build <'filename.go'>|compile go file source to a new executable file|
|go run <'filename'.go>|compile and run the new executable file|

### Conventions

Go lang has some conventions and good pratices to avoid problems and arguments.

|Convention|Description|
|------|-------|
|package main|Is the main application's package|
|func main()|The main function of a program must be called by that name|
|;|It's not necessary to end some instruction. On the contrary, they ask to not use in these cases|
|break|It's not necessary on switch statements. But if exists compiler does not care|

#### Hints:

There are some VS code extensions that automatically clean some bad pratices. Take a look on "Go by Go Team at Google" and accept all its dependencies.



## References

[Go Doc](www.golang.org)

https://cursos.alura.com.br/course/golang