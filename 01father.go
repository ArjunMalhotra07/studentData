package main

import (
	"fmt"
)

func studentByFather(fatherName string) ([]StudentData, error) {
	var students []StudentData

	rows, err := db.Query("SELECT * FROM student_Data WHERE fatherName = ?", fatherName)
	if err != nil {
		return nil, fmt.Errorf("error1: no student found with fatherName = %s, %v,  ", fatherName, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb StudentData
		if err := rows.Scan(&alb.Id, &alb.Cgpa, &alb.StudentId, &alb.FatherName, &alb.MotherName, &alb.StudentName, &alb.City); err != nil {
			return nil, fmt.Errorf("error2: no student found with fathername = %s, %v ", fatherName, err)
		}
		students = append(students, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error3 : no student found with fathername = %s, %v,  ", fatherName, err)
	}
	return students, nil
}
