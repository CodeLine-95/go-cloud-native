package logz

import (
	"sync"
	"time"
)

var fmtTime string
var fmtTimeMtx = new(sync.RWMutex)

// FmtTime 返回格式化时间
func FmtTime() string {
	fmtTimeMtx.RLock()
	defer fmtTimeMtx.RUnlock()

	return fmtTime
}

// SyncTime 同步时间
func SyncTime() {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			fmtTimeMtx.Lock()
			fmtTime = time.Now().Format(time.DateTime)
			fmtTimeMtx.Unlock()
		}
	}
}
