# golang-united-school-homework-4

Implement a function that computes the sum of two int numbers written as a string

For example, having an input string "3+5", it should return output string "8" and nil error

Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")

For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace) the function should return an empty string and an appropriate error from strconv package wrapped into your own error with fmt.Errorf function

Use the provided errors appropriately for cases when the input string is empty(contains not character, but whitespace) or an expression contains one or greater than two operands, again wrapping into fmt.Errorf
