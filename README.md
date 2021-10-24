# Learn-go-with-test

參考此 github repo(https://github.com/studygolang/learn-go-with-tests) 內容學習 golang testing，以下為學習筆記

## 寫測試的規則
- 測試需在 xxxx_test.go 檔案中轉撰寫
- 測試函式命名需為 Test 開頭
- 測試函式只接受一個 *testing.T 的變數 t

## 寫測試的週期
- 編寫一個測試
- 讓編譯通過
- 執行測試，查看失敗原因並檢查錯誤訊息是有意義的
- 編寫足夠的程式碼使其通過測試
- 重構

## #1 hello
- 可以將固定的字串定義為常數(const)，而不放在函式內，使用常數可以避免每次函式被使用時都創建一次，能幫助快速理含義解甚至提升效能
- 在測試函式中撰寫子測試，子測試用對同一個事情不同的情境做測試，並且可以合併重複的測試碼
- 在測試共用的程式碼中加上 t.Helper()，告訴測試套件這個函式為輔助函式，當測試失敗時所報告的行數會在呼叫的函式當中，而不是輔助函式內部

## #2 integers
- 在被測試的函式上方加上註解，這將會出現在文件中
- 可以添加 ExampleTest([示例](https://go.dev/blog/examples ))更了解測試，並透過`go test -v`執行，只是通常與實際程式碼相比是過時的，就像常忘記更新 Readme 一樣
- 通過指令 godoc -http=:6060 並訪問 http://localhost:6060/pkg/ ，可以在文件中找到對應的 pkg 以及說明

## #3 iteration
- 透過編寫 benchmarks([基準測試](https://pkg.go.dev/testing#hdr-Benchmarks ))，程式碼會將函式運行 b.N 次，並測量需要多長的時間處理，指令為 `go test -bench=.`

## #4 array
- array 的大小也會是類型的一部分，`[4]int` 作為 `[5]int` 類型作為參數傳入函式是不可編譯的
- 測試並不是越多越好，太多的測試會增加維護成本
- 可以透過 `go test -cover` 指令計算測試的覆蓋率([golang 測試覆蓋率文章](https://go.dev/blog/cover))
- 使用 [可變參數](https://gobyexample.com/variadic-functions) 接受或傳遞一個 array 或 slice

## #5 struct
- 將 circle 和 rectangle 賦予 Area() method，以此解決無法使用多型的問題
- 定義一個 Shape interface 並提供 Area() 函式，我們可以只關注形狀面積的產生，而不是長方形或是圓形，達成解耦的作用
- 可以為一系列相同測試方式使用 [表格驅動測試](https://github.com/golang/go/wiki/TableDrivenTests) ，如果要測試新的接口實現，或著對同個結構的數據有不同的情境測試，可以考慮使用

## #6 pointer
- 當提供函式給外部改變內部私有變數時，要使用指針的方式撰寫，才能因內存地址相同修改到同一個物件
  ``` golang
    type Wallet struct {
      balance int
    }
    
    // O
    func (w *Wallet) Deposit(amount int) {
      w.balance += amount
    }

    // X
    func (w Wallet) Deposit(amount int) {
      w.balance += amount
    }
  ```
- 當資料龐大或只想要有一個實例時，也可以使用指針傳遞
- 指針可以為 nil，當函數返回指針時，要檢查是否為 nil，否則會拋出執行異常

## #7 map
- map 有個特別的特性，就是不用指針傳遞就可以修改它們
- map 是引用類型，所以擁有對底層數據結構的引用，所以無論 map 有多大，都只會有一個
- 如果嘗試建立一個 nil 的 map 會導致異常
  ``` golang
    // O
    dictionary = map[string]string{}
    dictionary = make(map[string]string)
  
    // X
    var m map[string]string
  ```
  