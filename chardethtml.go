package chardet

//这个函数出去utf16 utf32的判断,因为实际测试会把gbk识别为utf16be
func Mosthtml(data []byte) string {
	if s := checkbom(data); s != "" {
		return s
	}
	lp := check(data, []detect{&gbk{}, &utf8{}, &big5{}, &eucJP{}, &shiftJIS{}, &iso2022JP{}, &eucKR{}, &gb18030{}, &hzgb2312{}})
	if len(lp) > 0 {
		x, y := -1, -100.0
		for i, l := range lp {
			if r := l.Priority(); y < r {
				x, y = i, r
			}
		}
		return lp[x].String()
	}
	return ""
}
