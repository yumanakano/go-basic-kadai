package main

import (
	"fmt"
	"math"
)

// エラー情報を格納する構造体
type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

// errorインターフェースを満たすためのメソッド
func (e ValidationError) Error() string {
	return e.Message
}

// 図形インターフェース
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 長方形
type Rectangle struct {
	Width  float64
	Height float64
}

// 円
type Circle struct {
	Radius float64
}

// 長方形の面積
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 長方形の周長
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 円の面積
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 円の周長
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 図形のバリデーション
func validateShape(shape Shape) error {
	switch s := shape.(type) {
	case Rectangle:
		if s.Width <= 0 {
			return ValidationError{
				Message: "幅は0より大きい値を指定してください",
				Field:   "Width",
				Value:   s.Width,
			}
		}

		if s.Height <= 0 {
			return ValidationError{
				Message: "高さは0より大きい値を指定してください",
				Field:   "Height",
				Value:   s.Height,
			}
		}

	case Circle:
		if s.Radius <= 0 {
			return ValidationError{
				Message: "半径は0より大きい値を指定してください",
				Field:   "Radius",
				Value:   s.Radius,
			}
		}
	}

	return nil
}

func main() {
	// 処理対象の図形スライス
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0},
		Circle{Radius: -1.0},
		Rectangle{Width: 1.0, Height: 0.0},
	}

	// 計算対象の図形を格納するスライス
	var calcShapes []Shape

	// バリデーション
	for _, shape := range shapes {
		err := validateShape(shape)

		if err != nil {
			// ValidationErrorかどうかを判定
			if validationErr, ok := err.(ValidationError); ok {
				fmt.Printf(
					"バリデーションエラー: %s（フィールド: %s, 値: %.1f）\n",
					validationErr.Message,
					validationErr.Field,
					validationErr.Value,
				)
			} else {
				fmt.Printf("その他のエラー: %v\n", err)
			}

			continue
		}

		// バリデーションOK
		calcShapes = append(calcShapes, shape)
	}

	// 計算対象の図形について面積・周長を計算
	fmt.Println("\n--- 計算結果 ---")

	for _, shape := range calcShapes {
		fmt.Printf(
			"型: %T, 面積: %.2f, 周長: %.2f\n",
			shape,
			shape.Area(),
			shape.Perimeter(),
		)
	}
}
