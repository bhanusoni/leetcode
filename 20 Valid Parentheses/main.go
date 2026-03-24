func isValid(s string) bool {
    var list = []rune{}
    var pair = map[rune]rune{'{':'}', '(':')', '[': ']' }
    var size int
    for _, ch := range s {
        size = len(list)
        if val, ok := pair[ch]; ok{
            list = append(list, val)
        } else if size>0 && list[size-1] == ch{
            list = list[:size-1]
        } else {
            return false
        }

    }
    return len(list) == 0
}