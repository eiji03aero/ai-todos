# ai todos
- This is the experimental project to explore the potential of cline.
- Basic policy
    - Never ever write a code yourself. Let the big guy take care of it all

---

# Documents
## How feature development works
- Create feature spec document
    - Update domain model
    - Update Features
- Create api spec
- Generate openapi doc
- Create page spec

## Document details
- docs/Application.md
    - By engineer / AI
- docs/Backend.md
    - By AI
- docs/Codebase.md
    - By AI
- docs/Frontend.md
    - By engineer / AI

---

# Thoughts
- context windowが結構すぐ膨れるので、やり方に工夫がいるかも
    - ある程度の大きさで2-3分作業する単位でタスクをやらせるとか
- コードベースのセットアップは結構苦手かもしれない
    - 資料に書いてあるのと違うことやり始めたり、たぶんコンテキストウィンドウも膨らみがち
    - shadcn/uiのインストールは諦めた
        - https://ui.shadcn.com/docs/installation/tanstack
        - ページにしたがって、コマンドのコピペと指示で対応
    - 設定ファイルの編集はドキュメント見ながら、さすがにコピペした方が早いしお金もかからない
- 設計してもらうのに、設計方針を理解してもらう必要があるが、設計資料をしっかり書くと、理解してもらえた
    - docs/Frontend.mdにFSDのフォルダ構成を追記したら、適切にlayerを活用してくれる件
- architect mode -> code modeが使い勝手良さそう
    - architect modeで壁打ちをしながら実装方針を決める
    - fixした内容をdocs配下に残す
    - そのままcode modeに移行して実装をすすめる
- ちょっと不具合修正が苦手かも
    - docker containerで、npmパッケージが見つからないエラーがでて修正してもらおうとしたときに、あっちこっち行ってさっさと調査に移らなかった
    - これはmodelの良し悪しも関係あるかもしれない
    - ただし、以前出した指示を覚えていて、あとからでも拾ってくれるのは有能
    - そもそもdebugモードを使った方がいいかも
- tailwindのスタイルが当たらない件は直せなかった
    - コードから原因を推測するのはたしかに難しかったと思われる
- 設定ファイルにimportと設定値を追加するのがどうしてもうまくいかない
    - import文追加して終わりになっとる
- 外部ページの情報を読み込ませるときは、@ディレクティブを使う
- 問題の解決の際には、debugモードを使う
- Boomerang mode
    - まるばつゲームの実装をお願いした
    - 仕様検討の壁打ちからやってくれると思ったが、いきなり実装から取り掛かってしまった
        - 内容的に確認するまでもないと判断したのだろうか
    - 既存の機能を削除して、まるばつゲームと入れ替えてしまった。指示が悪かったかも
    - 実装はすぐできた。すげーぞこれ
    - だが、アプリの実装では、検討から丁寧にやるcustom modeが必要そう

---

# Todo
## Backlog
- [ ] Create todo
- [ ] Mark todo done
- [ ] Edit todo
- [ ] Delete todo
- [ ] List and search todo
- [ ] Create project
- [ ] Edit project
- [ ] Delete project
- [ ] List and search Project
- [ ] Project detail
- [ ] update list and search todo
    - search by project
- [ ] Create tag
- [ ] Edit tag
- [ ] Delete tag
- [ ] update todo to attach
- [ ] update list and search todo
    - search by tag

## Done
- [x] create roadmap
    - list out features
- [x] try out boomerang mode
- [x] fix websocket error for local dev server
- [x] play around the architect mode
    - implement helloworlder
- [x] implement sample codebase
    - try to use custom instruction to make it read spec documents
        - .roo/rules
    - [x] api
        - get helloworld
    - [x] client
        - top helloworld
- [x] Implement local environment
    - docker compose
    - docker container
- [x] Create base for Application spec
- [x] Consider how to proceed with the development
- [x] Base of codebase
    - Codebase document
        - top level folders
        - backend (tech stack, folder structure)
        - frontend (tech stack, folder structure)