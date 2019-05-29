package janorm

import "testing"

func check(t *testing.T, exp, from string) {
	t.Helper()
	act := Normalize(from)
	if act != exp {
		t.Fatalf("normalize faield:\nexpect: %s\nactual: %s\n", exp, act)
	}
}

func TestNormalize_zen2han(t *testing.T) {
	check(t, "0123456789", "０１２３４５６７８９")
	check(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ")
	check(t, "abcdefghijklmnopqrstuvwxyz",
		"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ")
	check(t, "! \"#$%&'()*+,-./:;<=>?@[¥]^_`{|}",
		"！　”＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿｀｛｜｝")
}

func TestNormalize_han2zen(t *testing.T) {
	check(t, "。、・「」", "｡､･｢｣")
	check(t, "ハンカク", "ﾊﾝｶｸ")
	check(t, "ゼンカク", "ｾﾞﾝｶｸ")
}

func TestNormalize_bar(t *testing.T) {
	check(t, "o-o", "o₋o")
	check(t, "majikaー", "majika━")
	check(t, "スーパー", "スーパーーーー")
}

func TestNormalize_tilde(t *testing.T) {
	check(t, "わい", "わ〰い")
}

func TestNormalize_spaces(t *testing.T) {
	check(t, "ゼンカクスペース", "ゼンカク　スペース")
	check(t, "おお", "お             お")
	check(t, "おお", "      おお")
	check(t, "おお", "おお      ")
	check(t, "検索エンジン自作入門を買いました!!!", "検索 エンジン 自作 入門 を 買い ました!!!")
	check(t, "アルゴリズムC", "アルゴリズム C")
	check(t, "PRML副読本", "　　　ＰＲＭＬ　　副　読　本　　　")
	check(t, "Coding the Matrix", "Coding the Matrix")
}
