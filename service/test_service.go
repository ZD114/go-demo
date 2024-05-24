package service

import (
	"fmt"
	"zhangda.com/go-demo/config"
	"zhangda.com/go-demo/repository"
)

type testService struct {
}

var TestService = new(testService)

func (s *testService) Test(id int64) (*repository.Test, error) {
	ms := new(repository.Test)

	if _, err := config.GetDB().SQL("SELECT * FROM test WHERE id = ? ", id).Get(ms); err != nil {
		fmt.Println(err)
		return ms, err
	}

	return ms, nil
}
