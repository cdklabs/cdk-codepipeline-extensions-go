package cdklabscdkcodepipelineextensions


// searchTerms: a list of terms to match in the alarm description.
// Experimental.
type GetAlarmStateOptions struct {
	// Experimental.
	SearchTerms *[]*string `field:"required" json:"searchTerms" yaml:"searchTerms"`
}

