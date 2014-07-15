**Number**

- Note that `rune is alias of int32` and `byte is alias of uint8`.
- Float types have `float32` and `float64`, and no type called float.
- Go has complex number as well. 
    - `complex128` (with a 64-bit real and 64-bit imaginary part) 
    - `complex64` (with a 32-bit real and 32-bit imaginary part)   
    - `RE + IMi`, 
        - RE is real part
        - IM is imaginary part, the last i is imaginary number.

**Strings** 

- We just talked about that Go uses UTF-8 character set.
    - Strings are represented by double quotes "" or backtracks `` .
- It's impossible to change string values by index, you will get errors when you compile following code.
- Change one String:

```go
s := "hello"
c := []byte(s) // convert string to []byte type
c[0] = 'c'
s2 := string(c) // convert back to string type
fmt.Printf("%s\n", s2)
```

**Error types**

- Go has one error type for purpose of dealing with error messages. There is also a package called errors to handle errors. 

```go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
    fmt.Print(err)
}
```

