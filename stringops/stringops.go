// Go program to illustrate how to split a string
package stringops

import (
	"fmt"
	"strings"
)

// Main function

// SplitAfterN: This function splits a string into all substrings after each instance of the given separator and returns a slice which contains these substrings.
// Syntax:
// func SplitAfterN(str, sep string, m int) []string
// Here, str is the string, sep is the separator, and m is used to find the number of substrings to return.
// Here, if m>0, then it returns at most m substrings and the last string substring will not split. If m == 0, then it will return nil.
// If m<0, then it will return all substrings.
func SplitAfterNTest() {

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using SplitAfterN() function
	res1 := strings.SplitAfterN(str1, ",", 2)
	res2 := strings.SplitAfterN(str2, "", 4)
	res3 := strings.SplitAfterN(str3, "!", 1)
	res4 := strings.SplitAfterN("", "GeeksforGeeks, geeks", 3)

	// Displaying the result
	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)

	//result
	// String 1:  Welcome, to the, online portal, of GeeksforGeeks
	// String 2:  My dog name is Dollar
	// String 3:  I like to play Ludo

	// Result 1:  [Welcome,  to the, online portal, of GeeksforGeeks]
	// Result 2:  [M y   dog name is Dollar]
	// Result 3:  [I like to play Ludo]
	// Result 4:  []

}
func SplitTest() {
	// Here, str is the string and sep is the separator.
	// If str does not contain the given sep and sep is non-empty, then it will return a slice of length 1 which contain only str.
	// Or if the sep is empty, then it will split after each UTF-8 sequence.
	// Or if both str and sep are empty, then it will return an empty slice.

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using Split() function
	res1 := strings.Split(str1, ",")
	res2 := strings.Split(str2, " ")
	res3 := strings.Split(str3, "!")
	res4 := strings.Split("", "GeeksforGeeks, geeks")

	// Displaying the result

	fmt.Println("\nResult 1: ", res1, "length :", len(res1))
	fmt.Println("Result 2: ", res2, "length :", len(res2))
	fmt.Println("Result 3: ", res3, "length :", len(res3))
	fmt.Println("Result 4: ", res4, "length :", len(res4))

	//result
	// String 1:  Welcome, to the, online portal, of GeeksforGeeks
	// String 2:  My dog name is Dollar
	// String 3:  I like to play Ludo

	// Result 1:  [Welcome  to the  online portal  of GeeksforGeeks] length : 4
	// Result 2:  [My dog name is Dollar] length : 5
	// Result 3:  [I like to play Ludo] length : 1
	// Result 4:  [] length : 1
}

func SplitAfterTest() {

	// Creating and initializing the strings
	str1 := "Welcome, to the, online portal, of GeeksforGeeks"
	str2 := "My dog name is Dollar"
	str3 := "I like to play Ludo"

	// Displaying strings
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)
	fmt.Println("String 3: ", str3)

	// Splitting the given strings
	// Using SplitAfter() function
	res1 := strings.SplitAfter(str1, ",")
	res2 := strings.SplitAfter(str2, "")
	res3 := strings.SplitAfter(str3, "!")
	res4 := strings.SplitAfter("", "GeeksforGeeks, geeks")

	// Displaying the result
	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
	//result:
	// String 1:  Welcome, to the, online portal, of GeeksforGeeks
	// String 2:  My dog name is Dollar
	// String 3:  I like to play Ludo

	// Result 1:  [Welcome,  to the,  online portal,  of GeeksforGeeks]
	// Result 2:  [M y   d o g   n a m e   i s   D o l l a r]
	// Result 3:  [I like to play Ludo]
	// Result 4:  []
}
