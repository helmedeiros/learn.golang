## How it works
When you write a program in Go you will have a `main` package defined with a `main` func inside it. The `func` keyword is how you define a function with a name and a body.

With `import "fmt"` we are importing a package which contains the `Println` function that we use to print.

## How to test
How do you test this? It is good to separate your "domain" code from the outside world (side-effects). The `fmt.Println` is a side effect (printing to stdout) and the string we send in is our domain.

Create a new function with `func` that returns a `string`.

Now create a new file called `hello_test.go` where we are going to write a test for our `Hello` function

Run `go test` in your terminal. It should've passed! Just to check, try deliberately breaking the test by changing the `want` string.

## Writing tests
Writing a test is just like writing a function, with a few rules:
- It needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`

For now it's enough to know that your `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.

## t.Errorf

We are calling the `Errorf` method on our `t` which will print out a message and fail the test. The `f` stands for format which allows us to build a string with values inserted into the placeholder values `%s`. When you made the test fail it should be clear how it works.
