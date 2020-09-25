# lis.go

GoによるミニマムなScheme処理系です。

PythonによるミニマムなScheme処理系である[lis.py](http://www.aoky.net/articles/peter_norvig/lispy.htm)を元にしています。

## 実行方法

### ビルドせずに実行

```
$ make run
```

### ビルド
`./lisgo` という実行可能バイナリが生成されます。

```
$ make build
```

## 主なファイルと説明
* `lis.go` エントリポイントとreplのループ
* `tokenize.go` 字句・構文解析器
* `eval.go` 評価器
* `token.go` ASTを表す構造体
* `env.go` 「環境」を表す構造体
* `operators.go` 組み込み関数

## 実装されている(されていない)機能
[todo.md](/todo.md)を参照してください。

## LICENSE

MIT LICENSE
