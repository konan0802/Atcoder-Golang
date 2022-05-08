package main

// ーーーーーーーーーー入力ーーーーーーーーーー

// 標準入力
sc := bufio.NewScanner(os.Stdin)

// ファイル入力
file, _ := os.Open("test.txt")
sc := bufio.NewScanner(file)

// 1行ずつ
func input(sc *bufio.Scanner) string {
	sc.Scan()
	x := sc.Text()
	return N
}

// スペース区切りを配列に変換
arr := strings.Split(sc.Text(), " ")

// 複数行
for sc.Scan() {

}

// ーーーーーーーーーー出力ーーーーーーーーーー
fmt.Println(str)

// ーーーーーーーーーー終了ーーーーーーーーーー
os.Exit(0)

// ーーーーーーーーーーキャストーーーーーーーーーー

// string -> int
x, _ := strconv.Atoi(str)

// ーーーーーーーーーー三項演算子ーーーーーーーーーー


// ーーーーーーーーーー配列ーーーーーーーーーー

// 配列のstring -> int変換
A := []int{}
for _, v := range SA {
	iv, _ := strconv.Atoi(v)
	A = append(A, iv)
}

// 昇順ソート
sort.Ints(s)

// 降順ソート
sort.Sort(sort.Reverse(sort.IntSlice(A)))

// ーーーーーーーーーーライブラリーーーーーーーーーー


// ーーーーーーーーーー自作関数ーーーーーーーーーー

// ーーーーーーーーーーTIPSーーーーーーーーーー
// 配列を全探索する処理は減らす
// ⇒ inでの配列に存在するかを確認
// ⇒ sortでの配列の順序の入れ替え