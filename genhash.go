package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    var password string

    fmt.Print("パスワードを入力してください: ")
    fmt.Scanln(&password)

    // ハッシュ値を生成
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    fmt.Printf("ハッシュ値: %s\n", string(hashedPassword))
}
