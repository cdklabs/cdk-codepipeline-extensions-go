package cdklabscdkcodepipelineextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-codepipeline-extensions-go/cdklabscdkcodepipelineextensions/jsii"
)

// The calendar for determining if pipeline stage should be open or closed.
// Experimental.
type Calendar interface {
	// The ARN of the calendar in SSM.
	// Experimental.
	CalendarArn() *string
	// Experimental.
	SetCalendarArn(val *string)
	// The name of the calendar.
	// Experimental.
	CalendarName() *string
	// Experimental.
	SetCalendarName(val *string)
}

// The jsii proxy struct for Calendar
type jsiiProxy_Calendar struct {
	_ byte // padding
}

func (j *jsiiProxy_Calendar) CalendarArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Calendar) CalendarName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"calendarName",
		&returns,
	)
	return returns
}


// Experimental.
func NewCalendar_Override(c Calendar) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-codepipeline-extensions.Calendar",
		nil, // no parameters
		c,
	)
}

func (j *jsiiProxy_Calendar)SetCalendarArn(val *string) {
	if err := j.validateSetCalendarArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarArn",
		val,
	)
}

func (j *jsiiProxy_Calendar)SetCalendarName(val *string) {
	if err := j.validateSetCalendarNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"calendarName",
		val,
	)
}

// Creates a calendar from a local file.
// Experimental.
func Calendar_Path(options *LocalPathOptions) Calendar {
	_init_.Initialize()

	if err := validateCalendar_PathParameters(options); err != nil {
		panic(err)
	}
	var returns Calendar

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-codepipeline-extensions.Calendar",
		"path",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a calendar from a S3 bucket.
// Experimental.
func Calendar_S3Location(options *S3LocationOptions) Calendar {
	_init_.Initialize()

	if err := validateCalendar_S3LocationParameters(options); err != nil {
		panic(err)
	}
	var returns Calendar

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-codepipeline-extensions.Calendar",
		"s3Location",
		[]interface{}{options},
		&returns,
	)

	return returns
}

