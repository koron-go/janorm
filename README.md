# Japanese Charaters Normalization

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/janorm)](https://pkg.go.dev/github.com/koron-go/janorm)
[![Actions/Go](https://github.com/koron-go/janorm/workflows/Go/badge.svg)](https://github.com/koron-go/janorm/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/janorm)](https://goreportcard.com/report/github.com/koron-go/janorm)

日本語の文字(キャラクタ)を検索等に適した形に正規化します。

日本語のキャラクタセットには1つの文字にも拘わらず、
描画幅の違いで異なるコードポイントを割り当てられたものが存在します。
数字(`123` or `１２３`)やカタカナ(`イロハ` or `ｲﾛﾊ`)等がその代表です。
いわゆる半角・全角と言われるものです。

またほぼ同じ字形にも拘わらず複数のコードポイントを割り当てられた記号も存在します。

`janorm` パッケージはこのような日本語文字の多義性をいずれかに変換・統一または削除することで正規化し、
検索等の機械処理に適した形に変換します。

大まかな正規化(変換)ルールは以下の通りです。

文字種 | 正規化方法 | 正規化の例
-------|------------|----------
数字   |半角        |`012345` ← `０１２３４５`
アルファベット|半角 |`ABCxyz` ← `ＡＢＣｘｙｚ`
ASCII記号|半角      |` !"#$%` ← `　！”＃＄％`
句点,読点,中点,カッコ,調音記号|全角|`。、・「」ー` ← `｡､･｢｣ｰ`
カタカナ|全角       |`アイウエオ` ← `ｱｲｳｴｵ`
半カタ+濁点・半濁点|全角|`ヴガギグ` ← `ｳﾞｶﾞｷﾞｸﾞ`
ハイフンマイナス記号|統一|`-`
全角長音記号|統一|`ー`
チルダ状記号|削除|(n/a)
