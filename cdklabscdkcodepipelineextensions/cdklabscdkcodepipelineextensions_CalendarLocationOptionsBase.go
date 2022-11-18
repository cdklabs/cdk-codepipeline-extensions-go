// This project is for use in the workshop DOP 401: Get better at building AWS CDK constructs.
package cdklabscdkcodepipelineextensions


// Options for creating a calendar object.
// Experimental.
type CalendarLocationOptionsBase struct {
	// The name of the calendar file.
	// Experimental.
	CalendarName *string `field:"required" json:"calendarName" yaml:"calendarName"`
}

