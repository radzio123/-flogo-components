package filelist

// Imports
import (
	"io/ioutil"
	"strings"

	"path/filepath"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	ivDirectory = "directory"
	ivPattern = "pattern"
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
	directory := context.GetInput(ivDirectory).(string)
	pattern := context.GetInput(ivPattern).(string)

	files, error := ioutil.ReadDir(directory)
	
	if error != nil {
		log.Errorf("Error while reading files list %s\n", error.Error)
		return false, error
	}

	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}

	var all []string

	for _, file := range files {
		if !file.IsDir() {
			match, matchError := filepath.Match(pattern, file.Name())
			if matchError != nil {
				log.Errorf("Error matching pattern %s\n", matchError.Error)
				return false, matchError
			}
			if match {
				all = append(all, directory + file.Name())
			}
		}
	}

	context.SetOutput(ovResult, all)
	return true, nil
}