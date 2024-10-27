package sequences

import "strconv"

func lookSay(str string) (rstr string) {
	cbyte := str[0]
	inc := 1
	for i := 1; i < len(str); i++ {
		dbyte := str[i]
		if dbyte == cbyte {
			inc++
			continue
		}
		rstr = rstr + strconv.Itoa(inc) + string(cbyte)
		cbyte = dbyte
		inc = 1
	}
	return rstr + strconv.Itoa(inc) + string(cbyte)
}
