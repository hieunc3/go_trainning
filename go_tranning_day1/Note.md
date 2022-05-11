1> Go variable
    + Create variable
        - Using var keyword
            - Syntax: var variable_name type = expression
            - When you don't know the initial value
            - When you need a package scoped variable
            - When you want to group variables togather for greater readability
            - Important Point: 
                - In the above syntax, either type or = expression can be ommited, but not both
                - If the "= expression" is omiited, then the variable is determined by its type's default value
                - If the type is removed, then the type of the variable is determined by the value-initialize in the expression
        - Using short variable declaration
            - Syntax: variable_name:= expression     
            - When you know the initial value
            - Keep the code concise
            - For redeclaration
    + Default value 
        - 0 - Number
        - false - Booleans
        - "" - strings
        - nil - interface
        - nil - reference type

2> Data Type
    + List Data Type:
      - Basic Type: Numbers, strings, and booleans.
        - Strings:
          - Strings are the immutable chain of arbitrary bytes or string is a read-only slice of 
          - It is a sequence of variable-width characters where each and every character is represented by one or more bytes using UTF-8 Encoding
          - The len() function returns the number of bytes of the string
          - Rune:
            - Rune represents a Unicode code point in Go
            - The RuneCountInString() function returns the total number of rune presents in the string
      - Aggregate type: Array and structs
        - Array
          - Store a collection of data of the same type
          - An array is a fixed-length sequence that is used to store the same type element
          - Important Point: 
            - An array is of value type not of reference type
            - The length of the array is fixed and unchangeable.
            - In an array, if the element type of the array is comparable, then the array type is also comparable
            - In an array, if an array does not initialized explicitly, then the default value of the array is the default value of the type array
            - Arrays are mutable, so that you can use array[index] syntax to the left-hand side of the assignment to set the elements of the array at the given index
      - Reference type: Pointers, slices, maps, functions and channels
        - Slices
          - The size of the slice is resized they are not in fixed-size just like an array
          - A slice is a reference to an underlying array
          - When a slice is created, the program allocates an array and returns a slice that refers to that array
          - It consists of a pointer to the array, the length of the segment, and its capacity
          - The slice does not copy the slice's data. It creates a new slice value that points to the original array
          - To increase the capacity of a slice one must creaate a new, larger slice and copy the contents of the original slice into it.
      - Interface type
3> Encode & Decode from Json
    + Encode:
        - Using Struct Tags to control Encoding:
            - The struct tage that encoding/json recognizes has a key of json and a value that controls the output.
            -  By placing the camel-cased version of the field names as the value to the json key, the encoder will use that name instead
        - Using `json: "name_change",omitempty"` 
            - omitempty: show empty if the field is nil
            - "-" ignore the field