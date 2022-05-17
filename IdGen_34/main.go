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

func (ig *IdGenerator) generate() (string, error) {
	// 简单写
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < 6; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	id := fmt.Sprintf(idTpl, time.Now().Unix(), result)
	ig.ctx = context.WithValue(ig.ctx, "id", id)
	return id, nil
}

func main() {
	ctx := context.TODO()
	idGen := NewIdGenerator(ctx)
	idGen.generate()
}
