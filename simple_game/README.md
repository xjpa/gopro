notes

-- default value of booleans: false

-- small, big letters

- lowercase: private
- uppercase first letter: public

ex.

- isRunning - private
- IsRunning

-- pointers

& operator yields the address of a variable

the \* operator retrieves the variable that the pointer refers to

-- unsigned integers

uint -- unsigned integer, w/c means it cant be negative, only positive

uint32/uint64 -- uints can be 32 or 64-bit, 64 bits = 8 bytes

-- struct

anything that needs multiple values, you wanna group those values into structures

-- constructors

no concept of classical constructosr in golang, we need to do it ourselves

-- loops

```go
for i := 0; i < 10; i++ {
}
```

this one also works but loops indefinitely

```go
for{

}
```

42.07
