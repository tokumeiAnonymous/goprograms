## This Readme is more of what I learn rather than how to use these packages
## This programs are build by following the tutorial from the book `Powerful Command-Line Applications in Go`
### Chapter 1
- How to initialize a go project
  ```
  go mod init module/name
  ```
- How to import dependencies
  ```
  import (
    "dependency1"
    "dependency2"
  )
  ```
- How to use inferred data type
  ```
  variable1 := "value1"
  ```
- Pointer vs Copy
  ```
  *p 
  ```
- Testing and the testing module
  ```
  t *testing.T
  m *testing.M
  ```
  All test files have `*_test.go` \
  To run the test `go test -v` and also make sure that you are in the directory where this test files exist.
- Build the app
  ```
  go build
  ```
- Command line flags
  ```
  lines := flag.Bool("l", false, "Count lines")
  // flag, default value, Description
  ```
- Build for different environment
  ```
  GOOS=windows go build
  ```
### Chapter 2
Since I know that there is a go extension for vscode I've tried to install it. (There is a vscode extension for everythin)
- How to use go.work
  - Since I used the same repository for multiple projects I am encountering errors
  - You can either simply ran `go work use ./module`. This will then create a `go.work` file.
  - You can also create the `go.work` file manually.
- Creating struct
  ```
  type item struct {
    Task string
    Done bool
  }
  ```
- When modifying an object copy that has pointer to the content. You can modify the underlying content. But when modifying the object copy directly you are modifying your own copy. If that makes sense
- When using the go extension in vscode upon saving it will remove all unused imports and automatically import unimported but used imports
- 