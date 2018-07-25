package persistence

import "litesoftToDo/app/model"

type Persistor interface {
	/*
	 * Return a "rough" Count of the current lists.
	 */
	ListCount() (uint, error)

	/*
	 * Return a slice (could be empty, but not Nil) of the matching Lists.
	 *
	 * The SearchQuery string is matched if the Name is equal (case insensitive) OR the Description contains it (case insensitive).
	 *
	 * Note: The returned List(s) are not pointer(s), so that the data is copied!
	 */
	SearchLists(pSearchQuery string, pSkip, pLimit uint) ([]model.TodoList, error)

	/*
	 * Return a List by its ID.
	 *
	 * If not found the "returned" List reference will be Nil.
	 *
	 * Note: While the returned List is a pointer, if found, it is a pointer to a copy!
	 */
	GetList(pListId string) (*model.TodoList, error)

	/*
	 * Add a new List (which may not currently exist) with optional (which also may not currently exist) Tasks.
	 *
	 * The returned string is the ID of the List (returned, in case it needed to be generated)
	 *
	 * Note: List is not a pointer, so that the data is copied!
	 */
	AddList(pList model.TodoList) (string, error)

	/*
	 * Add a Task (which may not be an existing member of the existing List with the ListId).
	 *
	 * Note: Task is not a pointer, so that the data is copied!
	 */
	AddTask(pListId string, pTask model.Task) error

	/*
	 * Update the List (which must exist, identified by the ListId) with the updated Task (which must already be a member of the List).
	 *
	 * Note: Task is not a pointer, so that the data is copied!
	 */
	UpdateTask(pListId string, pTask model.Task) error
}
