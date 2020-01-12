package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Printf("%s\n", nl2br(`This\r\nis\n\ra\nstring\r`))
}

func addslashes(s string, charlist []rune) string {
	if charlist == nil {
		charlist = []rune(`'"\`)

	}
	r := string(charlist)

	//fmt.Printf("%q %#[1]v\n", charlist)

	var buf bytes.Buffer

	rs := []rune(s)
	for _, v := range rs {
		if strings.ContainsRune(r, v) {
			buf.WriteString(`\`)
		}
		buf.WriteString(string(v))
	}

	return buf.String()
}

/*
afsdbcvc\'fff\\\"222
afsdbcvc'fff\"222

 */
func stripslashes(s string) string {
	return strings.ReplaceAll(s, `\`, "")
}

var charMap = map[rune]string{
	'&': "&amp;",
	'"': "&quot;",
	'\'': "&#039;",
	'<': "&lt;",
	'>': "&gt;",
}

func htmlspecialchars(s string) string {
	var buf bytes.Buffer
	for _, v := range s {
		if convert, ok := charMap[v]; ok {
			buf.WriteString(convert)
		} else {
			buf.WriteString(string(v))
		}
	}
	return buf.String()
}

func htmlspecialcharsDecode(encode string) string {
	oldnew := make([]string, 0)
	for k, v := range charMap {
		oldnew = append(oldnew, v, string(k))
	}
	return strings.NewReplacer(oldnew...).Replace(encode)
}

func chunkSplit(body string, chunklen int, end string) string {
	if chunklen == 0 {
		chunklen = 3
	}
	if end == "" {
		end = "\r\n"
	}

	bodyRunes := []rune(body)
	max := len(bodyRunes)
	//fmt.Println("max=", max)
	start := 0

	var buf bytes.Buffer
	for {
		//fmt.Printf("start:%d, chunklen:%d\n", start, chunklen)

		to := int(math.Min(float64(max), float64(start + chunklen)))
		//to := start + chunklen
		chunk := bodyRunes[start:to]
		buf.WriteString(string(chunk))
		buf.WriteString(end)

		start += chunklen
		if start >= max {
			break
		}
	}
	return buf.String()
}

func countChars(s string) (r map[rune]int) {
	r = make(map[rune]int)
	for _, v := range s {
		r[v]++
	}
	return r
}

func join(gule string, pieces []string) string {
	return strings.Join(pieces, gule)
}

func explode(delimiter string, str string) []string {
	return strings.Split(str, delimiter)
}

func nl2br(s string) string {
	rns := []string{`\r\n`, `\r`, `\n`}
	oldnew := make([]string, 0)
	for _, v := range rns {
		oldnew = append(oldnew, v, `<br>`+v)
	}

	return strings.NewReplacer(oldnew...).Replace(s)
}