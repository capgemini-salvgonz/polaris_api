package commons

import (
	"hash/fnv"
	"regexp"
	"strconv"
	"strings"
)

// String is a custom type that represents a string.
type String string

// Append concatenates the given string to the current string and returns the result.
func (s *String) Append(other String) String {
	*s += other
	return String(*s)
}

// CharAt returns the character at the specified index of the string.
func (s String) CharAt(index int) rune {
	return []rune(s)[index]
}

// CompareIgnoreCase compares the current string with another string, ignoring case.
// It returns an integer comparing the two strings.
// The result will be 0 if the strings are equal, -1 if the current string is less than the other string,
// and 1 if the current string is greater than the other string.
func (s String) CompareIgnoreCase(other String) int {
	return strings.Compare(strings.ToLower(string(s)), strings.ToLower(string(other)))
}

// Concat concatenates the given string to the current string and returns the result.
func (s String) Concat(other String) String {
	return s + other
}

// Contains checks if the current string contains the specified substring.
// It returns true if the substring is found, otherwise false.
func (s String) Contains(other String) bool {
	return strings.Contains(string(s), string(other))
}

// Equals checks if the current string is equal to the specified string.
// It returns true if the strings are equal, otherwise false.
func (s String) Equals(other String) bool {
	return s == other
}

// EndsWidth checks if the current string ends with the specified suffix.
// It returns true if the string ends with the suffix, otherwise false.
func (s String) EndsWidth(suffix String) bool {
	return strings.HasSuffix(string(s), string(suffix))
}

// GetBytes returns the byte representation of the string.
func (s String) GetBytes() []byte {
	return []byte(s)
}

// GetChars returns the rune representation of the string.
func (s String) GetChars() []rune {
	return []rune(s)
}

// HashCode generates a hash value for the string using the FNV-1a algorithm.
// It returns a 32-bit unsigned integer representing the hash value.
func (s String) HashCode() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// IndexOf returns the index of the first occurrence of the specified substring in the string.
// If the substring is not found, it returns -1.
func (s String) IndexOf(substr String) int {
	return strings.Index(string(s), string(substr))
}

// IsEmpty checks if the current string is empty.
// It returns true if the string is empty, otherwise false.
func (s String) IsEmpty() bool {
	return s == ""
}

// IsNotEmpty checks if the current string is not empty.
// It returns true if the string is not empty, otherwise false.
func (s String) IsNotEmpty() bool {
	return s != ""
}

// LastIndexOf returns the index of the last occurrence of the specified substring in the string.
// If the substring is not found, it returns -1.
func (s String) LastIndexOf(substr String) int {
	return strings.LastIndex(string(s), string(substr))
}

// Length returns the length of the string.
// It returns an integer representing the number of characters in the string.
func (s String) Length() int {
	return len(s)
}

// Matches checks if the current string matches the specified regular expression.
// It returns true if there is a match, otherwise false.
func (s String) Matches(regex String) bool {
	match, _ := regexp.MatchString(string(regex), string(s))
	return match
}

// Replace replaces all occurrences of the old substring with the new substring in the string.
// It returns the resulting string after the replacement.
func (s String) Replace(old, new String) String {
	return String(strings.Replace(string(s), string(old), string(new), -1))
}

// Split splits the string into substrings using the specified separator.
// It returns a slice of strings containing the substrings.
func (s String) Split(separator string) []string {
	return strings.Split(string(s), separator)
}

// StartsWith checks if the current string starts with the specified prefix.
// It returns true if the string starts with the prefix, otherwise false.
func (s String) StartsWith(prefix String) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

// Substring returns a substring of the string starting from the specified start index (inclusive)
// and ending at the specified end index (exclusive).
func (s String) Substring(start, end int) String {
	return String(s[start:end])
}

// Trim removes leading and trailing white spaces from the string.
// It returns the trimmed string.
func (s String) Trim() String {
	return String(strings.TrimSpace(string(s)))
}

// ToLowerCase converts the string to lowercase.
// It returns the lowercase version of the string.
func (s String) ToLowerCase() String {
	return String(strings.ToLower(string(s)))
}

// ToUpperCase converts the string to uppercase.
// It returns the uppercase version of the string.
func (s String) ToUpperCase() String {
	return String(strings.ToUpper(string(s)))
}

// Value returns the string value.
func (s String) Value() string {
	return string(s)
}

// ValueOfInt returns the string representation of the specified integer.
func (s String) ValueOfInt(i int) String {
	return String(strconv.Itoa(i))
}

// ValueOfFloat returns the string representation of the specified float64.
func (s String) ValueOfFloat(f float64) String {
	return String(strconv.FormatFloat(f, 'f', -1, 64))
}

// ValueOfBool returns the string representation of the specified boolean.
func (s String) ValueOfBool(b bool) String {
	return String(strconv.FormatBool(b))
}

// ValueOfRune returns the string representation of the specified rune.
func (s String) ValueOfRune(r rune) String {
	return String(string(r))
}

// ValueOfBytes returns the string representation of the specified byte slice.
func (s String) ValueOfBytes(b []byte) String {
	return String(string(b))
}

// ValueOfChars returns the string representation of the specified rune slice.
func (s String) ValueOfChars(r []rune) String {
	return String(string(r))
}
