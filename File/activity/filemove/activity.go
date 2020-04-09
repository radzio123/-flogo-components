// Package readfile implements a file reader for Flogo
package filelist

// Imports
import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivCurrentPath = "currentPath"
	ivWhereToMove = "whereToMove"
	ovResult   = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-listfiles")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the actions
	currentPath := context.GetInput(ivCurrentPath).(string)
	whereToMove := context.GetInput(ivWhereToMove).(string)

	error := os.Rename(currentPath, whereToMove)

	if error != nil {
		log.Errorf("Error moving file %s\n", error.Error)
		context.SetOutput(ovResult, false)
		return false, error
	}

	context.SetOutput(ovResult, true)
	return true, nil
}