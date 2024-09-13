/*
Copyright © 2024 Enrique Marín

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package models

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          int
	Name        string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func Add(db *gorm.DB, task Task) *Task {
	task.Completed = false
	db.Create(&task)
	return &task
}

func GetAll(db *gorm.DB) []Task {
	var tasks []Task
	db.Order("created_at desc").Find(&tasks)
	return tasks
}

func GetByID(db *gorm.DB, id int) Task {
	var task Task
	db.Find(&task, id)
	return task
}

func DeleteByID(db *gorm.DB, id int) *gorm.DB {
	result := db.Delete(&Task{}, id)
	if result.Error != nil {
		log.Println(result.Error)

		os.Exit(1)
	}

	return result
}

func UpdateByID(db *gorm.DB, id int, task Task) *Task {
	var taskToUpdate Task
	resultTask := db.First(&taskToUpdate, id)
	if resultTask.Error != nil {
		log.Println(resultTask.Error)

		os.Exit(1)
	}

	taskToUpdate = task

	resultSave := db.Save(&taskToUpdate)

	if resultSave.Error != nil {
		log.Println(resultSave.Error)

		os.Exit(1)
	}

	return &taskToUpdate
}
