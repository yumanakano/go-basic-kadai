package main

import "fmt"

// ToDoアイテムを表す構造体
type ToDo struct {
	ID        int
	Title     string
	Completed bool
}

// ToDoを完了状態にするメソッド
func (todo *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", todo.ID)
	todo.Completed = true
}

// printTodos()関数を定義
func printTodos(todos []ToDo) {
	for _, todo := range todos {
		if todo.Completed {
			fmt.Printf("[完了]（ID：%d）%s\n", todo.ID, todo.Title)
		} else {
			fmt.Printf("[未完了]（ID：%d）%s\n", todo.ID, todo.Title)
		}
	}
}

func main() {
	// 1. ToDoスライスを作成
	todos := []ToDo{
		// 2. 3つのToDoアイテムを追加
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎文法", Completed: false},
	}

	// 3. 初期状態を表示
	printTodos(todos)

	// 4. IDが1・2のToDoを完了状態に更新
	todos[0].Complete()
	todos[1].Complete()

	// 5. 最終状態を表示
	printTodos(todos)
}
