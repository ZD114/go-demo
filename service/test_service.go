package service

import "log"

type testService struct {
}

var TestService = new(testService)

func (s *testService) Test() (string, error) {
	log.Println("测试接口联通性")

	return "测试接口联通性", nil
}
