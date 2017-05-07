package reldate

import (
    "testing"
    "time"
    "fmt"
    "strings"
)

func TestConvert(t *testing.T) {

    var result []string

    I18nFunc = func(t string, rest ...interface{}) string {

        var values []string

        for _,v := range rest {
            values = append(values, fmt.Sprintf("%v", v))
        }

        return fmt.Sprintf("%v:%v", t, strings.Join(values, ":"))
    }

    ts := time.Now();

    dat, _ := time.Parse(time.RFC822, "14 Feb 16 15:04 GMT")


    forge := []time.Time{
        ts,                            // just now
        ts.Add(-300 * time.Second),    // a couple of minutes
        ts.Add(-1200 * time.Second),   // n hours ago
        ts.Add(-11000 * time.Second),  // today
        ts.Add(-77000 * time.Second),  // yesterday
        ts.Add(-170000 * time.Second), // day before yesterday
        ts.Add(-350000 * time.Second), // n days ago
        ts.Add(-650000 * time.Second), // n weeks ago
        dat,
    }

    t.Run("JustNow", func(t *testing.T) {

        result = devideResult(Convert(forge[0]))
        if result[0] != "justNow" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("aCoupleOfMinutes", func(t *testing.T) {

        result = devideResult(Convert(forge[1]))
        if result[0] != "aCoupleOfMinutes" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("nHoursAgo", func(t *testing.T) {

        result = devideResult(Convert(forge[2]))
        if result[0] != "nHoursAgo" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("timeToday", func(t *testing.T) {

        result := devideResult(Convert(forge[3]))
        if result[0] != "timeToday" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("timeYesterday", func(t *testing.T) {

        result := devideResult(Convert(forge[4]))
        if result[0] != "timeYesterday" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("timeDayBeforeYesterday", func(t *testing.T) {

        result := devideResult(Convert(forge[5]))
        if result[0] != "timeDayBeforeYesterday" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("nDaysAgo", func(t *testing.T) {

        result := devideResult(Convert(forge[6]))
        if result[0] != "nDaysAgo" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("nWeeksAgo", func(t *testing.T) {

        result := devideResult(Convert(forge[7]))
        if result[0] != "nWeeksAgo" {
            t.Log(result)
            t.Fail()
        }
    })

    t.Run("actualDate", func(t *testing.T) {

        r := Convert(forge[8])
        if r != "14. Feb, 2016" {
            t.Log(r)
            t.Fail()
        }
    })
}

func devideResult(r string) []string {

    return strings.Split(r, ":")
}
