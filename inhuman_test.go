package inhuman_test

import (
	"testing"
	"time"

	"github.com/umayr/inhuman"
)

func TestParse(t *testing.T) {
	pair := map[string]time.Duration{
		"2 h":                               (2 * time.Hour),
		"2 hr":                              (2 * time.Hour),
		"2 hrs":                             (2 * time.Hour),
		"2 hour":                            (2 * time.Hour),
		"2 hours":                           (2 * time.Hour),
		"20h":                               (20 * time.Hour),
		"20hr":                              (20 * time.Hour),
		"20hrs":                             (20 * time.Hour),
		"20hour":                            (20 * time.Hour),
		"20hours":                           (20 * time.Hour),
		"1 m":                               (1 * time.Minute),
		"1 min":                             (1 * time.Minute),
		"1 mins":                            (1 * time.Minute),
		"1 minute":                          (1 * time.Minute),
		"1 minutes":                         (1 * time.Minute),
		"21m":                               (21 * time.Minute),
		"21min":                             (21 * time.Minute),
		"21mins":                            (21 * time.Minute),
		"21minute":                          (21 * time.Minute),
		"21minutes":                         (21 * time.Minute),
		"7 s":                               (7 * time.Second),
		"7 sec":                             (7 * time.Second),
		"7 secs":                            (7 * time.Second),
		"7 second":                          (7 * time.Second),
		"7 seconds":                         (7 * time.Second),
		"17s":                               (17 * time.Second),
		"17sec":                             (17 * time.Second),
		"17secs":                            (17 * time.Second),
		"17second":                          (17 * time.Second),
		"17seconds":                         (17 * time.Second),
		"1 h 20 m":                          (1 * time.Hour) + (20 * time.Minute),
		"1 hr 20 min":                       (1 * time.Hour) + (20 * time.Minute),
		"1 hrs 20 mins":                     (1 * time.Hour) + (20 * time.Minute),
		"1 hour 20 minute":                  (1 * time.Hour) + (20 * time.Minute),
		"1 hours 20 minutes":                (1 * time.Hour) + (20 * time.Minute),
		"12h 2m":                            (12 * time.Hour) + (2 * time.Minute),
		"12hr 2min":                         (12 * time.Hour) + (2 * time.Minute),
		"12hrs 2mins":                       (12 * time.Hour) + (2 * time.Minute),
		"12hour 2minute":                    (12 * time.Hour) + (2 * time.Minute),
		"12hours 2minutes":                  (12 * time.Hour) + (2 * time.Minute),
		"1h20m":                             (1 * time.Hour) + (20 * time.Minute),
		"1hr20min":                          (1 * time.Hour) + (20 * time.Minute),
		"1hrs20mins":                        (1 * time.Hour) + (20 * time.Minute),
		"1hour20minute":                     (1 * time.Hour) + (20 * time.Minute),
		"1hours20minutes":                   (1 * time.Hour) + (20 * time.Minute),
		"3 h 32 s":                          (3 * time.Hour) + (32 * time.Second),
		"3 hr 32 sec":                       (3 * time.Hour) + (32 * time.Second),
		"3 hrs 32 secs":                     (3 * time.Hour) + (32 * time.Second),
		"3 hour 32 second":                  (3 * time.Hour) + (32 * time.Second),
		"3 hours 32 seconds":                (3 * time.Hour) + (32 * time.Second),
		"13h 2s":                            (13 * time.Hour) + (2 * time.Second),
		"13hr 2sec":                         (13 * time.Hour) + (2 * time.Second),
		"13hrs 2secs":                       (13 * time.Hour) + (2 * time.Second),
		"13hour 2second":                    (13 * time.Hour) + (2 * time.Second),
		"13hours 2seconds":                  (13 * time.Hour) + (2 * time.Second),
		"3h32s":                             (3 * time.Hour) + (32 * time.Second),
		"3hr32sec":                          (3 * time.Hour) + (32 * time.Second),
		"3hrs32secs":                        (3 * time.Hour) + (32 * time.Second),
		"3hour32second":                     (3 * time.Hour) + (32 * time.Second),
		"3hours32seconds":                   (3 * time.Hour) + (32 * time.Second),
		"1 m 30 s":                          (1 * time.Minute) + (30 * time.Second),
		"1 min 30 sec":                      (1 * time.Minute) + (30 * time.Second),
		"1 mins 30 secs":                    (1 * time.Minute) + (30 * time.Second),
		"1 minute 30 second":                (1 * time.Minute) + (30 * time.Second),
		"1 minutes 30 seconds":              (1 * time.Minute) + (30 * time.Second),
		"15m 130s":                          (15 * time.Minute) + (130 * time.Second),
		"15min 130sec":                      (15 * time.Minute) + (130 * time.Second),
		"15mins 130secs":                    (15 * time.Minute) + (130 * time.Second),
		"15minute 130second":                (15 * time.Minute) + (130 * time.Second),
		"15minutes 130seconds":              (15 * time.Minute) + (130 * time.Second),
		"1241m2430s":                        (1241 * time.Minute) + (2430 * time.Second),
		"1241min2430sec":                    (1241 * time.Minute) + (2430 * time.Second),
		"1241mins2430secs":                  (1241 * time.Minute) + (2430 * time.Second),
		"1241minute2430second":              (1241 * time.Minute) + (2430 * time.Second),
		"1241minutes2430seconds":            (1241 * time.Minute) + (2430 * time.Second),
		"3 h 40 m 100 s":                    (3 * time.Hour) + (40 * time.Minute) + (100 * time.Second),
		"3 hr 40 min 100 sec":               (3 * time.Hour) + (40 * time.Minute) + (100 * time.Second),
		"3 hrs 40 mins 100 secs":            (3 * time.Hour) + (40 * time.Minute) + (100 * time.Second),
		"3 hour 40 minute 100 second":       (3 * time.Hour) + (40 * time.Minute) + (100 * time.Second),
		"3 hours 40 minutes 100 seconds":    (3 * time.Hour) + (40 * time.Minute) + (100 * time.Second),
		"03h 4m 001s":                       (3 * time.Hour) + (4 * time.Minute) + (1 * time.Second),
		"03hr 4min 001sec":                  (3 * time.Hour) + (4 * time.Minute) + (1 * time.Second),
		"03hrs 4mins 001secs":               (3 * time.Hour) + (4 * time.Minute) + (1 * time.Second),
		"03hour 4minute 001second":          (3 * time.Hour) + (4 * time.Minute) + (1 * time.Second),
		"03hours 4minutes 001seconds":       (3 * time.Hour) + (4 * time.Minute) + (1 * time.Second),
		"3123h4990m139947s":                 (3123 * time.Hour) + (4990 * time.Minute) + (139947 * time.Second),
		"3123hr4990min139947sec":            (3123 * time.Hour) + (4990 * time.Minute) + (139947 * time.Second),
		"3123hrs4990mins139947secs":         (3123 * time.Hour) + (4990 * time.Minute) + (139947 * time.Second),
		"3123hour4990minute139947second":    (3123 * time.Hour) + (4990 * time.Minute) + (139947 * time.Second),
		"3123hours4990minutes139947seconds": (3123 * time.Hour) + (4990 * time.Minute) + (139947 * time.Second),
	}

	for k, v := range pair {
		r, err := inhuman.Parse(k)
		if err != nil {
			t.Error(err)
			return
		}

		if r != v {
			t.Errorf("expected %dns (%s) but got %dns", v, k, r)
		}
	}
}
