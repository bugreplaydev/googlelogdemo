package googlelogdemo

import "net/http"

func ExampleCloudLogger_gce() {
	logger, err := New("myserviceaccount", "myprivatekey", "myprojectid", "mylogname")
	if err != nil {
		panic(err)
	}
	err = logger.WriteLogEntry(LevelDebug, struct {
		EventName string
		EventID   string
	}{
		EventName: "NewUser",
		EventID:   "qxzzr65",
	})
	if err != nil {
		panic(err)
	}
}

func ExampleCloudLogger_appengine() {
	//normally you'd get the request from a Handler,
	r := new(http.Request)
	logger, err := NewAppEngineLogger(r, "myprojectid", "mylogsID")
	err = logger.WriteLogEntry(LevelDebug, struct {
		EventName string
		EventID   string
	}{
		EventName: "NewUser",
		EventID:   "qxzzr65",
	})
	if err != nil {
		panic(err)
	}
}
