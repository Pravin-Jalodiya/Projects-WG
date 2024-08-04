package models

type UserData struct {
	Username    string
	Password    string
	ToDo        []Course
	DailyStatus []DailyStatus
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

//type UserProgress struct {
//	CompletedModules []float32
//}

type DailyStatus struct {
	Date            string
	Time            string
	TopicsCompleted []Module
}
