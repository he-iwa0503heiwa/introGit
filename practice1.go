/*実行するときはターミナルを開き、自分のディレクトリにいき
pwd 今のディレクトリがわかる
cd Desktop/　とかで移動

go build practice1.go　これでコンパイルされファイルができる
./Practice1				ファイルを実行
もしくは
go run practice1.go　でコンパイルと実行までできる
*/
package main
import (
	"fmt"
)

func main() {
	// var num int 
	// num = 1 変数宣言して使わなかったら「不要だから消せ」とコンパイルエラーになる
	num := 1 
	num += 10

	index := []string{"iti","ni","san"}
	fmt.Println(index[0])
	index[0] = "tanaka"
	fmt.Println(index[0])

	fmt.Println(num)

	num++

	fmt.Println(num)

	if kekka := 7; kekka >= 5 && kekka < 10{
		fmt.Println("syuturyoku")
	}
}