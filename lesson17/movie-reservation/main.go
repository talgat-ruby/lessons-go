package main

import (
	"context"
	"log/slog"
)

//var SomeError = fmt.Errorf("some error")

func main() {
	ctx := context.Background()

	//fmt.Println(ParentTask())
	//fmt.Println(SomeError == ParentTask())
	//fmt.Println(errors.Is(SomeError, ParentTask()))

	d, err := newDB(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}
	if err := d.Init(ctx); err != nil {
		panic(err)
	}

	a := newApi(slog.With("service", "api"))
	if err := a.Start(ctx); err != nil {
		panic(err)
	}
}

//func ParentTask() error {
//	if err := SomeTask(); err != nil {
//		return fmt.Errorf("SomeTask: %w", err)
//	}
//
//	return nil
//}
//
//func SomeTask() error {
//	return SomeError
//}
