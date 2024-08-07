package models

import "time"

type UserData struct {
	Username    string
	Password    string
	ToDo        []Course
	DailyStatus []DailyStatus
	Progress    UserProgress
	GeneralTodo []DoList
}

type DoList struct {
	Task     string
	Deadline time.Time
}

type Module struct {
	MID       float32
	Title     string
	Weightage int
}

type Course struct {
	CID     int
	Title   string
	Modules []Module
}

type UserProgress struct {
	Courses         []Course
	ModulesFinished []float32
}

type DailyStatus struct {
	Date            string
	Time            string
	TopicsCompleted []Module
}
