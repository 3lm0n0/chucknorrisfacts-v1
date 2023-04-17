package main

import (
	"context"
	"fmt"
	"time"
)

type Logger struct {
	svc Service
}

func NewLogger(svc Service) Service {
	return &Logger{svc: svc}
}

func (l *Logger) GetFact(ctx context.Context) (fact *Fact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v \n", fact.Fact, err, time.Since(start))
	}(time.Now())
	return l.svc.GetFact(ctx)
}
