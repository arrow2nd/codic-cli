# codic-cli

[codic](https://codic.jp/) をターミナルから利用する為の CLI ツール

[![test](https://github.com/arrow2nd/codic-cli/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/codic-cli/actions/workflows/test.yml)
[![release](https://github.com/arrow2nd/codic-cli/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/codic-cli/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/codic-cli)](https://goreportcard.com/report/github.com/arrow2nd/codic-cli)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/codic-cli/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/codic-cli)](https://github.com/arrow2nd/codic-cli/blob/main/LICENSE)

![codic-cli](https://user-images.githubusercontent.com/44780846/146678773-5c518844-f5b9-4ada-a2b4-db3c50a02fc7.gif)

## インストール

### Homebrew

```sh
brew tap arrow2nd/tap
brew install codic
```

### それ以外

[Releases](https://github.com/arrow2nd/codic-cli/releases) からお使いの環境にあったファイルをダウンロードしてください。

## コマンド

### get

`get [<日本語>] [--case 名前付け規則] [--prefix 接頭辞のスタイル]`

ネーミングを生成します。

```sh
$ codic こんにちは世界 --case snake --prefix camel
hello_world
```

フラグを省略すると設定したデフォルト値に基づいて生成されます。

また、ネーミングの生成に失敗した場合でもそのまま出力します。

### config

#### api

`config api`

API キーを設定します。

#### case

`config case`

デフォルトの名前付け規則の種類を設定します。

#### prefix

`config prefix`

デフォルトの接頭辞のスタイルを指定します。
