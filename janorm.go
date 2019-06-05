package janorm

import "golang.org/x/text/width"

var defaultNormalizer *normalizer

func init() {
	b := newBuilder(1000)

	// map Zen-kaku chars to Han-kaku chars.
	b.putEach(
		"０１２３４５６７８９", 1,
		"0123456789", 1)
	b.putEach(
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ", 1,
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1)
	b.putEach(
		"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ", 1,
		"abcdefghijklmnopqrstuvwxyz", 1)
	b.putEach(
		"　！”＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿｀｛｜｝", 1,
		" !\"#$%&'()*+,-./:;<=>?@[¥]^_`{|}", 1)

	// map Han-kaku chars to Zen-kaku chars
	b.putEach(
		"｡､･｢｣ｰ", 1,
		"。、・「」ー", 1)
	b.putEach(
		"ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾒﾐﾑﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ", 1,
		"アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマメミムモヤユヨラリルレロワヲン", 1)
	b.putEach(
		"ｧｨｩｪｫｬｭｮｯ", 1,
		"ァィゥェォャュョッ", 1)
	b.putEach(
		"ｳﾞｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ", 2,
		"ヴガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ", 1)

	// hyphen/minus
	b.putMap("-",
		"\u02d7", "\u058a", "\u2010", "\u2011", "\u2012", "\u2013",
		"\u2043", "\u207b", "\u208b", "\u2212")

	// Zen-kaku 長音
	b.putMap("ー",
		"\u2014", "\u2015", "\u2500", "\u2501", "\ufe63", "\uff70")

	// remove tilde like characters
	b.putMap("", "~", "∼", "∾", "〜", "〰", "～")

	defaultNormalizer = b.normalizer()
}

// Normalize normalizes a string as Japanese text.
func Normalize(s string) string {
	s = defaultNormalizer.normalize(s)
	s = cleanup(s)
	return s
}

func cleanup(s string) string {
	var (
		b           = make([]rune, 0)
		lastSpace   = true
		lastZenbar  = false
		lastZenkaku = false
	)
	for _, r := range s {
		switch r {
		case ' ':
			if lastSpace || lastZenkaku {
				continue
			}
			break
		case 'ー':
			if lastZenbar {
				continue
			}
			break
		default:
			if lastSpace && isZenkaku(r) && len(b) > 0 {
				b = b[:len(b)-1]
				lastSpace = false
			}
			break
		}
		b = append(b, r)
		lastSpace = r == ' '
		lastZenbar = r == 'ー'
		lastZenkaku = isZenkaku(r)
	}
	if lastSpace && len(b) > 0 {
		b = b[:len(b)-1]
	}
	return string(b)
}

func isZenkaku(r rune) bool {
	k := width.LookupRune(r).Kind()
	switch k {
	case width.EastAsianAmbiguous,
		width.EastAsianWide,
		width.EastAsianFullwidth:
		return true
	default:
		return false
	}
}
