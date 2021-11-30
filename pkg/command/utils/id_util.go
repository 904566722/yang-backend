package utils

import (
    "fmt"
    "math/rand"
    "time"
)

func GenerateId(prefix string, randStrLen int) string {
    str := "0123456789abcdefghijklmnopqrstuvwxyz"
    bytes := []byte(str)
    result := []byte{}
    rand.Seed(time.Now().UnixNano()+ int64(rand.Intn(100)))
    for i := 0; i < randStrLen; i++ {
        result = append(result, bytes[rand.Intn(len(bytes))])
    }
    return fmt.Sprintf("%s-%s", prefix, result)
}

func GenerateTime()  {
    
}