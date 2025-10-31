function generateParanthesis(n) {}

// first = (
// last = )
// n = 1 , first last
// n = 2, (first first last last), (first last first last)
// n = 3, (first first first last last last), (first last first last first last),
// (first first last first last last)

// a = (, b = )
// n = 2, a-b, c-d = abcd, acbd
// n = 3, a-b, c-d, e-f = abcdef, acebdf, acdbef, abcefd, acdefb
