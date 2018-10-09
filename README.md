#### 使用方式(參數token須先手動登入取得)
systemAdmin 登入 取systemToken
gameAdmin 登入 取gameAdminToken
game 登入 取gameToken
paltformAdmin 登入 取platformAdminToken

```
建立gameAdmin帳號 商代:BBGA usernaem : gameAdmin 密碼:54321 group:gameAdmin
```
go run main.go -systemToken=93DEAD123EE


```
建立game帳號 商代:BBGA usernaem : [自訂] 密碼:54321 group:gameTest（設定同gameAdmin）
```
go run main.go -gameAdminToken=93DEAD123EE -groupID=888


```
建立platformAdmin帳號 商代:BBPL usernaem : platformAdmin 密碼:54321 group:platformAdmin
```
go run main.go -gameToken=93DEAD123EE


```
建立platform帳號 商代:BBPL usernaem : [自訂] 密碼:54321 group:platformTest（設定同platformAdmin）
```
go run main.go -platformAdminToken=93DEAD123EE -groupID=999