## ysrenamer
シンプルなリネームツールです。

## 概要
あらかじめ設定ファイルに記述しておいた**リネームコマンド**をファイルに適用させることができます。

例えば、

```json
{
  "ltrim": {
    "action": "regexp_replace",
    "args": ["^\\s+", "", "true"]
  },

  "rtrim": {
    "action": "regexp_replace",
    "args": ["\\s+$", "", "true"]
  },

  "trim": {
    "action": "order",
    "args": ["ltrim", "rtrim"]
  }
}
```

このようにJSON形式で設定ファイルを書いておくと、

```sh
$ ysrenamer rtrim "foo/bar .hoge"
foo/bar .hoge -> foo/bar.hoge
```

このようにコマンドに従ってファイルリネームを実行します。

## インストール
### go get
```sh
go get github.com/yuukichi/ysrenamer/cmd/ysrenamer
```

## 設定ファイルの位置

### Windows
```
%APPDATA%\ysrenamer\ysrenamer.json
```

### Others
```
~/.ysrenamer.json
```

## action
### replace
`args[0]`に一致した部分を`args[1]`で置換します。`args[2]`回置換を実行します。  
`args[2]`の初期値は1で、負の値を入力すると全て置換します。  
`args[3]`に`true`を指定すると拡張子を除外してリネームを実行します。

### regexp_replace
`args[0]`で指定した正規表現に一致した部分を`args[1]`で**全て**置換します。  
`args[2]`に`true`を指定すると拡張子を除外してリネームを実行します。

### order
設定ファイルに記述されている他コマンドを`args`に記述した順に実行します。
