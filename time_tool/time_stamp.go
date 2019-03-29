/*
参数： t time.Time
返回值：result int64  (时间戳)
*/

package time_tool

import "time"

//如传入time.Now(), 返回今天0点时间的时间戳
func GetFirstDateOfToday(d time.Time) int64 {
	return GetZeroTime(d).Unix()
}

//如传入time.Now(), 返回明天0点时间的时间戳
func GetFirstDateOfTomorrow(d time.Time) int64 {
	return GetZeroTime(d).Unix() + 86400
}

//如传入time.Now(), 返回当前月份的第一天0点时间的时间戳
func GetFirstDateOfMonth(d time.Time) int64 {
	t := GetZeroTime(d.AddDate(0, 0, -d.Day()+1))
	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"), time.Local)
	return stime.Unix()
}

//如传入time.Now(), 返回当前月份的最后一天24点时间的时间戳
func GetLastDateOfMonth(d time.Time) int64 {
	t := GetZeroTime(d.AddDate(0, 0, -d.Day()+1)).AddDate(0, 1, 0)
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:04:05"), time.Local)
	return etime.Unix()
}
