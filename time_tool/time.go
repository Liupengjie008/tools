/*
参数： t time.Time
返回值：result time.Time or time.Duration
*/

package time_tool

import "time"

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取今天剩余秒数
func GetRemainingTime() (expiration time.Duration) {
	now := time.Now()
	lastTime, _ := time.ParseInLocation("2006-01-02", now.Format("2006-01-02"), time.Local)
	return time.Duration(lastTime.AddDate(0, 0, 1).Unix()-now.Unix()) * time.Second
}
