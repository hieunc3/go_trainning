1> Function Arguments

- Functions of the same name but different type are not allowed to be defined in the program
- By default Go language use call by value method to pass argument in function
- Go language supports two ways to pass arguments to your function:
  - Call by value:
    - In this parameter passing method, values of actual parameters are copied to function's formal
      parameters and the two types of parameters are stored in different memory locations
      -> So any changes made inside functions are not reflected in actual parameters of the caller.
    - Call by reference: - Both the actual and formal parameters refer to the same locations, so any changes made inside the function are actually reflected in actual parameters of the caller
- Variadic function:
  - The function that called with the varying number of arguments is known as variadic funcdtion
  - In the declaraction of the variadic function, the type of the last parameters is preceded by an ellipsis(...). the ...type behaves like a slice
  - Variadic function is used when you want to pass a slice in a function
  - Variadic function is used when we don’t know the quantity of parameters.
  - When you use variadic function in your program, it increase the readability of your program.
- Anonymous function: (function literal)
  - A function is which doesn't contain any name.
  - It's useful when you want to create an inline function
  - Go language allows to assign an anonymous function to a variable.
  - You can also pass an anonymous function as an argument into other function
