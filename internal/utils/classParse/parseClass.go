package classParse

import (
	ics "github.com/arran4/golang-ical"
	"homeworktodolist/internal/service/class"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Class struct {
	Summary     string
	Start       string
	End         string
	Description string
	Location    string
}

var (
	summary     string
	start       string
	end         string
	location    string
	description string
)

func IcalParse(icalLink string) ([]class.UpdateClass, error) {

	resp, err := http.Get(icalLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)

	bodyString = strings.Replace(bodyString, "X-SCHEDULE_VERSION-ID:8", "X-SCHEDULE-VERSION-ID:8", -1)

	calendar, err := ics.ParseCalendar(strings.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	var classes []class.UpdateClass
	for _, event := range calendar.Events() {
		class := toClass(event)
		if class.Description == "" && class.Location == "" {

		}
		classes = append(classes, class.toClass())
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].Start.Before(classes[j].Start)
	})

	return classes, nil
}

func (c *Class) toClass() class.UpdateClass {
	start, _ := time.Parse("20060102T150405", c.Start)
	end, _ := time.Parse("20060102T150405", c.End)
	return class.UpdateClass{
		Summary:     c.Summary,
		Start:       start,
		End:         end,
		Description: c.Description,
		Location:    c.Location,
	}
}
