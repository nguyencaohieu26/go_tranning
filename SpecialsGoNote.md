- Go is not a pure object-oriented programming language and it does not support classes.
- Hence methods on types are a way to achieve behaviour similar to classes
- Build && Installation
  - Create executable file in Mac or linux we have to using syntax: GOOS="darwin" for Mac or GOOS="linux" for Ubuntu
  - Go build: command to compile the code into an executable
    - For windown we can run executable file by using syntax: .\interfaces.exe
    - For linux or Max(darwin) we can run executable file by using syntax: .\interfaces
- Memory Management
  - GC happens automatically
  - With new() keyword:
    - Allocate memory but no init
    - You will get a memory address
    - zeroed storage
  - With make() keyword:
    - Allocate memory and init
    - You will get a memory address
    - Non-zeroed storage
- Data Types
  - Basic Type: numbers, strings, boolean
  - Aggregate Type: array, structs
  - Reference Type: slice, map, function, chanel
  - Interface Type: interface
