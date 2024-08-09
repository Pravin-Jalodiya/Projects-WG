package readers

import (
	"projects/models"
	"reflect"
	"testing"
)

func TestFReaderCourses(t *testing.T) {
	type args struct {
		f    string
		flag int
	}
	tests := []struct {
		name string
		args args
		want []models.Course
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FReaderCourses(tt.args.f, tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FReaderCourses() = %v, want %v", got, tt.want)
			}
		})
	}
}
