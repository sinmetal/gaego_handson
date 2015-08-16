# gaego_handson

## Initialize

* [GCP Free Trial](https://cloud.google.com/free-trial/?hl=ja)
* [Google Cloud SDK](https://cloud.google.com/sdk/)
* [App Engine Go SDK](https://cloud.google.com/appengine/downloads?hl=ja#Google_App_Engine_SDK_for_Go)

## Tool

* [DHC](https://chrome.google.com/webstore/detail/dhc-resthttp-api-client/aejoelaoggembcahagimdiliamlcdmfm)

## Resouces

* [golang](http://golang.org/)
* [The Go Playground](http://play.golang.org/)
* [A Tour of Go](https://go-tour-jp.appspot.com/#1)
* [Go言語の初心者が見ると幸せになれる場所](http://qiita.com/tenntenn/items/0e33a4959250d1a55045)
* [GOPATH は適当に決めて問題ない](http://qiita.com/yuku_t/items/c7ab1b1519825cc2c06f)
* [GAE/Goのハマったところ( ・᷄ὢ・᷅ )](http://qiita.com/hogedigo/items/fae5b6fe7071becd4051)
* [setup go-lang-idea-plugin for gae/go](http://qiita.com/sinmetal/items/0073a5cf9e613e05786f)
* [goenvでgae/goと普通のgoの環境を切り替える](http://qiita.com/sinmetal/items/71cfba4ae27cc2366572)
* [GAE/GoのWebアプリをCircleCIで自動テスト&自動デプロイする](http://qiita.com/kyokomi/items/84af37e9774faf9072ed)

## Hangs-on

### Part 0

環境整備

[Google Cloud SDK](https://cloud.google.com/sdk/), [App Engine Go SDK](https://cloud.google.com/appengine/downloads?hl=ja#Google_App_Engine_SDK_for_Go)をDownloadし、Pathを通す

goapp commandが使えるようになればOK

```
$ goapp
Go is a tool for managing Go source code.

Usage:

	goapp command [arguments]
...
```

### Part 1

Hello Worldを作成して、deployする

#### local server

```
goapp serve {your src dir}
```

#### deploy command

```
goapp deploy {your src dir}
```

### Part 2

