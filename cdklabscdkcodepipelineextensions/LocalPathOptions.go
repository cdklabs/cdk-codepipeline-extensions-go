package cdklabscdkcodepipelineextensions


// Options for creating a calendar from a local file path.
// Experimental.
type LocalPathOptions struct {
	// The name of the calendar file.
	// Experimental.
	CalendarName *string `field:"required" json:"calendarName" yaml:"calendarName"`
	// The relative path to the calendar file.
	// Experimental.
	CalendarPath *string `field:"required" json:"calendarPath" yaml:"calendarPath"`
}

