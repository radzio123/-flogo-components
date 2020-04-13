package filemove

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
var log = logger.GetLogger("activity-movefile")

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
		log.Errorf("Error moving file %s\n", error)
		return false, error
	}

	return true, nil
}