package utils

import (
	"fmt"
	"html/template"
	"time"
)

func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatDate": func(dateStr string) string {
			t, err := time.Parse("2006-01-02T15:04", dateStr)
			if err != nil {
				t, err = time.Parse("2006-01-02", dateStr)
				if err != nil {
					return dateStr
				}
			}
			return t.Format("02/01/2006")
		},
		"formatDateFlight": func(dateStr string) string {
			t, err := time.Parse("2006-01-02T15:04", dateStr)
			if err != nil {
				t, err = time.Parse("2006-01-02", dateStr)
				if err != nil {
					return dateStr
				}
			}
			return t.Format("Mon 02 Jan '06")
		},
		"calculateNights": func(checkin, checkout string) int {
			layout := "2006-01-02"
			t1, err := time.Parse(layout, checkin)
			if err != nil {
				return 0
			}
			t2, err := time.Parse(layout, checkout)
			if err != nil {
				return 0
			}
			return int(t2.Sub(t1).Hours() / 24)
		},
		"tripDuration": func(departure, arrival string) string {
			layout := "2006-01-02"
			t1, err := time.Parse(layout, departure)
			if err != nil {
				return ""
			}
			t2, err := time.Parse(layout, arrival)
			if err != nil {
				return ""
			}
			days := int(t2.Sub(t1).Hours()/24) + 1
			nights := days - 1
			return fmt.Sprintf("%d Days %d Nights", days, nights)
		},
		"formatActivityDateTime": func(dateStr, timeStr string) string {
			layout := "2006-01-02 15:04"
			t, err := time.Parse(layout, dateStr+" "+timeStr)
			if err != nil {
				return dateStr
			}
			return t.Format("02-01-2006/03:04PM")
		},
		"formatCurrency": func(n int) string {
			s := fmt.Sprintf("%d", n)
			if len(s) <= 3 {
				return s
			}
			lastThree := s[len(s)-3:]
			otherNumbers := s[:len(s)-3]
			var result string
			for i, r := range []rune(otherNumbers) {
				if i != 0 && (len(otherNumbers)-i)%2 == 0 {
					result += ","
				}
				result += string(r)
			}
			return result + "," + lastThree
		},
	}
}
