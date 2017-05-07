# go.relative-date
A relative date formed from time.Time

### Prequisites

Since this lib has no built in functionality for i18n I would recomment using e.g. [nicksnyder/i18n](https://github.com/nicksnyder/go-i18n)

### Installation

```go

go get github.com/dtop/go.relative-date

```

### Usage

```go

// the following is an example for i18n with the previously mentioned lib

import (
    "github.com/nicksnyder/go-i18n/i18n"
)

i18n.MustLoadTranslationFile("../goi18n/testdata/expected/en-us.all.json")
T, _ := i18n.Tfunc("en-US")


// declare your I18n func, here you would convert tokens and vars
// to a human understandable string 

reldate.I18nFunc = func(token string, vars ...interface{}) {

    if len(vars) {
       
       // vars can be numbers if numbers are part of 
       // the expression
    }   
    
    return T("token", Count{vars[0]})
}

reldate.TimeFormat = "15:04"
reldate.DateFormat = "02. Jan, 2006"
reldate.Timezone   = "Europe/Berlin"

ts := time.Now()

expression := reldate.Convert(ts)

```
