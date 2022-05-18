package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	idTpl = "%s-%d-%s"

	ctxKey = "id"
)

type IdGenerator interface {
	Generate() (string, error)
}
type CtxTraceIdGenerator interface {
	IdGenerator
}

type RandomIdGenerator struct {
	ctx context.Context
}

/*
	Id 生成器是用来生成随机不重复非递增的id信息的

	该生成器存在一定的可能性相同，但概率极低。放心食用
*/
func NewRandomIdGenerator(ctx context.Context) CtxTraceIdGenerator {
	return &RandomIdGenerator{ctx: ctx}
}

// GetRandomString wants n>=0, return "" if illegal
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

func (ig *RandomIdGenerator) Generate() (string, error) {
	// 简单写
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	if len(hostname) == 0{
		return "", errors.New("hostname is empty")
	}

	id := fmt.Sprintf(idTpl, hostname, time.Now().Unix(), GetRandomString(6))
	ig.ctx = context.WithValue(ig.ctx, ctxKey, id)
	return id, nil
}

func main() {
	ctx := context.TODO()
	idGen := NewRandomIdGenerator(ctx)
	idGen.Generate()
}
