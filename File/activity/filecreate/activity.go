package filecreate

// Imports
import (
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivPath = "path"
	ivContent = "content"
	ovResult   = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-file-create")

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
	path := context.GetInput(ivPath).(string)
	content := context.GetInput(ivContent).(string)

	if _, error := os.Stat("file-exists.go"); error == nil {
		log.Errorf("File exists")
		return false, error
	}

	f, error := os.Create(path)
	if error != nil {
		log.Errorf("Error creating file %s\n", error.Error)
		return false, error
	}

	_ ,error = f.WriteString(content);
	if error != nil {
		log.Errorf("Error writing to file %s\n", error.Error)
		return false, error
	}

	return true, nil
}