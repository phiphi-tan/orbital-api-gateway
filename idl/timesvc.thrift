namespace go timesvc

struct TimeReq {
    1: required bool twentyfourhour,
}

struct TimeResp {
    1: string time,
}

struct DateReq {
     1: required bool american_formatting,
}

struct DateResp {
    1: string date,
}

service TimeSvc {
    TimeResp getTime(1: TimeReq req)( api.post = '/timesvc/time', api.param = 'true')

    DateResp getDate(1: DateReq req)( api.post = '/timesvc/date', api.param = 'true')
}
