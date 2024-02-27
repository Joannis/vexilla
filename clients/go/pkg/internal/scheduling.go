package internal

import (
	"time"
)

func IsScheduledFeatureActive(feature Feature) bool {
	return IsScheduleActive(feature.Schedule, feature.ScheduleType)
}

func IsScheduleActive(schedule Schedule, scheduleType ScheduleType) bool {
	return IsScheduleActiveWithNow(schedule, scheduleType, time.Now().UTC())
}

func IsScheduleActiveWithNow(schedule Schedule, scheduleType ScheduleType, now time.Time) bool {

	nowMillis := now.UnixMilli()

	switch scheduleType {
	case EmptyScheduleType:
		return true
	case GlobalScheduleType, EnvironmentScheduleType:

		startDate := time.UnixMilli(schedule.Start).UTC()
		startOfStartDate := time.Date(
			startDate.Year(),
			startDate.Month(),
			startDate.Day(),
			0,
			0,
			0,
			0,
			time.UTC,
		)
		endDate := time.UnixMilli(schedule.End).UTC()
		endOfEndDate := time.Date(
			endDate.Year(),
			endDate.Month(),
			endDate.Day(),
			23,
			59,
			59,
			999_999_999,
			time.UTC,
		)

		if startOfStartDate.UnixMilli() > nowMillis || nowMillis > endOfEndDate.UnixMilli() {
			return false
		}

		startTime := time.UnixMilli(schedule.StartTime).UTC()
		endTime := time.UnixMilli(schedule.EndTime).UTC()

		switch schedule.TimeType {
		case NoneScheduleTimeType:
			return true
		case StartEndScheduleTimeType:

			startOfEndDate := time.Date(
				endDate.Year(),
				endDate.Month(),
				endDate.Day(),
				0,
				0,
				0,
				0,
				time.UTC,
			)

			startDateTimestampWithStartTime := startOfStartDate.UnixMilli() + schedule.StartTime
			endDateTimestampWithEndTime := startOfEndDate.UnixMilli() + schedule.EndTime

			return startDateTimestampWithStartTime <= nowMillis && nowMillis <= endDateTimestampWithEndTime

		case DailyScheduleTimeType:

			zeroDay := time.UnixMilli(0).UTC()
			nowTimestamp := now.UnixMilli()

			todayZeroTimestamp := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				0,
				0,
				0,
				0,
				time.UTC,
			).UnixMilli()

			zeroedStartTimestamp := time.Date(
				zeroDay.Year(),
				zeroDay.Month(),
				zeroDay.Day(),
				startTime.Hour(),
				startTime.Minute(),
				startTime.Second(),
				startTime.Nanosecond(),
				time.UTC,
			).UnixMilli()

			zeroedEndDateTime := time.Date(
				zeroDay.Year(),
				zeroDay.Month(),
				zeroDay.Day(),
				endTime.Hour(),
				endTime.Minute(),
				endTime.Second(),
				endTime.Nanosecond(),
				time.UTC,
			)

			zeroedEndTimestamp := zeroedEndDateTime.UnixMilli()

			startTimestamp := todayZeroTimestamp + zeroedStartTimestamp
			endTimestamp := todayZeroTimestamp + zeroedEndTimestamp

			if zeroedStartTimestamp > zeroedEndTimestamp {
				return nowTimestamp >= startTimestamp || nowTimestamp <= endTimestamp
			} else {
				return nowTimestamp >= startTimestamp && nowTimestamp <= endTimestamp
			}
		default:
			return false
		}
	default:
		return false
	}
}
