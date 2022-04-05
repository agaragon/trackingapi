package utils

import (
    "github.com/mileusna/useragent"
    "net/http"
    . "trackingApp/shared/customError"
)



func BlockBot(req *http.Request) error {

        ua := ua.Parse(req.Header.Get("user-agent"))
        if (!ua.Mobile && !ua.Tablet && !ua.Desktop) || ua.Bot {
            return &AccessDenied{"Bot detectado"}
        }
        return nil
}