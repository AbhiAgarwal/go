Setting up go
===

Setting up go is easy, but sometimes it gets a little tricky if you're not used to packages. 

**Set your $GOPATH:**

`mkdir ~/go`
`export GOPATH=~/go`

- Go commands all rely on one important environment variable which is called $GOPATH.
- Notice that this is not the $GOROOT where Go is installed.
- This variable points to the workspace of Go in your computer.
- In Windows, you need to create a new environment variable called GOPATH, then set its value to c:\mygo
- It is OK to have more than one path(workspace) in $GOPATH, but remember that you have to use `:`, and `;` in Windows to break up them.

**In $GOPATH, you must have three folders as follows:**

- `src` for source files whose suffix is .go, .c, .g, .s.
- `pkg` for compiled files whose suffix is .a.
- `bin` for executable files

Package directory
===

- Lets create a package name called `mymath`
- Create package source files and folders like `$GOPATH/src/mymath/`.
- Make a file called `sqrt.go` inside that folder
- Every time you create a package, you should create a new folder in the `src` directory, folders' name usually as same as the package's that you are going to use.
- You can have multi-level directory if you want to. 
    - For example, you create directories `$GOPATH/src/github.com/astaxie/beedb`, then the package path is `github.com/astaxie/beedb`. 
    - The package name will be the last directory in your path, which is beedb in this case.
- Run the following:
    - `cd $GOPATH/src`, and `mkdir mymath`
    - Inside that folder create a go file called `sqrt.go`

```go
// $GOPATH/src/mymath/sqrt.go

package mymath

func Sqrt(x float64) float64 {
    z := 0.0
    for i := 0; i < 1000; i++ {
        z -= (z*z - x) / (2 * x)
    }
    return z
}
```

- We've already created our package above, but how to compile it for practical? There are two ways to do it.

Compile packages
===

- Switch your work path to the directory of your package, then execute command `go install`.    
- Execute above command with file name like `go install mymath`.
- Look in `$GOPATH/pkg/___/`, where `___` could be `darwin_amd64`
- There should be `mymath.a`

Writing programs with it
===

- `cd ~/goprograms`
- Lets write a program to use the thing we just created

```go
package main

import (
    "mymath"
    "fmt"
)

func main(){        
    fmt.Println(mymath.Sqrt(2))
}
```

Now we can run this, and this should print out `1.414213562373095`.
