package main

import (
	"encoding/binary"
	"fmt"
)

func main(){

    data := []byte{1, 2}
    fmt.Println(data)
    binary.LittleEndian.PutUint16(data, 1)
    fmt.Println(data)
    fmt.Println(binary.LittleEndian.Uint16(data))
    // fmt.Println(byte(1))
    // fmt.Println(byte(1 >> 8))
    // fmt.Println("%T", binary.LittleEndian.Uint16(data))

    // 1. Hello World
    // helloWorld()

    // 2. Variables
    // values()

    // 3. Constants
    // constants()

    // 4. Loops
    // loop()

    // 5. Conditionals
    // conditions()

    // 6. Switch case
    // switchCase()

    // 7. array examples
    // arrayEx()

    // 8. Slices
    // slicey()

    // 9. Maps
    // mapEx()

    // 10. Range
    // rangeEx()

    // 11. Functions
    // callFunc()

    // 12. Variadic functions
    // variadic()

    // 13. Closures
    // closure() 

    // 14. Recursion
    // recursion()

    // 15. Pointers
    // pointers()

    // 16. Strings and Runes
    // stringNRunes()

    // 17. Structs
    // structs()

    // 18. Methods

    // 19. Interfaces

    // 20. Struct Embedding

    // 21. Generics

    // 22. Errors

    // 23. Custom Errors

    // 24. Goroutines

    // 25. Channels

    // 26. Channel Buffering

    // 27. Channel Synchronization

    // 28. Channel Directions

    // 29. Select

    // 30. Timeouts

    // 31. Non-Blocking Channel Operations
    
    // 32. Closing Channels

    // 33. Timers

    // 34. Tickers
    
    // 35. Worker Pools

    // 36. WaitGroups

    // 37. Rate Limiting (api stuff ?)
    
    // 38. Atomic Counters

    // 39. Mutexes

    // 40. Stateful Goroutines

    // 41. Sorting
    // sortIt() 

    // 42. Sorting by Functions
    // sortByFunc()

    // 43. Panic
    // panicC()

    // 44. Defer
    // BegToDefer() 

    // 45. Recover
    // LetsRecover() 

    // 46. String Functions
    // stringF()

    // 47. String Formatting

    // 48. Text Templates
   
    // 49. Regular expressions
    // RegexCheck()

    // 50. JSON
    // jsonEx()

    // 51. XML

    // 52. Time 
    // timeTrial()

    // 53. Epoch
    // epoch()

    // 54. Time Formatting / Parsing
    // timeParse()  

    // 55. Random Numbers

    // 56. Number Parsing

    // 59. URL Parsing

    // 60. SHA256 Hashes
    // shasha()

    // 61. Base64 Encoding
    // baseAndURLEnc()

    // 62. Reading files

    // 63. Writing files

    // 64. Line Filters

    // 65. File Paths

    // 66. Directories

    // 69. Temporary Files and Directories

    // 70. Embed Directive

    // 71. Testing and Benchmarking

    // 72. Command-line arguments

    // 72. Command-line flags

    // 73. Command-line subcommands

    // 74. Environement Variables

    // 75. Logging

    // 76. HTTP Client

    // 77. HTTP Server

    // 78. Context

    // 79. Spawning Processes

    // 80. Exec'ing Processes

    // 81. Signals

    // 82. Exit
    // exitProg()

    // Misc
    // misc1()
}
