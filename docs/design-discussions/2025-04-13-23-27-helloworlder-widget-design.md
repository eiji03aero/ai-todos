# HelloWorlder Widgetの設計議論

## 日時
2025年4月13日 23:27

## 背景
/helloworldページに新しいコンポーネントを実装するための設計を行いました。

## 設計方針

### コンポーネント配置
- レイヤー: Widgets Layer
- パス: `client/src/widgets/helloworld/ui/HelloWorlder.tsx`

### 機能要件
1. キャンバス
   - サイズ: 600px x 400px
   - 花火アニメーションの表示
   - ストップウォッチ表示

2. ボタン
   - 「Great」ボタン: 花火アニメーションをトリガー
   - 「Start stopwatch」ボタン: ストップウォッチ開始

### 関連ファイル構造
```
client/src/
├── widgets/
│   └── helloworld/
│       └── ui/
│           └── HelloWorlder.tsx
├── features/
│   └── helloworld/
│       ├── lib/
│       │   ├── fireworks.ts
│       │   └── stopwatch.ts
└── shared/
    ├── ui/
    └── lib/
```

### 技術スタック
- React (Tanstack Start)
- TypeScript
- Tailwind CSS
- shadcn/ui

## 設計の特徴
- Feature Sliced Design (FSD)に準拠
- 自己完結型のWidget
- 再利用可能で拡張性の高い設計

## 合意事項
- Widgetsレイヤーに配置することで合意
- 基本的な花火とストップウォッチ機能を実装

## 次のステップ
1. 実装モードに切り替え
2. 詳細な実装を進める