package main

import "github.com/webview/webview"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://127.0.0.1:13280/")
	w.Run()
}
