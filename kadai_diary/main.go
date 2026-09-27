package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 日記エントリー
type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

// 日記データ保存
func saveDiary(entries []Entry) error {
	data, err := json.MarshalIndent(entries, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("diary.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 日記データ読み込み
func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile("diary.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil
		}
		return nil, err
	}

	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func main() {
	// コマンドライン引数チェック
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	command := os.Args[1]

	switch command {

	case "add":
		// タイトル・本文の有無をチェック
		if len(os.Args) < 4 {
			fmt.Println("使い方：go run main.go add \"タイトル\" \"本文\"")
			return
		}

		title := os.Args[2]
		content := os.Args[3]

		// 既存データ読み込み
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		// 新規エントリー作成
		entry := Entry{
			Timestamp: time.Now(),
			Title:     title,
			Content:   content,
		}

		// 追加
		entries = append(entries, entry)

		// 保存
		err = saveDiary(entries)
		if err != nil {
			fmt.Println("保存エラー:", err)
			return
		}

		fmt.Println("日記を保存しました")

	case "list":
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		if len(entries) == 0 {
			fmt.Println("日記はまだ登録されていません。")
			return
		}

		for _, entry := range entries {
			fmt.Println("日時:", entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Println("タイトル:", entry.Title)
			fmt.Println("本文:", entry.Content)
			fmt.Println("--------------------")
		}

	default:
		fmt.Println("不明なコマンドです:", command)
	}
}
