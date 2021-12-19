# codic-cli

[codic](https://codic.jp/) をターミナルから利用する為の CLI ツール

![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/codic-cli/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/codic-cli)](https://github.com/arrow2nd/codic-cli/blob/main/LICENSE)

## コマンド

### get

`get [<日本語>] [--case 名前付け規則] [--prefix 接頭辞のスタイル]`

ネーミングを生成します。

```sh
$ codic こんにちは世界 --case snake --prefix camel
```

### config

#### api

`config api`

API キーを設定します。

#### case

`config case`

名前付け規則の種類を設定します。

#### prefix

`config prefix`

接頭辞のスタイルを指定します。
