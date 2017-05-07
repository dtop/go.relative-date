package reldate

import (
    "time"
    "math"
)

type I18nConverterFunc func(token string, count ...interface{}) string

var I18nFunc I18nConverterFunc
var DateFormat string
var TimeFormat string
var Timezone string


func defaults() {

    if nil == I18nFunc {
        panic("I18n Converter func is not set, this results in unconverted tokens")
    }

    if "" == DateFormat {
        DateFormat = "02. Jan, 2006"
    }

    if "" == TimeFormat {
        TimeFormat = "15:04"
    }

    if "" == Timezone {
        Timezone = "Europe/Berlin"
    }
}

func Convert(t time.Time) string {

    defaults()

    var relDays, relTime float64

    loc, err := time.LoadLocation(Timezone)
    if err != nil { loc = time.Now().In(time.Local).Location() }

    ts  := t.In(loc)
    now := time.Now().In(loc)

    relTime = float64(now.Unix() - ts.Unix())
    relDays = relTime / 86400.0

    if relTime < 60.0 {
        return I18nFunc("justNow")
    }

    if relTime < 600 {
        return I18nFunc("aCoupleOfMinutes")
    }

    if relTime < 10800 {
        val := int(math.Floor((relTime / 3600) + 0.5))
        return I18nFunc("nHoursAgo", val)
    }

    if relDays < 7 {

        relDayDiff := relativeDayDiff(now, ts)
        switch relDayDiff {

        case 0:
            return I18nFunc("timeToday") + ", " + I18nFunc("localizedTimeUnit", formatTimePart(t, loc))
        case 1:
            return I18nFunc("timeYesterday") + ", " + I18nFunc("localizedTimeUnit", formatTimePart(t, loc))
        case 2:
            return I18nFunc("timeDayBeforeYesterday") + ", " + I18nFunc("localizedTimeUnit", formatTimePart(t, loc))
        default:
            return I18nFunc("nDaysAgo", relDayDiff)
        }
    }

    if relDays < 31 {
        rdays := int(math.Abs(math.Floor(relDays / 7)))
        return I18nFunc("nWeeksAgo", rdays)
    }

    return t.In(loc).Format(DateFormat)
}

func formatTimePart(t time.Time, loc *time.Location) string {

    return t.In(loc).Format(TimeFormat)
}

func relativeDayDiff(now, ts time.Time) int {

    dnow := int(now.Weekday())
    dts  := int(ts.Weekday())

    relDayDiff :=  dnow - dts

    if relDayDiff < 0 {
        return 7 - int(math.Abs(float64(relDayDiff)))
    }

    return relDayDiff
}
