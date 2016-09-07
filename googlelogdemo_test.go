package googlelogdemo

func ExampleCloudLogger() {
	logger := New("john@joetown.com", "MYKEY", "digerrity-doo", "logs1")
	err := logger.WriteLogEntry(LevelDebug, struct {
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
