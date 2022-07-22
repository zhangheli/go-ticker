go-ticker是高性能定时器库
## feature
* 支持周期性定时器

## 一次性定时器
```go
import (
    "github.com/antlabs/timer"
    "log"
)

func main() {
        tm := timer.NewTimer()

        tm.AfterFunc(1*time.Second, func() {
                log.Printf("after\n")
        })

        tm.AfterFunc(10*time.Second, func() {
                log.Printf("after\n")
        })
        tm.Run()
}
```

## 周期性定时器
```go
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangheli/go-ticker/ticker"
)

func TestTicker(t *testing.T) {
	action := func() {
		fmt.Println(time.Now())
	}
	stopC := make(chan string, 1)

	st := ticker.SecondTicker{}
	st.SetPoint([]int{1, 100, 200, 300, 400, 500})
	st.Run(stopC, action)

	<-make(chan int, 1)
}

```

运行结果
```
2022-07-22 12:44:54.001083 +0800 CST m=+4.560425084
2022-07-22 12:44:54.100082 +0800 CST m=+4.659426001
2022-07-22 12:44:54.200079 +0800 CST m=+4.759425417
2022-07-22 12:44:54.300076 +0800 CST m=+4.859424334
2022-07-22 12:44:54.400075 +0800 CST m=+4.959425834
2022-07-22 12:44:54.500073 +0800 CST m=+5.059426042
```
平均 0.07 ms 误差


## 爆破型定时器，加特林突击
```go
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhangheli/go-ticker/ticker"
)

func TestGatlin(t *testing.T) {
	action := func() {
		fmt.Println(time.Now())
	}
	now := time.Now().UnixMilli()
	p := now + (1000 - now%1000)
	runBegin := time.UnixMilli(p)

	ticker.GatlinWithLimit(runBegin, 10, action)
}
```
运行效果
```
2022-07-22 12:50:14.000916 +0800 CST m=+0.511891751
2022-07-22 12:50:14.005861 +0800 CST m=+0.516837167
2022-07-22 12:50:14.010862 +0800 CST m=+0.521838209
2022-07-22 12:50:14.015867 +0800 CST m=+0.526842751
2022-07-22 12:50:14.020858 +0800 CST m=+0.531834667
2022-07-22 12:50:14.025865 +0800 CST m=+0.536841709
2022-07-22 12:50:14.030859 +0800 CST m=+0.541835459
2022-07-22 12:50:14.035908 +0800 CST m=+0.546884584
2022-07-22 12:50:14.040165 +0800 CST m=+0.551141667
2022-07-22 12:50:14.045885 +0800 CST m=+0.556862126
2022-07-22 12:50:14.050858 +0800 CST m=+0.561834751
```
误差 0.1 ~ 0.8ms 左右