- Definition:
  - Reflection is the ability of a program to inspect its variables and values at run time and find their type.
- What is the need for reflection?
  - Often times the data passed to these empty interfaces are not primitives -> we need to perform procedures on such data without knowing their type or the values present in it
  - In such a situation in order to perform various operations on the struct, such as interpreting the data present in it to query a database or create a shcema for a database we need to know the types present in it as well as the number of fields -> These problems can be dealt with during run-time using reflection.
- Some methods
  - reflect.ValueOf(x interface{}): Return list values in that argument
  - reflect.TypeOf(x interface{}):Return type of passed argument
  - Type.Kind(): data types are defined in go lang
  - NumField():This method returns the number of fields present in a struct. If the passed argument is not of the Kind reflect.Struct then it panics.
  - Field(): This method allows us to access each field in the struct using an Indexing variable.
