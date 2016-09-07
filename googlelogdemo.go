package googlelogdemo

/*
	Log entries to Google Cloud Logging service.

// Severity: Severity of log.
// Possible values:
//   "DEFAULT"
//   "DEBUG"
//   "INFO"
//   "NOTICE"
//   "WARNING"
//   "ERROR"
//   "CRITICAL"
//   "ALERT"
//   "EMERGENCY"
*/

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	logging "google.golang.org/api/logging/v1beta3"
	"log"
)

const (
	LevelDefault   = "DEFAULT"
	LevelDebug     = "DEBUG"
	LevelInfo      = "INFO"
	LevelNotice    = "NOTICE"
	LevelWarning   = "WARNING"
	LevelError     = "ERROR"
	LevelCritical  = "CRITICAL"
	LevelAlert     = "ALERT"
	LevelEmergency = "EMERGENCY"
)

//CloudLogger will write structured logs to the cloud.
type CloudLogger struct {
	service         *logging.ProjectsService
	logEntryService *logging.ProjectsLogsEntriesService
	projectID       string
	logsID          string
}

func New(serviceAccount, privateKey, projectID, logsID string) *CloudLogger {
	oauthConf := &jwt.Config{
		Email:      serviceAccount,
		PrivateKey: []byte(privateKey),
		Scopes:     []string{logging.LoggingWriteScope},
		TokenURL:   google.JWTTokenURL,
	}
	oauthHTTPClient := oauthConf.Client(oauth2.NoContext)
	loggingService, err := logging.New(oauthHTTPClient)
	if err != nil {
		panic(err)
	}

	s := logging.NewProjectsService(loggingService)
	return &CloudLogger{
		service:         s,
		projectID:       projectID,
		logsID:          logsID,
		logEntryService: s.Logs.Entries,
	}
}

func (cl *CloudLogger) WriteLogEntry(severity string, e interface{}) error {
	le := &logging.LogEntry{
		StructPayload: e,
		Metadata: &logging.LogEntryMetadata{
			ServiceName: "compute.googleapis.com",
			//Severity can be set to other levels once we need them.
			Severity: severity,
		},
	}
	req := &logging.WriteLogEntriesRequest{
		Entries: []*logging.LogEntry{le},
	}
	call := cl.logEntryService.Write(cl.projectID, cl.logsID, req)
	_, err := call.Do()
	if err != nil {
		log.Println("Got an error trying to write logEntry", err)
		return err
	}
	return nil
}
