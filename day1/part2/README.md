# Notes

## Idea

- left list - just a normal list
- right list - a map[key]occurence


## Flow

1. populate left (list) and right (map)
2. Iterate through values of left
3. Use left's value as a key to get right map's value

## Learning (Not useful in this exercise lol)

### Set in Go

Go doesn't have a built-in `set` data structure, but can use a `map` with empty struct values (`map[T]struct{}`) to efficiently implement a Set. This approach is idiomatic to Go and makes use of the fact that `map` keys must be unique.

> Why use `struct{}`? `struct{}` has zero memory overhead compared to other placeholder types like `bool` and `int`. It communicates that the value is irrelevant, only the presence of key matters.

Example:

```Go
func main() {
    // create a set
    set := make(map[string]struct{})

    // add elements to the set
    set["apple"] = struct{}{} // an instace (literal) of the empty struct type `struct{}`
    set["banana"] = struct{}{}

    // check if exists
    _, exists := set["apple"]

    // delete an element
    delete(set, "banana")

    // iterate over it
    for key := range set {
        fmt.Println(key)
    }
}
```

### Method Receivers recap

In Go, methods are functions associated with specific type. They are two types of receivers:

1. Pointer Receiver `(* Type)`

    - The method operates on a pointer to the type.
    - Allows the method to modify the original value.
    - Common when working with large structs or when you need to mutate the receiver.

2. Value Receiver `(Type)`

    - The method operates on a copy of the type
    - The original value remains unchanged
    - Suitable for small structs or when mutation isn't required

