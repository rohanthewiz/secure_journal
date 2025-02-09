package web

import "fmt"

func menuItem(item string) string {
	return fmt.Sprintf("<li>%s</li>", item)
}

// key + password = hash
func WebOptions() (menu string) {
	menu = "===== Menu =====<br>"
	menu += menuItem("Register")
	menu += menuItem("Login")
	menu += menuItem("Delete me")
	menu += menuItem("Exit")
	menu += "Enter your choice"
	return
}
