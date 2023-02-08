## 
```bash
# RSA 2048ビットの秘密鍵を作成
$ openssl genrsa -out ca.key 2048

# 証明書署名要求(CSR)を作成
$ openssl req -new -sha256 -key ca.key -out ca.csr -config openssl.cnf

# 証明書を自分の秘密鍵で署名して作成
$ openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA
```

```bash
# RSA 2048ビットの秘密鍵を作成
$ openssl genrsa -out server.key 2048

# 証明書署名要求(CSR)を作成
$ openssl req -new -nodes -sha256 -key server.key -out server.csr -config openssl.cnf

# 証明書を自分の秘密鍵で署名して作成
$ openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server
```

```bash
# RSA 2048ビットの秘密鍵を作成
$ openssl genrsa -out client.key 2048

# 証明書署名要求(CSR)を作成
$ openssl req -new -nodes -sha256 -key client.key -out client.csr -config openssl.cnf

# 証明書を自分の秘密鍵で署名して作成
$ openssl x509 -req -days 365 -in client.csr -sha256 -out client.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Client
```

## Tips 
### certificate signed by unknown authority エラー
Dockerfileに以下を追記してオレオレ認証局を信用させる必要がある           
```
RUN apk --no-cache add ca-certificates
ADD ca.crt /etc/ssl/certs/      
```
