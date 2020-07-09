# Homework

The homework consists of Go source files including unimplemented functions.
Your job is to implement these functions based on their function comments and get the unit tests passing.
You will find `// TODO` comments in all the places you need to add code.

## Get Started

Please check out a branch with the following format: `yourname/202007`

For example, my branch would be: `neil/202007`

Fill in the empty functions in your branch and push it to origin when you are done.

## Install Go

To do the homework you will need to install Go 1.14 on your local machine.
You can do this with your package manager of choice.
Or, if you prefer, you can download it as a tar following these instructions: https://golang.org/doc/install

## How to verify your answers locally

To check your answers to one directory of problems, run all the unit tests in that directory.
The unit tests can be run with `go test`.

For example, you can run the tests for `001-slices` with this command (from the root dir of this Git repo):

```
go test 202007/001-slices/*
```

For the tests in `008-channels` you may want to use the `-timeout` flag.
Because these tests are concurrent, if there is a bug in your code it may lead to a deadlock.
Here's an example with a timeout of 10 seconds (which should be reasonable).

```
go test -timeout 10s 202007/008-channels/*
```

## When you're done...

When you finish the homework, push your branch open a pull request to the `master` branch.
(Don't forget to `go fmt` your code first.)
The instructor will check your code and leave comments - or let you know if the code looks good.
If you need help or can't complete any problems, open a pull request and let the instructor know.
The instructor will check your code and leave comments to help you solve your problem.
