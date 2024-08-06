package models

type UserData struct {
	Username    string
	Password    string
	ToDo        []Course
	DailyStatus []DailyStatus
	Progress    []UserProgress
}

type Module struct {
	MID   float32
	Title string
}

type Course struct {
	CID     int
	Title   string
	Modules []Module
}

type UserProgress struct {
	CID        int
	Completion int
}

type DailyStatus struct {
	Date            string
	Time            string
	TopicsCompleted []Module
}
