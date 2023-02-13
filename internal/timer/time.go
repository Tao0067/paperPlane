package timer

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) // 解析持续时间 其支持的有效单位有"ns”, “us” (or “µ s”), “ms”, “s”, “m”, “h”，例如：“300ms”, “-1.5h” or “2h45m”
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), nil // 得到之后的时间
}
