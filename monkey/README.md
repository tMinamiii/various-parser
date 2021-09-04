# Monkey Parser

1. lexerがトークンを探し
    - トークンの種別をType
    - トークンの文字情報をリテラル
2. parserがトークンの種別を判定し、
3. astがトークンの連続からStatement(命令)を判別する
