# Z
A simple library for O(n) pattern matching and search.

## func Match
`Match(pat, str string) bool`   
If **pat** exists in **str**, **Match** returns true at the first occurance.


## func FirstIndex
`FirstIndex(pat, str string) int`   
If **pat** exists in **str**, **FirstIndex** returns the index of the start of the pattern at the first occurance.     

Example:   
    
    pat := "cat"    
    str := "I see cattle and a cat."
    
    fIdx := FirstIndex(pat, str)
    fmt.Println(fIdx)
    
Output:      
`6`



