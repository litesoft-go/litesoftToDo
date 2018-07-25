package persistence

import (
	"fmt"
	"strings"
	"container/list"
	"litesoftToDo/app/model"
	"litesoftToDo/utils/rw"
	"litesoftToDo/utils/uuid"
	"litesoftToDo/utils/sets"
)

/*
 * Note: all interaction with this InMemoryPersistor are Synchronized!
 */

type enhancedToDoList struct {
	mNameLC        string
	mDescriptionLC string
	mList          *model.TodoList
}

func new_enhancedToDoList(pList *model.TodoList) *enhancedToDoList {
	return &enhancedToDoList{
		mNameLC:        strings.ToLower(pList.Name),
		mDescriptionLC: strings.ToLower(pList.Description),
		mList:          pList}
}

func (this *enhancedToDoList) selectedBy(pSearchQueryLC string) bool {
	return (this.mNameLC == pSearchQueryLC) || strings.Contains(this.mDescriptionLC, pSearchQueryLC)
}

func (this *enhancedToDoList) getList() (rList *model.TodoList) {
	if this != nil {
		rList = this.mList
	}
	return
}

func (this *enhancedToDoList) isOrderedBefore(them *enhancedToDoList) bool {
	if this.mNameLC < them.mNameLC {
		return true
	}
	if this.mNameLC > them.mNameLC {
		return false
	}
	return this.mDescriptionLC < them.mDescriptionLC
}

type InMemoryPersistor struct {
	mUUIDsource       uuid.UUIDsource
	mMutex            rw.Mutex
	mLinkedList       *list.List // for Ordering
	mListElementsById map[string]*list.Element
	mExistingTaskIDs  sets.StringSet // for "global uniqueness of Tasks (which IMO makes no sense)
}

func NewInMemoryPersistor(pUUIDsource uuid.UUIDsource) *InMemoryPersistor {
	return &InMemoryPersistor{
		mUUIDsource:       pUUIDsource,
		mMutex:            rw.NewMutex(),
		mLinkedList:       list.New(),
		mListElementsById: make(map[string]*list.Element),
		mExistingTaskIDs:  sets.NewString()}
}

/*
 * Return a "rough" Count of the current lists.
 */
func (this *InMemoryPersistor) ListCount() (rCount uint, err error) {
	if this != nil {
		defer this.mMutex.RLock().Unlock()
		rCount = uint(len(this.mListElementsById))
	}
	return
}

/*
 * Return a slice (could be empty, but not Nil - unless there is an error) of the matching Lists.
 *
 * The SearchQuery string is matched if the Name is equal (case insensitive) OR the Description contains it (case insensitive).
 */
func (this *InMemoryPersistor) SearchLists(pSearchQuery string, pSkip, pLimit uint) (rFound []model.TodoList, err error) {
	pSearchQuery = strings.ToLower(pSearchQuery)
	rFound = []model.TodoList{}
	if (this != nil) && (pLimit > 0) {
		defer this.mMutex.RLock().Unlock()
		for zElement := this.mLinkedList.Front(); zElement != nil; zElement = zElement.Next() {
			zEnhancedToDoList := enhancedToDoListFromElement(zElement)
			if zEnhancedToDoList.selectedBy(pSearchQuery) {
				if pSkip > 0 {
					pSkip--
				} else {
					rFound = append(rFound, *zEnhancedToDoList.getList())
					pLimit--
					if pLimit == 0 {
						return
					}
				}
			}
		}
	}
	return
}

/*
 * Return a List by its ID.
 *
 * If not found the "returned" List reference will be Nil.
 */
func (this *InMemoryPersistor) GetList(pListId string) (rInstance *model.TodoList, err error) {
	if this == nil {
		return
	}
	defer this.mMutex.RLock().Unlock()
	zInstance, err := this.getListUnderLock(pListId)
	if (err == nil) && (zInstance != nil) {
		rInstance = zInstance.Clone() // force a copy to be created (protect the version in the "persistor")
	}
	return
}

/*
 * Add a new List (which may not currently exist) with optional (but not currently existing Tasks).
 */
func (this *InMemoryPersistor) AddList(pList model.TodoList) (rListId string, err error) {
	err = ensureThisNotNil(this, "AddList", pList.Validate())
	if err != nil {
		return
	}
	defer this.mMutex.WLock().Unlock()
	zListId, err := pList.EnsureId(this.uniqueListIdSource) // Set ID if not there
	if err != nil {
		return
	}
	if this.existingListId(zListId) { // Check that ID does NOT already exist
		err = fmt.Errorf("attempt to AddList with list with already existing ID of: %s", zListId)
		return
	}
	zTaskIds, err := pList.EnsureTaskIds(this.uniqueTaskIdSource) // Set IDs if not there
	if err != nil {
		return
	}
	for _, zTaskId := range zTaskIds { // Check that the Task IDs do not already exist
		err = this.ensureNotExistingTaskId("AddList's list contains", zTaskId)
		if err != nil {
			return
		}
	}
	// OK it is safe to add...

	zEnhancedToDoList := new_enhancedToDoList(&pList) // 1st create enhancedToDoList

	zElement := this.addToLinkedList(zEnhancedToDoList) // 2nd insert into LinkedList (creating *Element)

	this.mListElementsById[zListId] = zElement // 3rd add Element to List Map

	for _, zTaskId := range zTaskIds { // 4th add TaskIds to Existing TaskID Set
		this.mExistingTaskIDs.Add(zTaskId)
	}

	rListId = zListId
	return
}

/*
 * Add a Task (which may not be an existing member of the existing List with the ListId).
 */
func (this *InMemoryPersistor) AddTask(pListId string, pTask model.Task) (err error) {
	err = ensureThisNotNil(this, "AddTask", pTask.Validate())
	if err != nil {
		return
	}
	defer this.mMutex.WLock().Unlock()
	zTaskId, err := pTask.EnsureId(this.uniqueTaskIdSource) // Set ID if not there
	if err != nil {
		return
	}
	err = this.ensureNotExistingTaskId("AddTask's", zTaskId)
	if err != nil {
		return
	}
	zList, err := this.getRequiredList("AddTask", pListId)
	if err != nil {
		return
	}
	// OK it is safe to add...
	zList.AddTask(&pTask)
	this.mExistingTaskIDs.Add(zTaskId)
	return
}

/*
 * Update the List (which must exist, identified by the ListId) with the updated Task (which must already be a member of the List).
 */
func (this *InMemoryPersistor) UpdateTask(pListId string, pTask model.Task) (err error) {
	err = ensureThisNotNil(this, "UpdateTask", pTask.Validate())
	if err != nil {
		return
	}
	defer this.mMutex.WLock().Unlock()
	zList, err := this.getRequiredList("UpdateTask", pListId)
	if err != nil {
		return
	}
	zTaskId := pTask.Id
	zExistingTask := zList.GetTask(zTaskId)
	if zExistingTask == nil {
		err = fmt.Errorf("unable to update Task (id:%s) on List (id:%s), as list doesnot not currently have that task", zTaskId, pListId)
		return
	}
	// OK it is safe to add...
	zExistingTask.UpdateWithValidated(&pTask)
	return
}

func ensureThisNotNil(this *InMemoryPersistor, pWhat string, pError error) error {
	if (pError == nil) && (this == nil) {
		pError = fmt.Errorf("can't %s to nil Persistor", pWhat)
	}
	return pError
}

func (this *InMemoryPersistor) getRequiredList(pWhat, pListId string) (rInstance *model.TodoList, err error) {
	rInstance, err = this.getListUnderLock(pListId)
	if (err == nil) && (rInstance == nil) {
		err = fmt.Errorf("attempt to %s, but list does NOT exist with ID: %s", pWhat, pListId)
	}
	return
}

func (this *InMemoryPersistor) getListUnderLock(pListId string) (rInstance *model.TodoList, err error) {
	zEnhancedToDoList := enhancedToDoListFromElement(this.mListElementsById[pListId])
	rInstance = zEnhancedToDoList.getList()
	return
}

func enhancedToDoListFromElement(pElement *list.Element) *enhancedToDoList {
	if pElement != nil {
		return pElement.Value.(*enhancedToDoList) // NOT checking for failure!
	}
	return nil
}

func (this *InMemoryPersistor) ensureNotExistingTaskId(pWhat string, pTaskId string) (err error) {
	if this.existingTaskId(pTaskId) {
		err = fmt.Errorf("%s Task with already existing ID of: %s", pWhat, pTaskId)
	}
	return
}

func (this *InMemoryPersistor) existingListId(pId string) bool {
	return nil != this.mListElementsById[pId]
}

func (this *InMemoryPersistor) existingTaskId(pId string) bool {
	return this.mExistingTaskIDs.Contains(pId)
}

func (this *InMemoryPersistor) uniqueListIdSource() (rUUID string, err error) {
	rUUID, err = this.uniqueIdSource("ListId", this.existingListId)
	return
}

func (this *InMemoryPersistor) uniqueTaskIdSource() (rUUID string, err error) {
	rUUID, err = this.uniqueIdSource("TaskId", this.existingTaskId)
	return
}

func (this *InMemoryPersistor) uniqueIdSource(pWhat string, pCheckExistsFunc func(string) bool) (rUUID string, err error) {
	var zPrevErr, zCurErr error
	for i := 0; i < 9; i++ {
		rUUID, zCurErr = this.mUUIDsource()
		if zCurErr == nil {
			if !pCheckExistsFunc(rUUID) {
				return
			}
			zPrevErr = nil // clear
		} else {
			if zPrevErr != nil { // 2 errors in a row - return current error
				err = fmt.Errorf("unable to generate '%s': %v", pWhat, zCurErr)
				return
			}
			zPrevErr = zCurErr
		}
	}
	err = fmt.Errorf("unable to generate '%s': max attempts exceeded", pWhat)
	return
}

func (this *InMemoryPersistor) addToLinkedList(pEList *enhancedToDoList) *list.Element {
	zElement := this.mLinkedList.Front()
	if zElement == nil {
		return this.mLinkedList.PushFront(pEList)
	}
	for enhancedToDoListFromElement(zElement).isOrderedBefore(pEList) {
		zElement = zElement.Next()
		if zElement == nil {
			return this.mLinkedList.PushBack(pEList)
		}
	}
	return this.mLinkedList.InsertBefore(pEList, zElement)
}
