package common

import "regexp"

//var re = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
var RePhone = regexp.MustCompile(`[84|]+(3|5|7|8|9|1[2|6|8|9])+([0-9]{8})`)
