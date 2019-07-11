# 简介
这是一个将以太坊keystore+passphrase转成私钥的小工具。

核心方法来自以太坊（go-ethereum）。这里只是调用了相关方法。

# 使用方式
## 编译

在当前目录下生成keystore_privatekey可执行文件
```
go build keystore_privatekey.go
```

## 输出私钥到文件

```
./keystore_privatekey [keystore文件位置]
passphrase: 
have write to file: [地址.txt]
```

## 输出私钥到终端

```
./keystore_privatekey --toconsole [keystore文件位置]
passphrase: 
[输出私钥]
```

# 其他
## 私钥如何转keystore
使用geth即可

```
geth account import [包含私钥的文件位置]
```

所以可以用本工具将新生成的账户keystore转成私钥，
移除keystore后（重要账户请备份！），
再使用 `geth account import` 导入geth成keystore文件来测试有无问题。
 
## 测试
不要问目前为啥没test;)




