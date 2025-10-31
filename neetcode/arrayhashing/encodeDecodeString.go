package arrayhashing

import "strconv"

type Solution struct{}

/*
Only a pointer (s *Solution) receiver can modify the struct's fields directly.
For example, if the Solution struct had fields,
a pointer receiver could change their values.

without pointer receiver (s Solution) would create a copy of it and would not
make any changes to the origin Solution
*/
func (s *Solution) Encode(strs []string) string {
	/*
	   In Go, strconv.Itoa is a function from the strconv package that converts an integer to its string representation.

	   The name Itoa stands for Integer to ASCII:

	   Itoa: Converts an integer (int) to a string.
	   The result is the string representation of the number.
	   import "strconv"

	   func main() {
	       num := 42
	       str := strconv.Itoa(num) // Convert integer 42 to string "42"
	       fmt.Println(str)         // Output: "42"
	   }


	*/

	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str

	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	// encoded = "4#abcd2#gh3#ijk" , 4##abc 4#a4bc
	// [abcd, gh, ijk]
	res := []string{}

	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++

		}
		length, _ := strconv.Atoi(encoded[i:j]) // second return value is error

		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length

	}
	return res

}
