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