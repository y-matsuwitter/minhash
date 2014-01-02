minhash
=======

b-bit minhash implementation for golang

Requirements
------

* golang

Installation
------

```
go get github.com/y-matsuwitter/minhash
```

Usage
------

```
minhash.Minhash([]string{
    "21", "歳", "ビール", "飲む", "アサヒ", "リクルート", "連携",
    }, []string{
    "21", "歳", "ビール", "無料", "アサヒ", "若者", "向け", "企画",
    })
```

Author
------
Yuki Matsumoto (@y_matsuwitter)
