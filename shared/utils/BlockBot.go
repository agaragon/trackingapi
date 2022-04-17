package utils

import (
    "github.com/mileusna/useragent"
    . "trackingApp/shared/customError"
)



func BlockBot(userAgent string) error {

        ua := ua.Parse(userAgent)
        if (!ua.Mobile && !ua.Tablet && !ua.Desktop) || ua.Bot {
            return &AccessDenied{"Bot detectado"}
        }
        return nil
}