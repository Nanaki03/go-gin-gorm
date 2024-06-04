package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "go_gin_gorm/libraries"
    "fmt"
)

type User struct{
    Id int
    UserId string
    Password string
}

type JsonRequest struct{
    UserId string 'json:"user_id"'
    Password string 'json:"password"'
}

func main(){
    engine := gin.Default()

    dsn := "seed:Tech_123@tcp(127.0.0.1:3306)/results?charset=utf8mb4&parseTime=True&loc=Local"
    db,err:=gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("DB接続失敗");
        return
    }

    // Sign up用ルーティング
    engine.POST("/signin", func(c *gin.Context) {
        var json JsonRequest
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error" : err.Error(),
            })
            return
        }
        userId := json.UserId
        pw := json.Password

        user := User{}
        db.Where("user_id = ?", userId).First(&user)
        if user.Id == 0 {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message" : "ユーザーが存在しません。",
            })
            return
        }

        err := crypto.CompareHashAndPassword(user.Password, pw)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message" : "パスワードが一致しません。",
            })
            return
        }
        
        c.JSON(http.StatusOK, gin.H{
            "message" : "ログイン成功",
        })
    })

        // 同一ユーザIDの検証
        user := User{}
        db.Where("user_id = ?", userId).First(&user)
        if user.Id != 0 {
            c.JSON(http.StatusBadRequest, gin.H{
                "message" : "そのUserIdは既に登録されています。",
            })
            return
        }

        // パスワードの暗号化
        encryptPw, err := crypto.PasswordEncrypt(pw)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "message" : "パスワードの暗号化でエラーが発生しました。",
            })
            return
        }
        
        // DBへの登録
        user = User{UserId: userId, Password: encryptPw}
        db.Create(&user)
        
        c.JSON(http.StatusOK, gin.H{
            "message" : "アカウント登録完了",
        })

        engine.Run(":8080")
    }
   
