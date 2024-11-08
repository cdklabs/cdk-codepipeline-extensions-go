package cdklabscdkcodepipelineextensions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-codepipeline-extensions-go/cdklabscdkcodepipelineextensions/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-codepipeline-extensions-go/cdklabscdkcodepipelineextensions/internal"
)

// A change controller.
//
// When added to a stage in a pipeline, this will check against
// a calendar and enable or disable the stage transition based off that calendar,
// defaulting to closed when the calendar cannot be found or when
// the check against it fails. It also checks to against alarms.
// Experimental.
type ChangeController interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChangeController
type jsiiProxy_ChangeController struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ChangeController) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewChangeController(scope constructs.Construct, id *string, props *ChangeControllerProps) ChangeController {
	_init_.Initialize()

	if err := validateNewChangeControllerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChangeController{}

	_jsii_.Create(
		"@cdklabs/cdk-codepipeline-extensions.ChangeController",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewChangeController_Override(c ChangeController, scope constructs.Construct, id *string, props *ChangeControllerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-codepipeline-extensions.ChangeController",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ChangeController_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChangeController_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-codepipeline-extensions.ChangeController",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChangeController) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

