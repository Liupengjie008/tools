/*
参数： t time.Time
返回值：result time.Time
*/

package time_tool

import "time"

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
