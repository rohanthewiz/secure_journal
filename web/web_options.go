package web

import (
	"fmt"
	"strings"
)

func menuItem(item string) string {
	arr := strings.Split(item, " ")
	if len(arr) > 0 {
		item = arr[0]
	}

	return fmt.Sprintf(`<li><a href="/%s">%s</a></li>`, strings.ToLower(item), item)
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
