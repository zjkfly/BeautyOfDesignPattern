package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const idTpl = "%d-%s"

type IdGeneratorService interface {
	generate() (string, error)
}

type IdGenerator struct {
	ctx context.Context
}

func NewIdGenerator(ctx context.Context) IdGeneratorService {
	return &IdGenerator{ctx: ctx}
}

// GetRandomString copy from https://www.csdn.net/tags/OtTaUg4sOTYzOTctYmxvZwO0O0OO0O0O.html
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func (ig *IdGenerator) generate() (string, error) {
	// 简单写
	id := fmt.Sprintf(idTpl, time.Now().Unix(), GetRandomString(6))
	ig.ctx = context.WithValue(ig.ctx, "id", id)
	return id, nil
}

func main() {
	ctx := context.TODO()
	idGen := NewIdGenerator(ctx)
	idGen.generate()
}
