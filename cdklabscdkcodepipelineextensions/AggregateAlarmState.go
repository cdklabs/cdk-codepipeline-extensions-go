package cdklabscdkcodepipelineextensions


// Experimental.
type AggregateAlarmState struct {
	// Experimental.
	AlarmDetails *[]IAlarmDetail `field:"required" json:"alarmDetails" yaml:"alarmDetails"`
	// Experimental.
	State AlarmState `field:"required" json:"state" yaml:"state"`
	// Experimental.
	Summary *string `field:"required" json:"summary" yaml:"summary"`
}

