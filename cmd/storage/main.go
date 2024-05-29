package main

import (
	"fmt"

	"github.com/Bostanova/route256_tgbot/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println(st)
}
