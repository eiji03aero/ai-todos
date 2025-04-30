# Setting up openapi-fetch with react-query
```
can you do following?

- research and introduce openapi-fetch and openapi-react-query while following instructions at @https://openapi-ts.dev/openapi-react-query/  
- please define a npm script that do the generation work for this package
```

補足:
- これくらいなら対話しつつうまく行っていそう

# Signup feature prompts
```
I want to create a feature design for SignUp

1. User opens signup page
2. User fills signup form
3. User submits the signup form
4. If signup was a success, user will be redirected to top page
```

# Implement sample feature on helloworld
```
/helloworldページに以下の機能を実装するための設計をしてください
おもにどのフォルダにどういうコードを作成するかを考えてください

機能の内容:
- コンポーネント名: HelloWorlder
- 仕様:
    - canvasを持ち操作に応じたアニメーションを表示する
    - canvasのサイズは600px x 400px
    - Greatボタンをもつ。このボタンを押すと、canvasに花火を表示する
        - 花火のアニメーションは基本的な線形と円形を駆使したシンプルなもので良い
    - Start stopwatchボタンを持つ。押したら押された瞬間から何分何秒たったかをcanvasに描画する
        - ストップウォッチの表示フォーマットは分:秒のフォーマットで毎秒描画して
```

補足:
- どうしてもpages配下にコードを置きがちになる。ここは流石に難しいので、キャッチボールをして方針修正をする

# Refine folder structure
```
client/appフォルダーをsrcにrenameしたいです
以下にしたがって修正してみて
- client/app.config.tsのdefineConfigの引数に以下を追加する
  tsr: {
    appDirectory: 'src',
  },
- client/appフォルダーをsrcにrenameする
```

# Add Dockerfile for client container
```
Dockerfileを以下条件で作成して
- client/Dockerfileに作成
- docker-compose.ymlのclientサービス向けのもの
- nodejs v22.14-bookwormのdocker imageを使う
- commandでnpm run devを実行
```

# Impelment hellow world on client part 2
```
client/app/routesフォルダーをclient/app配下に移動してください
パスの変更に必要なコードの修正もお願い
```

# Impelment hellow world on client
```
clientに以下の実装を加えてください

- clientにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- サンプル実装は以下の内容のページを作成
    - ページ: /helloworld
    - h1でhello worldと表示するだけ
```

補足:
- 設計資料にtanstack startを使ってと書いているのに、nextでやり始めてしまった。やり直し

```
clientに以下の実装を加えてください

- clientにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- 構成や技術スタックはdocs/Frontend.mdを参考にする
- サンプル実装は以下の内容のページを作成
    - ページ: /helloworld
    - h1でhello worldと表示するだけ
```

補足:
- パッケージマネージャの選択からやりはじめた。これは資料に書いておくべき
- 関連ライブラリのインストールをし始めた
- なんかtanstack startの構成ではなかったっぽいのでやりなおし

```
clientに以下の実装を加えてください

- clientにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- 構成や技術スタックはdocs/Frontend.mdを参考にする
    - こちらを参考にtanstack startでコードベースを初期化する https://tanstack.com/start/latest/docs/framework/react/build-from-scratch
- サンプル実装は以下の内容のページを作成
    - ページ: /helloworld
    - h1でhello worldと表示するだけ
```

補足:
- なんか知らんが、tanstack startではなく、viteをねじ込んでこようとする。学習量の差とかなのかも
- パッケージ指定がうまくいかないので、やりなおし

```
こちらを参考にtanstack startでコードベースを初期化する https://tanstack.com/start/latest/docs/framework/react/build-from-scratch
clientディレクトリにコードベースを作る
```

```
docs/Frontend.mdを以下指示に沿って修正してください

- フォルダ構成の部分を、現在のclient配下のディレクトリ構成に沿ったまま、Feature sliced designの構成が組み込めるようにする
    - Feature sliced designの構成はclient/app配下に作るようにする
```

```
clientに以下の実装を加えてください

- 以下の内容でサンプル実装を作成する
    - ページパス: /helloworld
    - h1でhello worldと表示するだけ
```

```
clientに以下の実装を加えてください

- clientにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- 構成や技術スタックはdocs/Frontend.mdを参考にする
- サンプル実装は以下の内容のページを作成
    - ページ: /helloworld
    - h1でhello worldと表示するだけ
```

補足:
- 結局tanstack startの初期化で、シェルでインタラクティブに作成する部分は人間の手でやってしまった

# Impelment hellow world on api
```
apiに以下の実装を加えてください

- apiにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- 構成や技術スタックはdocs/Backend.mdを参考にする
- サンプル実装は以下の内容のapiを一本作成
    - get: /helloworld
    - レスポンスにはhelloworoldと返すだけ
```

補足:
- go mod initをコンテナ内でやってほしく、その分のやり取りが発生した
- go mod initがコンテナが即時終了する (hello worldをechoするだけ終了するため、の件をログのアウトプットを確認しがら修正してくれたすごいぞ)
- docker compose ymlを修正するときに勝手に他の関係ないサービスも修正していた
- なぜかファイルの作成と内容の作成をコンテナ内のcatコマンドでやろうとしていた
- helloworldの実装はかなりステップバイステップ
- 結局コンテキストウィンドウが膨らみすぎたため、別のタスクで仕切り直すことに

```
apiに以下の実装を加えてください

- apiにサンプル実装を作成しながらフォルダ / ファイルの構成のベースを作成する
- 構成や技術スタックはdocs/Backend.mdを参考にする
- サンプル実装は以下の内容のapiを一本作成
    - get: /helloworld
    - レスポンスにはhelloworoldと返すだけ
```

補足:
- また制限にひっかかったため、再度やり直し

# Implement docker compose
```
以下の指示を守りながらローカル開発用のdocker-composeを実装してください

- docker-compose.ymlをルートディレクトリに作成する
- db用のコンテナを以下内容で作成
    - サービス名はdb
    - postgres v17.4-bookwormのdocker imageを使う
    - デフォルトdbはai-todo
    - デフォルトユーザーはai-todo
    - デフォルトパスワードはpass
- api用のコンテナを以下内容で作成
    - サービス名はapi
    - golang v1.24のdocker imageを使う
    - コンテナ内の/appディレクトリをworking_dirに指定する
    - apiディレクトリをコンテナ内の/appディレクトリにマウントする
    - 4200 portを開ける
    - commandはhello worldをechoするだけ
- フロントエンド用のコンテナを以下内容で作成
    - サービス名はclient
    - nodejs v22.14のdocker imageを使う
    - コンテナ内の/appディレクトリをworking_dirに指定する
    - clientディレクトリをコンテナ内の/appディレクトリにマウントする
    - 3000 portを開ける
    - commandはhello worldをechoするだけ
```

# Signup feature spec
```
以下の指示を守りながらSignup機能の仕様を追記してください

- docs/Application.mdのFeatures sectionに追記する
- 内容は以下
  - 概要
    - ユーザーがai-todoにユーザー登録するための機能を提供する
  - フロー
    - ユーザーはサインアップページを開く
    - ユーザーはEメールとパスワードを入力する
    - ユーザーはフォームを送信する
    - 登録が成功すると、アプリのトップページにリダイレクトされる
```

# Create frontend document
```
以下の指示を守りながらフロントエンド開発の方針となる資料を作成してください

- docs/Frontend.mdというファイルを作成し、その中に資料を作成する
- 技術スタックについて以下の組み合わせで記載する
    - 言語: Typescript
    - UIフレームワーク: React
    - Reactフレームワーク: tanstack start
    - UIライブラリ: shadcn/ui
    - スタイリング: tailwind, css modules
    - フォーム: tanstack form
    - http通信: axios
    - http通信管理: tanstack query
- フォルダ構成は以下で
    - Feature sliced design
```

# Create backend document
```
以下の指示を守りながらバックエンド開発の方針となる資料を作成してください

- docs/Backend.mdというファイルを作成し、その中に資料を作成する
- 技術スタックについて以下の組み合わせで記載する
    - 言語: Golang
    - httpサーバーフレームワーク: Gin
    - ORM: Gorm
- フォルダ構成は以下で
    - オニオンアーキテクチャ
```