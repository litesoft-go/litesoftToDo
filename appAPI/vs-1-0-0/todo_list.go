/*
 * Simple To Do API
 *
 * This is a simple API for managing a TO DO List
 *
 * API version: 1.0.0
 * Originally Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vs_1_0_0

import "litesoftToDo/app/model"

type TodoList struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	Tasks []Task `json:"tasks,omitempty"`
}

func (this *TodoList) Init() {
	if this != nil {
		this.Tasks = []Task{}
	}
}

func (src *TodoList) ToModelTodoList() (dst model.TodoList) {
	if src != nil {
		dst.Init()

		dst.Id = src.Id
		dst.Name = src.Name
		dst.Description = src.Description

		if src.Tasks != nil {
			for _, zTask := range src.Tasks {
				zDstTask := zTask.ToModelTask()
				dst.Tasks = append(dst.Tasks, &zDstTask )
			}
		}
	}
	return
}

func FromModelTodoList(src *model.TodoList) (dst TodoList) {
	if src != nil {
		dst.Init()

		dst.Id = src.Id
		dst.Name = src.Name
		dst.Description = src.Description

		if src.Tasks != nil {
			for _, zTask := range src.Tasks {
				dst.Tasks = append(dst.Tasks, FromModelTask(zTask))
			}
		}
	}
	return
}

func FromModelTodoLists(src []model.TodoList) (dst []TodoList) {
	dst = []TodoList{}
	if src != nil {
		for _, zEntry := range src {
			dst = append(dst, FromModelTodoList(&zEntry))
		}
	}
	return
}