package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ToDo struct {
	Title, Status string
}

type EdiToDo struct {
	Title, NewTitle, NewStatus string
}

type Task int

var todoSlice []ToDo

func(t *Task) GetToDo(title string, reply *ToDo) error {
	var found ToDo

	for _, v:= range todoSlice {
		if v.Title == title {
			found = v
		}
	}

	*reply = found
	return nil
}

func (t *Task) GetSlice(title string, reply *[]ToDo) error {
	*reply = todoSlice
	return nil
}

func (t *Task) MakeToDo(todo ToDo, reply *ToDo) error {
	todoSlice = append(todoSlice, todo)
	*reply = todo
	return nil
}

func (t *Task) EdiToDo(todo EdiToDo, reply *ToDo) error {
	var edited ToDo

	for i, v := range todoSlice {
		if v.Title == todo.Title {
			todoSlice[i] = ToDo{todo.NewTitle, todo.NewStatus}
			edited = ToDo{todo.NewTitle, todo.NewStatus}
		}
	}

	*reply = edited
	return nil
}

func (t *Task) DeleteToDo(todo ToDo, reply *ToDo) error {
	var deleted ToDo
	for i, v := range todoSlice {
		if v.Title == todo.Title && v.Status == todo.Status {
			todoSlice = append(todoSlice[:i], todoSlice[i+1:]...)
			deleted = todo
			break
		}
	}

	*reply = deleted
	return nil
}

func main() {
	task := new(Task)
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("Format of service Task isn't correct.", err)
	}

	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}