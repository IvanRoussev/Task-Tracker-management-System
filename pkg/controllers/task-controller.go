package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/IvanRoussev/taskManager/pkg/config"
	"github.com/IvanRoussev/taskManager/pkg/models"
	"github.com/IvanRoussev/taskManager/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var newTask models.Task

var db = config.GetDB()

func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks, err := models.GetAllTasks(db)

	if err != nil {
		fmt.Printf("Could not get tasks: %v", err)
	}

	res, _ := json.Marshal(newTasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		fmt.Printf("Could not get all Task as Requested: %v", err)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}

	utils.ParseBody(r, CreateTask)
	b := CreateTask.CreateTask(db)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		fmt.Printf("Could not create requested task: %v", err)
	}
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	fmt.Println(args)
	taskId := args["taskId"]
	fmt.Println(taskId)
	task := models.GetTaskById(db, taskId)

	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)

	if err != nil {
		fmt.Printf("Could not Fin Task by ID: %v error: %v", taskId, err)
	}

}
