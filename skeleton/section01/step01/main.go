// STEP01: Hello, 世界

// mainパッケージの定義
package main

// main関数から実行される
func main() {
	// 「Hello, 世界」と出力
	println("Hello, 世界")
}

/* go.mod files

事象
go.modで定義しないと以下の通りエラーが発生する。
・go: no module declaration in go.mod. To specify the module path: go mod edit -module=example.com/mod
・gopls was not able to find modules in your workspace. When outside of GOPATH, gopls needs to know which modules you are working on.
　https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory


対処
どこの階層でも良いので、modを定義する必要がある。
https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire

参考
https://zenn.dev/optimisuke/articles/105feac3f8e726830f8c

*/

/* go.work

事象
ビルド不可
go: no modules were found in the current workspace; see 'go help work'

*/
