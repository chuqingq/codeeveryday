# HTTPS双向认证

## 生成证书

### CA

```shell
# 私钥
openssl genrsa -out ca.key 2048
# 公钥
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt
```

### 服务端

```shell
# 私钥
openssl genrsa -out server.pem 1024
openssl rsa -in server.pem -out server.key
# 签发请求
openssl req -new -key server.pem -out server.csr
# 用CA签发
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -out server.crt
```

### 客户端

```shell
# 私钥
openssl genrsa -out client.pem 1024
openssl rsa -in client.pem -out client.key
# 签发请求
openssl req -new -key client.pem -out client.csr
# 用CA签发
openssl x509 -req -sha256 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -out client.crt
```

## nginx配置

```config
    server{
        listen       10443 ssl; # https listen port
        server_name  117.78.43.39;

        ssl_certificate      /root/cqq/server.crt;
        ssl_certificate_key  /root/cqq/server.key;

        ssl_session_cache    shared:SSL:1m;
        ssl_session_timeout  5m;

        ssl_ciphers  HIGH:!aNULL:!MD5;
        ssl_prefer_server_ciphers  on;

        ssl_client_certificate /root/cqq/client.crt;
        ssl_verify_client on;

        location /push {
            echo "hello world";
        }
    }
```

详细参见nginx.conf。 

## 验证

### curl验证

```shell
# 失败：400 No required SSL certificate was sent
curl -k  ./client.crt https://testlfs.powerapp.io:10443/push
# 失败：400 The SSL certificate error
curl -k --key ./server.key --cert ./server.crt https://testlfs.powerapp.io:10443/push
# 成功
curl -k --key ./client.key --cert ./client.crt https://testlfs.powerapp.io:10443/push
```

### java验证（OkHttp）

java端的验证，主要就是keymanager用的keystore和trustmanager用的keystore，分别是客户端的JKS和服务端的证书（也是JKS格式）。

```shell
# 生成客户端jks
openssl pkcs12 -export -in client.crt -inkey client.pem -out client.p12 -name client -passin pass:123456 -passout pass:123456
keytool -importkeystore -srckeystore client.p12 -srcstoretype PKCS12 -srcstorepass 123456 -alias client -deststorepass 123456 -destkeypass 123456 -destkeystore client.jks
# 生成服务端jks
keytool -import -rfc -file server.crt -keystore server.jks
```

keymanager: client.jks
trustmanager: server.jks

场景2：

如果是客户端给JSK，例如hitouch_clients_pub.jks，alias是hitouch_clients_pub，密码是123456，则需要：

导出证书：
```shell
keytool -export -alias hitouch_clients_pub -keystore hitouch_clients_pub.jks -rfc -file hitouch_clients_pub.crt
```

然后在nginx中使用这个证书作为ssl_client_certificate的证书。
java代码中keymanager使用这个hitouch_clients_pub.jks。

```java
        // 加载客户端证书
        SSLSocketFactory sslSocketFactory = null;
        try {
            // keystore: client.keystore
            KeyStore keyStore = KeyStore.getInstance(KeyStore.getDefaultType());
            InputStream clientCrt = new FileInputStream("D:\\work\\01.https双向\\hitouch_clients_pri.jks");
            keyStore.load(clientCrt, "975832".toCharArray());
            KeyManagerFactory keyManagerFactory = KeyManagerFactory.getInstance("SunX509");// 安卓可能需要把SunX509换成X509
            keyManagerFactory.init(keyStore, "975832".toCharArray());
            clientCrt.close();

            // truststore: ca.keystore
            System.out.println("KeyStore.getDefaultType(): " + KeyStore.getDefaultType());
            KeyStore trustStore = KeyStore.getInstance(KeyStore.getDefaultType());
            InputStream caCrt = new FileInputStream("D:\\work\\01.https双向\\server2.jks");
            trustStore.load(caCrt, "123456".toCharArray());
            TrustManagerFactory trustManagerFactory = TrustManagerFactory.getInstance("SunX509");
            trustManagerFactory.init(trustStore);
            caCrt.close();

            // sslSocketFactory
            SSLContext sslContext = SSLContext.getInstance("TLS");
            sslContext.init(keyManagerFactory.getKeyManagers(), trustManagerFactory.getTrustManagers(), null);
            sslSocketFactory = sslContext.getSocketFactory();
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }

        // 请求
        OkHttpClient httpClient = new OkHttpClient.Builder().sslSocketFactory(sslSocketFactory)
        // .hostnameVerifier(new TestOkHttpHttps().new UnSafeHostnameVerifier())
                .build();
        try {
            Request request = new Request.Builder().url("https://testlfs.powerapp.io:9879/push").get().build();
            long start = System.currentTimeMillis();
            Response response = httpClient.newCall(request).execute();
            long stop = System.currentTimeMillis();
            System.out.println("time: " + (stop - start));

            // 失败
            if (response == null || response.code() != 200) {
                System.out.println("error: " + response.code() + " " + response.message() + response.body().string());
                return;
            }

            System.out.println("success: " + response.body().string());
        } catch (Exception e) {
            System.out.println("request exception: " + e);
            return;
        }
```

### 浏览器导入密钥

浏览器默认无法访问双向认证的服务，需要导入xxx

keytool -importkeystore -srcstoretype JKS -srckeystore hitouch_clients_pri.jks -srcalias hitouch_clients_pri -deststoretype PKCS12 -destkeystore hitouch_clients_pri.p12  -destalias hitouch_clients_pri

## 参考

[NGINX 配置 SSL 双向认证](http://www.cnblogs.com/UnGeek/p/6049004.html "NGINX 配置 SSL 双向认证")
[Java实现 SSL双向认证](http://www.cnblogs.com/yqskj/p/3142861.html "Java实现 SSL双向认证")
[keytool和openssl生成的证书转换](http://www.cnblogs.com/cuimiemie/p/6442668.html "keytool和openssl生成的证书转换")
