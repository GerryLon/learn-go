package talker

import (
	"testing"

	"github.com/GerryLon/learn-go/gomock/talker/mock"

	"github.com/golang/mock/gomock"
)

// 传统测试
func TestSayHello(t *testing.T) {
	talker := NewPerson("王尼美")
	company := NewCompany(talker)
	t.Log(company.Meeting("王尼玛"))
}

// 使用Mock对象测试
func TestSayHello2(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTalker := mock.NewMockTalker(ctrl)

	mockTalker.EXPECT().SayHello(gomock.Eq("王尼美")).Return("王尼玛")

	company := NewCompany(mockTalker)
	t.Log(company.Meeting("王尼美"))
}
