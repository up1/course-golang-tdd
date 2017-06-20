package task

import "testing"

func TestNewTodoWithEmptyTitle(t *testing.T) {
	_, err := NewTodo("")
	if err == nil {
		t.Error("Expected 'Empty Title', error, got nil")
	}
}

func TestNewTodo(t *testing.T) {
	title := "My Title"
	todo, _ := NewTodo(title)
	if todo.Title != title {
		t.Errorf("Expected title %q, got %q", title, todo.Title)
	}
	if todo.Done {
		t.Errorf("New todo is done")
	}
}

func TestSaveOneTodoAndGetOne(t *testing.T) {
	todo, err := NewTodo("My todo")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	taskManager := NewTodoManager()
	taskManager.save(todo)

	allTask := taskManager.GetAll()
	if len(allTask) != 1 {
		t.Errorf("Expected 1 task, got %v", len(allTask))
	}

	if *allTask[0] != *todo {
		t.Errorf("Expected %v, got %v", todo, allTask)
	}
}

func TestSaveTwoAndGetTwo(t *testing.T) {
	todo1, err := NewTodo("My todo 1")
	if err != nil {
		t.Errorf("New todo : %v", err)
	}

	todo2, err := NewTodo("My todo 2")
	if err != nil {
		t.Errorf("New todo : %v", err)
	}

	todoManager := NewTodoManager()
	todoManager.save(todo1)
	todoManager.save(todo2)

	allTodo := todoManager.GetAll()
	if len(allTodo) != 2 {
		t.Errorf("Expected 2 todo, got %v", len(allTodo))
	}
}

func TestSaveUpdateWithDoneAndGet(t *testing.T) {
	todo, err := NewTodo("My todo")
	if err != nil {
		t.Errorf("New todo : %v", err)
	}

	todoManager := NewTodoManager()
	todoManager.save(todo)

	todo.Done = true

	allTodo := todoManager.GetAll()
	if allTodo[0].Done {
		t.Errorf("Save todo not done")
	}
}

func TestDuplicateSaveAndGetOne(t *testing.T) {
	todo, err := NewTodo("My todo")
	if err != nil {
		t.Errorf("New todo : %v", err)
	}

	todoManager := NewTodoManager()
	todoManager.save(todo)
	todoManager.save(todo)

	allTodo := todoManager.GetAll()
	if len(allTodo) != 1 {
		t.Errorf("Expected 1 todo, got %v", len(allTodo))
	}
}

func TestSaveAndFindByID(t *testing.T) {
	todo, err := NewTodo("My task")
	if err != nil {
		t.Errorf("New task : %v", err)
	}

	todoManager := NewTodoManager()
	todoManager.save(todo)

	myTodo, ok := todoManager.Find(todo.ID)
	if !ok {
		t.Errorf("Todo not found")
	}

	if *todo != *myTodo {
		t.Errorf("Expected %v, got %v", todo, myTodo)
	}
}

func TestSaveFindAndEdit(t *testing.T) {
	todo, err := NewTodo("My todo")
	if err != nil {
		t.Errorf("New todo : %v", err)
	}

	todoManager := NewTodoManager()
	todoManager.save(todo)

	myTask, _ := todoManager.Find(todo.ID)
	myTask.Done = true
	todoManager.save(myTask)

	allTodo := todoManager.GetAll()
	if !allTodo[0].Done {
		t.Errorf("Save todo not done")
	}

	if len(allTodo) != 1 {
		t.Errorf("Expected 1 todo, got %v", len(allTodo))
	}
}
