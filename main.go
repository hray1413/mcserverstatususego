package main

import (
    "encoding/json"
    "fmt"
    "os"

    "github.com/Tnze/go-mc/net"
    "github.com/Tnze/go-mc/status"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("用法：go run main.go <伺服器 IP>")
        return
    }

    server := os.Args[1]
    conn, err := net.DialMC(server)
    if err != nil {
        fmt.Printf("❌ 無法連線到伺服器 %s：%v\n", server, err)
        return
    }
    defer conn.Close()

    resp, err := status.HandshakeAndStatus(conn, server, 25565)
    if err != nil {
        fmt.Printf("❌ 查詢失敗：%v\n", err)
        return
    }

    var data map[string]interface{}
    if err := json.Unmarshal(resp, &data); err != nil {
        fmt.Printf("❌ JSON 解析失敗：%v\n", err)
        return
    }

    fmt.Println("✅ 伺服器狀態：")
    fmt.Printf("📝 MOTD：%v\n", data["description"])
    fmt.Printf("👥 玩家：%v\n", data["players"])
    fmt.Printf("🎮 版本：%v\n", data["version"])
}
