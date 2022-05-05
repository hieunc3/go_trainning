- Interfaces
  - In Go, an interface is set of method signatures.
  - It's possible to create new interfaces by embedding other interfaces
  - When a type provides definition for all the methods in the interface, it's said to implement the interface.
  - Interface specifies what methods a type should have and the type decides how to implement these methods.
  - This is not needed in Go and Go interfaces are implemented implicitly if a type contains all the methods declared in the interface.
  - Type assertion is used to extract the underlying value of the interface
  - Implement interfaces using pointer receivers & value receivers:
    - It's legal to call a value method on anything which is a value or whose value can be dereferenced