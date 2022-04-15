package main

import (
	"fmt"
)

func studentById(studentId string) ([]StudentData, error) {

	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE studentId = ?", studentId)
	if err != nil {
		return nil, fmt.Errorf("error1: no student found with studentId = %s, %v,  ", studentId, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.Id, &alb.Cgpa, &alb.StudentId, &alb.FatherName, &alb.MotherName, &alb.StudentName, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with studentId = %s, %v ", studentId, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with studentId = %s, %v,  ", studentId, err)
	}
	return students, nil
}
