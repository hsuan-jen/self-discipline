## ð ç®å½ç»æ

```
âââ api             (apiå±)
â   âââ v1          (v1çæ¬æ¥å£)
âââ conf            (éç½®æ)
âââ configs         (éç½®å)
âââ core            (æ ¸å¿æä»¶)
âââ docs            (swaggerææ¡£ç®å½)
âââ global          (å¨å±å¯¹è±¡)
âââ initialize      (åå§å)
â   âââ internal    (åå§ååé¨å½æ°)
âââ middleware      (ä¸­é´ä»¶å±)
âââ model           (æ¨¡åå±) 
â   âââ request     (å¥åç»æä½)                
â   âââ response    (åºåç»æä½)
âââ router          (è·¯ç±å±)               
âââ service         (serviceå±)              
âââ utils           (å·¥å·å)
```

## ä½¿ç¨å¦ä¸å½ä»¤ä¸è½½swag
```
go get -u github.com/swaggo/swag/cmd/swag
```

### çæAPIææ¡£
```
swag init
```
## å¼å¯pprof,å¨éç½®æä»¶ä¸­è®¾ç½®
```
system:
  pprof: true
```

### éè¿æµè§å¨è®¿é®
```
http://127.0.0.1:8800/debug/pprof/
```

### ç«ç°å¾
```
# æ§è¡å½ä»¤å,ä¼å¨æµè§å¨æå¼ä¸ä¸ªçªå£
go tool pprof -http=:1234 http://localhost:8800/debug/pprof/goroutine
# ç®åè§£é
# -http è¡¨ç¤ºä½¿ç¨äº¤äºå¼webæ¥å£æ¥çè·åçæ§è½ä¿¡æ¯,æå®å¯ç¨çç«¯å£å³å¯
# debug/pprof/éè¦æ¥ççææ  (allocs,block,goroutine,heap...)
# go tool pprof http://localhost:8800/debug/pprof/goroutine?second=20
# ééåç¨æ°æ®å¹¶æç»­20S
# å¸¸ç¨çå½ä»¤ætop,tree,web,listç­
```
## å¼å¯prometheus,å¨éç½®æä»¶ä¸­è®¾ç½®
```
system:
  promhttp: true
```