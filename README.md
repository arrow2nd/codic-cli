# codic-cli

[codic](https://codic.jp/) をターミナルから利用する為の CLI ツール

[![test](https://github.com/arrow2nd/codic-cli/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/codic-cli/actions/workflows/test.yml)
[![release](https://github.com/arrow2nd/codic-cli/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/codic-cli/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/codic-cli)](https://goreportcard.com/report/github.com/arrow2nd/codic-cli)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/codic-cli/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/codic-cli)](https://github.com/arrow2nd/codic-cli/blob/main/LICENSE)

![codic-cli](https://user-images.githubusercontent.com/44780846/146678773-5c518844-f5b9-4ada-a2b4-db3c50a02fc7.gif)

## 必要なもの

[codic](https://codic.jp/)の API キーが必要です。[こちら](https://codic.jp/my/api_status)から取得してください。

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

`get [<日本語>] [--case 名前付け規則] [--prefix 接頭辞スタイル]`

ネーミングを生成します。

```txt
$ codic get こんにちは世界 --case snake --prefix camel
hello_world
```

フラグを省略すると設定したデフォルト値に基づいて生成されます。

また、ネーミングの生成に失敗した場合でもそのまま出力します。

#### 名前付け規則のパラメータ

| パラメータ名 | 規則名                       | 例         |
| ------------ | ---------------------------- | ---------- |
| pascal       | パスカルケース               | PascalCase |
| camel        | キャメルケース               | camelCase  |
| snake        | スネークケース               | snake_case |
| screaming    | スクリーミングスネークケース | SNAKE_CASE |
| kebab        | ケバブケース                 | kebab-case |
| space        | なし（スペース区切り）       | space case |

#### 接頭辞スタイルのパラメータ

| パラメータ名 | 説明                            | 例          |
| ------------ | ------------------------------- | ----------- |
| microsoft    | MS ネーミングガイドラインに準拠 | IOException |
| camel        | CamelCase のルールに準拠        | IoException |
| literal      | リテラル（変換しない)           |             |

### config api

API キーを設定します。

### config case

デフォルトの名前付け規則の種類を設定します。

### config prefix

デフォルトの接頭辞スタイルを設定します。
