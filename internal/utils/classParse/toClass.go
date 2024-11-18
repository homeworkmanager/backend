package classParse

import ics "github.com/arran4/golang-ical"

func toClass(event *ics.VEvent) Class {

	if event.GetProperty(ics.ComponentPropertySummary) != nil {
		summary = event.GetProperty(ics.ComponentPropertySummary).Value
	}

	if event.GetProperty(ics.ComponentPropertyDescription) != nil {
		description = event.GetProperty(ics.ComponentPropertyDescription).Value
	}

	if event.GetProperty(ics.ComponentPropertyDtStart) != nil {
		start = event.GetProperty(ics.ComponentPropertyDtStart).Value
	}

	if event.GetProperty(ics.ComponentPropertyDtEnd) != nil {
		end = event.GetProperty(ics.ComponentPropertyDtEnd).Value
	}

	if event.GetProperty(ics.ComponentPropertyLocation) != nil {
		location = event.GetProperty(ics.ComponentPropertyLocation).Value
	}

	return Class{
		Summary:     summary,
		Start:       start,
		End:         end,
		Description: description,
		Location:    location,
	}

}
