# Security

## What is Security

- 安全は危険を考えること
- 安全は存在しない
- 安全とは危険を除去・制御すること

## アカウンタビリティとトレーサビリティ

- オープンソース
- プロプライエタリ

## ３大要件

- 秘匿性
    - 知られてはいけない情報
- 完全性
    - 改ざんされない
- 可用性
    - 情報が利用可能
    - 攻撃が容易

[スライド](https://drive.google.com/drive/folders/1xaliw9WpYixeKDj6tmO2rhENJ9dTYW9M)

## OpenSSL
encrypt  
```openssl aes-256-cbc -e -in raw.txt -out enc.txt```  
decrypt  
```openssl aes-256-cbc -d -in enc.txt -out dec.txt```  

## https
- CA証明書
    - Digisertなどの認証局が発行
    - ブラウザ内にルート証明書が埋め込まれている
- 中間認証局
    - ルートとサーバ間の中間的な存在
    - 破棄や再発行がスムーズにできる

# Hardware Security
## Hardware Securityとは
- ハードウェアに直接的に実行される攻撃に対するセキュリティ
    - IoTを想定
    - ソフトウェア的な手段は除外
- IoTシステムは、ハード的な脆弱性も考慮した上でセキュアなシステムを構築する
- HomeGateway
    - セキュリティ調査会社に依頼して評価してもらった

## 書籍
[ハッカーの学校 IoTハッキングの教科書](https://www.amazon.co.jp/ハッカーの学校-IoTハッキングの教科書-黒林檎/dp/4781702368)

## 一般的なITセキュリティとの違い
- ハードウェアを入手、解析される
- 被害
    - WEBカメラの侵入
        - ドンキホーテのWifiカメラ
    - Miraiマルウェア
        - DDoS
        - Anna-senpai
        - github:[jgamblin](https://github.com/jgamblin/Mirai-Source-Code)
    - Jeep Cherokee(実験)
        - HeadUnit(Linuxベース)の内蔵ソフトの脆弱性で権限昇格
        - 車載マイコンのファームウェア書き換え機構を利用してハードウェアへアクセス
    - Tesla(実験)

## ハードウェアハッキングの手順
1. デバイスの入手、分解
    - [ifixit](https://www.ifixit.com)
2. ボードの観察
    - エントリポイントを探す
    - UART, JTAG, etc
3. ポートから侵入、無線経由で侵入
    - [Shikra](https://int3.cc/products/the-shikra)
4. チップの取り外し(チップオフ)
    - Frash ROMを直接読み出し
5. データの吸出し
    - SPIインタフェース
    - メモリをダンプする
6. 吸い出したバイナリの解析
    - [ツール](https://ghidra-sre.org)を使って逆コンパイル

[ドンキの見守りカメラ](http://honeylab.hatenablog.jp/entry/2019/05/03/234417)

# XSS実演
## 攻撃の基本
- 既存の手法が繰り返されている
- 既存の方法をきちんと学ぶことが重要

## 脆弱性
- 掲示板にscriptが実装できる
    - 管理者が意図していないスクリプトが走る

## XSS攻撃
- 攻撃者のスクリプトを脆弱性のあるページで走らせる
- GET/POSTメソッドでも可能