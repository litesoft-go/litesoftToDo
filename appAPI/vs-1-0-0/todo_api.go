/*
 * Simple To Do API
 *
 * API version: 1.0.0
 * Original Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vs_1_0_0

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"litesoftToDo/openAPI"
	"litesoftToDo/app/persistence"
	"litesoftToDo/utils/uints"
	"litesoftToDo/utils/validate"
)

const DEFAULT_SEARCH_LIMIT = 10

type Methods struct {
	mPersistor persistence.Persistor
}

func MethodsNew(pPersistor persistence.Persistor) *Methods {
	return &Methods{mPersistor: pPersistor}
}

// AddGet("/lists") +
//
// searchString 	string (query)				pass an optional search string for looking up a list
// skip				integer($int32) (query)		number of records to skip for pagination
// limit			integer($int32) (query)		maximum number of records to return
//
// Result "array" of vs_1_0_0
func (this Methods) SearchLists(w http.ResponseWriter, r *http.Request) {
	zQP := openAPI.NewQueryParams(r.URL.Query())
	zSearchString := zQP.GetString("searchString")
	zSkip := zQP.GetUintOr("skip", 0)
	zLimit := zQP.GetUintOr("limit", DEFAULT_SEARCH_LIMIT)
	if !uints.IsBetweenInclusive(1, zLimit, 1000) {
		zLimit = DEFAULT_SEARCH_LIMIT
	}
	zModelLists, err := this.mPersistor.SearchLists(zSearchString, zSkip, zLimit)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	zTodoLists := FromModelTodoLists(zModelLists)
	zBytes, err := json.Marshal(zTodoLists)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
	} else {
		openAPI.NewJsonResponseOK(string(zBytes)).ApplyTo(w)
	}
}

// AddRoutePost("/lists").
func (this Methods) AddList(w http.ResponseWriter, r *http.Request) {
	var zList TodoList
	zNoBody, err := openAPI.HttpRequestBodyToJsonBinding(&zList, r.Body)
	if zNoBody {
		openAPI.NewInvalidTextResponse("No Post Body provided").ApplyTo(w)
		return
	}
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	zModelList := zList.ToModelTodoList()
	zListId, err := this.mPersistor.AddList(zModelList)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	zBytes, err := json.Marshal(&ListCreated{ListId: zListId})
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
	} else {
		openAPI.NewJsonResponseCreated(string(zBytes)).ApplyTo(w)
	}
}

// AddGet("/list/{id}").
func (this Methods) GetList(w http.ResponseWriter, r *http.Request) {
	zPP := openAPI.NewPathParams(mux.Vars(r))
	zListId := zPP.GetString("id")
	err := validate.NoWhiteSpace("list ID", zListId)
	if (len(zListId) == 0) || (err != nil) {
		openAPI.NewInvalidTextResponse("List ID invalid: '" + zListId + "'").ApplyTo(w)
		return
	}
	zModelList, err := this.mPersistor.GetList(zListId)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	if zModelList == nil {
		openAPI.NewNotFoundTextResponse("List not found for ID: '" + zListId + "'").ApplyTo(w)
		return
	}
	zTodoList := FromModelTodoList(zModelList)
	zBytes, err := json.Marshal(zTodoList)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
	} else {
		openAPI.NewJsonResponseOK(string(zBytes)).ApplyTo(w)
	}
}

// AddRoutePost("/list/{id}/tasks").
func (this Methods) AddTask(w http.ResponseWriter, r *http.Request) {
	zPP := openAPI.NewPathParams(mux.Vars(r))
	zListId := zPP.GetString("id")
	err := validate.NoWhiteSpace("list ID", zListId)
	if (len(zListId) == 0) || (err != nil) {
		openAPI.NewInvalidTextResponse("List ID invalid: '" + zListId + "'").ApplyTo(w)
		return
	}

	var zTask Task
	zNoBody, err := openAPI.HttpRequestBodyToJsonBinding(&zTask, r.Body)
	if zNoBody {
		openAPI.NewInvalidTextResponse("No Post Body provided").ApplyTo(w)
		return
	}
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	zModelTask := zTask.ToModelTask()
	err = this.mPersistor.AddTask(zListId, zModelTask)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	openAPI.NewTextResponseCreated("Task").ApplyTo(w)
}

// AddRoutePost("/list/{id}/task/{taskId}/complete")
func (this Methods) CompleteTask(w http.ResponseWriter, r *http.Request) {
	zPP := openAPI.NewPathParams(mux.Vars(r))
	zListId := zPP.GetString("id")
	err := validate.NoWhiteSpace("list ID", zListId)
	if (len(zListId) == 0) || (err != nil) {
		openAPI.NewInvalidTextResponse("List ID invalid: '" + zListId + "'").ApplyTo(w)
		return
	}
	zTaskId := zPP.GetString("taskId")
	err = validate.NoWhiteSpace("task ID", zTaskId)
	if (len(zTaskId) == 0) || (err != nil) {
		openAPI.NewInvalidTextResponse("Task ID invalid: '" + zTaskId + "'").ApplyTo(w)
		return
	}

	var zCompletedTask CompletedTask
	zNoBody, err := openAPI.HttpRequestBodyToJsonBinding(&zCompletedTask, r.Body)
	if zNoBody {
		openAPI.NewInvalidTextResponse("No Post Body provided").ApplyTo(w)
		return
	}
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}

	zModelList, err := this.mPersistor.GetList(zListId)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	if zModelList == nil {
		openAPI.NewInvalidTextResponse("No List for ID: '" + zTaskId + "'").ApplyTo(w)
		return
	}
	zModelTask := zModelList.GetTask(zTaskId)
	if zModelTask == nil {
		openAPI.NewInvalidTextResponse("No Task with ID '" + zTaskId + "' on List for ID: '" + zTaskId + "'").ApplyTo(w)
		return
	}

	if zModelTask.Completed == zCompletedTask.Completed {
		openAPI.NewTextResponseOK("Task already in requested 'completion' state").ApplyTo(w)
		return
	}

	zModelTask = zModelTask.Clone()
	zModelTask.Completed = zCompletedTask.Completed

	err = this.mPersistor.UpdateTask(zListId, *zModelTask)
	if err != nil {
		openAPI.NewFailedResponse(err).ApplyTo(w)
		return
	}
	openAPI.NewTextResponseUpdated("Task").ApplyTo(w)
}
