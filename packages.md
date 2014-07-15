# Go notes

- One thing that you should know in the first is that Go programs are composed by package
- `package<pkgName>` (In this case is `package main`) 
    - tells us this source file belongs to main package
    - the keyword main tells us this package will be compiled to a program instead of package files whose extensions are `.a`

- Every executable program has one and only one main package, and you need an entry
function called `main` without any argument and return value in the main package.
- The function main.main() (this function must be in the main package) is the entry point of any program.
- Go supports UTF-8 characters because one of the creators of Go is a creator of UTF-8, so Go supports multi-language from the time it was born.
