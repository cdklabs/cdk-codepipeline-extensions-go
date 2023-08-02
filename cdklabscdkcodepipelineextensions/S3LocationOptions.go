package cdklabscdkcodepipelineextensions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Options for creating a calendar from a file in a S3 Bucket.
// Experimental.
type S3LocationOptions struct {
	// The name of the calendar file.
	// Experimental.
	CalendarName *string `field:"required" json:"calendarName" yaml:"calendarName"`
	// The bucket where the calendar is stored.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The role used for getting the calendar file.
	// Default: - A role created by the Custom Resource.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

