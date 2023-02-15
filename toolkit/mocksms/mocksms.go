package mocksms

import (
	"sync"
	"time"
	"toolkit/random"
)

// StatusOpts 验证码状态选项
var StatusOpts status

// 枚举值类型
type statusType int

// 定义枚举
type status struct {
	Normal  statusType
	Expired statusType
	InValid statusType
}

// 初始化枚举
func init() {
	StatusOpts = status{
		Normal:  0,
		InValid: 1,
		Expired: 2,
	}
}

// 存储验证码
var mockMaps = make(map[string]*mockSMS, 10)

// mockSMS 定义验证码数据结构
type mockSMS struct {
	code   int64
	status statusType
}

// Code 返回验证码
func (ms *mockSMS) Code() int64 {
	return ms.code
}

// Status 返回状态码
func (ms *mockSMS) Status() statusType {
	return ms.status
}

// SetStatus 设置状态
func (ms *mockSMS) SetStatus(phone string, status statusType) {
	mu.Lock()
	defer mu.Unlock()
	ss, ok := mockMaps[phone]
	if ok {
		ss.status = status
	}
}

var mu = sync.RWMutex{}

func NewMockSMS(phone string) *mockSMS {
	code := random.RandomInt(1000, 9999)
	ss := mockSMS{
		code: code,
	}

	mu.Lock()
	mockMaps[phone] = &ss
	mu.Unlock()

	go func() {
		// 5 分钟后过期
		time.Sleep(5 * time.Minute)
		ss.SetStatus(phone, StatusOpts.Expired)

		// 10 分钟后删除
		time.Sleep(10 * time.Minute)
		mu.Lock()
		delete(mockMaps, phone)
		mu.Unlock()
	}()

	return &ss
}

// GetMockSMS 获取一个验证码
func GetMockSMS(phone string) (sms *mockSMS, ok bool) {
	sms, ok = mockMaps[phone]
	return
}
