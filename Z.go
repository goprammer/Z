package Z

// Returns first index of pattern if found in string.
func FirstIndex (pat, str string) int {
	total := pat + "$" + str
	length := len(total)
	z := make([]int, length)
	l := 0 
	r := 0

	for i := 1; i < length; i++ {
		if r < i {
			l := i 
			r := i

			for r < length && total[r] == total[r-l] {
				r++
			}

			z[i] = r-l
			r--
		} else {
			k := i - l 
			if z[k] < r - i + 1 {
				z[i] = z[k]
			} else {
				l = i

				for r < length && total[r] == total[r-l] {
					r++
				}

				z[i] = r-l
				r--
			}
		}
		
		if z[i] == len(pat) {
			return i - len(pat)
		}
	}	

	return -1
}

// Boolean check if pattern exists in string.
func Match (pat, str string) bool {
    res := FirstIndex(pat, str)
    if res != -1 {
        return true
    }
    
    return false
}
