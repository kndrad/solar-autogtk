package main

import (
	"autogtk/log"
	"autogtk/sunphase"
	"autogtk/theme"
	"flag"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// Command-line flags for setting light theme, dark theme, and geographical coordinates.
var (
	light     = flag.String("light-theme", "Light", "system light theme name")
	dark      = flag.String("dark-theme", "Dark", "system dark theme name")
	latitude  = flag.Float64("latitude", 50.26, "host latitude")
	longitude = flag.Float64("longitude", 19.02, "host longitude")
)

var logger = log.New()

func main() {
	flag.Parse()

	cr := cron.New()

	// Function to update theme based on solar phases
	updateTransitions := func() {
		// Remove all previous cron jobs
		for _, entry := range cr.Entries() {
			cr.Remove(entry.ID)
		}

		// Schedule light theme at sunrise
		sunrise := sunphase.Sunrise(*latitude, *longitude)
		id, err := cr.AddFunc(
			fmt.Sprintf("%d %d * * *", sunrise.Minute(), sunrise.Hour()), func() {
				err := theme.Set(*light)
				if err != nil {
					logger.Errorf("Error while switching to %s theme: %s.", *light, err.Error())
				}
			},
		)
		if err != nil {
			logger.Fatalf("Error scheduling %s theme: %s.", *light, err.Error())
		}
		logger.Infof("(%d): %s theme will switch at: %s.", id, *light, sunrise.Format(time.TimeOnly))

		// Schedule dark theme at sunset
		sunset := sunphase.Sunset(*latitude, *longitude)
		id, err = cr.AddFunc(
			fmt.Sprintf("%d %d * * *", sunset.Minute(), sunset.Hour()), func() {
				err := theme.Set(*dark)
				if err != nil {
					logger.Errorf("Error while switching to %s theme: %s.", *dark, err.Error())
				}
			},
		)
		if err != nil {
			logger.Fatalf("Error scheduling %s theme: %s.", *dark, err.Error())
		}
		logger.Infof("(%d): %s theme will switch at: %s.", id, *dark, sunset.Format(time.TimeOnly))
	}

	// Update theme transitions initially
	updateTransitions()

	// Schedule daily update of theme transitions at 1 am
	id, err := cr.AddFunc("0 1 * * *", updateTransitions)
	if err != nil {
		logger.Fatalf("Error while scheduling daily theme transitions %s.", err.Error())
	}
	logger.Infof("Scheduled daily theme transitions: %d", id)

	select {}
}
