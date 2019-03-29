日常开发常用工具封装总结：



mobile_tool 目录：

手机号真伪判断：
func MobileVerification(mobile string) bool 
手机号中间中间四位标*：
func MobileNote(mobile string) string


time_tool 目录：
返回今天0点时间的时间戳：
func GetFirstDateOfToday(d time.Time) int64
返回明天0点时间的时间戳：
func GetFirstDateOfTomorrow(d time.Time) int64
返回当前月份的第一天0点时间的时间戳：
func GetFirstDateOfMonth(d time.Time) int64 
返回当前月份的最后一天24点时间的时间戳：
func GetLastDateOfMonth(d time.Time) int64



