package main

import "fmt"

// transform a string composed of 0-9a-zA-Z?! to a int64
func shortLinkToLongInt(shortLink string) int64 {
	shortLinkLength := len(shortLink)
	index := int64(0)

	for i := 0; i < shortLinkLength; i++ {
		ch := int64(shortLink[shortLinkLength-i-1])
		// the ascii of ? id 63, so we will not handle it
		// the ascii of ! is 33, turn it to 62
		if ch == 33 {
			ch = 62
		}
		// ch = 0, 1, 2 ... 9
		if ch >= 48 && ch < 58 {
			ch = ch - '0'
		}
		// ch = A, B, C...Z
		if ch >= 65 && ch < 91 {
			ch = ch - 'A' + 10
		}
		// ch = a,b,c...z
		if ch >= 97 && ch < 123 {
			ch = ch - 'a' + 36
		}
		index += ch << (6 * i)
	}
	return index
}

// transform  an int64 to a string composed of 0-9a-zA-Z?!
func longIntToShortLink(index int64) string {
	var shortLink string
	if index == 0{
		shortLink = "0"
	}
	for index != 0 {
		ch := index - (index>>6)<<6
		if ch >= 0 && ch < 10 {
			shortLink = fmt.Sprintf("%c%s", ch+'0', shortLink)
		}
		if ch >= 10 && ch < 36 {
			shortLink = fmt.Sprintf("%c%s", ch-10+'A', shortLink)
		}
		if ch >= 36 && ch < 62 {
			shortLink = fmt.Sprintf("%c%s", ch-36+'a', shortLink)
		}
		if ch == 62 {
			shortLink = fmt.Sprintf("%c%s", '!', shortLink)
		}
		if ch == 63 {
			shortLink = fmt.Sprintf("%c%s", '?', shortLink)
		}
		index = index >> 6
	}

	for len(shortLink) < 6 {
		shortLink = fmt.Sprintf("%d%s", 0, shortLink)
	}

	return shortLink
}