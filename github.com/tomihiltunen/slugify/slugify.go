/***************************************************************
*
*   Slugify
*   The Go-lang slug generator.
*   
*   Created by Tomi Hiltunen 2013.
*   http://www.linkedin.com/in/tomihiltunen
*
*   https://github.com/TomiHiltunen/Slugify
*
*       - Use this script however you wish.
*       - Do not remove any copyrights/comments on any files included.
*       - All use is on your own risk.
*
*   Note worthy:
*       - Use this to generate slugs to embed in URLs for SEO purposes.
*       - Do NOT use for generating unique identifiers for resources.
*
***************************************************************/
package slugify

import (
	// Go packages
	"regexp"
	"strings"
)

/*
 *  Dictionary of accented characters and their translitterations.
 */
var dictionary = map[string]string{
	"Š": "S",
	"š": "s",
	"Đ": "Dj",
	"đ": "dj",
	"Ž": "Z",
	"ž": "z",
	"Č": "C",
	"č": "c",
	"Ć": "C",
	"ć": "c",
	"À": "A",
	"Á": "A",
	"Â": "A",
	"Ã": "A",
	"Ä": "A",
	"Å": "A",
	"Æ": "A",
	"Ç": "C",
	"È": "E",
	"É": "E",
	"Ê": "E",
	"Ë": "E",
	"Ì": "I",
	"Í": "I",
	"Î": "I",
	"Ï": "I",
	"Ñ": "N",
	"Ò": "O",
	"Ó": "O",
	"Ô": "O",
	"Õ": "O",
	"Ö": "O",
	"Ø": "O",
	"Ù": "U",
	"Ú": "U",
	"Û": "U",
	"Ü": "U",
	"Ý": "Y",
	"Þ": "B",
	"ß": "Ss",
	"à": "a",
	"á": "a",
	"â": "a",
	"ã": "a",
	"ä": "a",
	"å": "a",
	"æ": "a",
	"ç": "c",
	"è": "e",
	"é": "e",
	"ê": "e",
	"ë": "e",
	"ì": "i",
	"í": "i",
	"î": "i",
	"ï": "i",
	"ð": "o",
	"ñ": "n",
	"ò": "o",
	"ó": "o",
	"ô": "o",
	"õ": "o",
	"ö": "o",
	"ø": "o",
	"ù": "u",
	"ú": "u",
	"û": "u",
	"ý": "y",
	"þ": "b",
	"ÿ": "y",
	"Ŕ": "R",
	"ŕ": "r",
}

/*
 * Creates a lower-case trimmed string with dashes for white-spaces.
 *
 *      - Converts to lower case.
 *      - Trims the leading/trailing white spaces.
 *      - Converts applicable accented caharcters to non-accented and removes invalid ones.
 *      - Converts leftover white-spaces to dashes regardles of type.
 */
func Slug(original string) (edited string) {
	// Remove invalid characters
	re, _ := regexp.Compile(`[^A-Za-z0-9\s-_]`)
	edited = re.ReplaceAllStringFunc(original, convertAccent)
	// Trim leading and trailing white-space
	edited = strings.TrimSpace(edited)
	// Convert all white-spaces to dashes
	re, _ = regexp.Compile(`\s+`)
	edited = re.ReplaceAllString(edited, "_")
	// All done!
	return strings.ToLower(edited)
}

/*
 * Same as above, but shorteness the string if it is over the size limit.
 */
func SlugWithMaxLength(original string, length int) (edited string) {
	// Create the slug
	edited = Slug(original)
	// Cut to size
	if len(edited) > length {
		edited = edited[0:(length - 1)]
	}
	// All done!
	return edited
}

/*
 * Converts accented characters if found from the dictionary.
 * Otherwise will replace the character with an empty string.
 */
func convertAccent(found string) string {
	if newValue, ok := dictionary[found]; ok {
		return newValue
	}
	return ""
}
