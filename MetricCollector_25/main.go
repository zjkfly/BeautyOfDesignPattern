package main

import (
	"time"
)

type Metrics struct {
	responseTimes map[string][]int64
	timestamps    map[string][]int64
}

func NewMetrics(responseTimes map[string][]int64, timestamps map[string][]int64) MetricsInterface {
	return &Metrics{responseTimes: responseTimes, timestamps: timestamps}
}

type MetricsInterface interface {
	recordResponseTime(apiName string, responseTs int64) error
	recordTimestamp(apiName string, responseTs int64) error
	startRepeatedReport(apiName string, period int64) error
	max(apiName string, startTs, endTs int64) (int64, error)
	min(apiName string, startTs, endTs int64) (int64, error)
}

func (m *Metrics) recordResponseTime(apiName string, responseTs int64) error {
	if m.responseTimes == nil {
		m.responseTimes = map[string][]int64{}
	}
	m.responseTimes[apiName] = append(m.responseTimes[apiName], responseTs)
	return nil
}
func (m *Metrics) recordTimestamp(apiName string, timestamp int64) error {
	if m.timestamps == nil {
		m.timestamps = map[string][]int64{}
	}
	m.timestamps[apiName] = append(m.timestamps[apiName], timestamp)
	return nil
}

func (m *Metrics) startRepeatedReport(apiName string, period int64) error {
	panic("implement me")
}

func (m *Metrics) max(apiName string, startTs, endTs int64) (int64, error) {
	panic("implement me")
}

func (m *Metrics) min(apiName string, startTs, endTs int64) (int64, error) {
	panic("implement me")
}

type User struct {
	account  string
	password string
	sex      int
	addr     string
}

func NewUser(account string, password string, sex int, addr string) *User {
	return &User{account: account, password: password, sex: sex, addr: addr}
}

type UserInterface interface {
	register(user User) error
	login(account, password string) *User
}
type UserController struct {
	metricsInterface MetricsInterface
}

func NewUserController(metricsInterface MetricsInterface) UserInterface {
	return &UserController{
		metricsInterface: metricsInterface,
	}
}

func (u *UserController) register(user User) error {
	apiName := "register"
	u.metricsInterface.recordTimestamp(apiName, time.Now().Unix())
	/*
		todo 业务操作
	*/
	u.metricsInterface.recordResponseTime(apiName, time.Now().Unix())

	return nil
}
func (u *UserController) login(account, password string) *User {
	apiName := "login"
	u.metricsInterface.recordTimestamp(apiName, time.Now().Unix())
	/*
		todo 业务操作
	*/
	u.metricsInterface.recordResponseTime(apiName, time.Now().Unix())
	return nil
}

func main() {
	userController := NewUserController(NewMetrics(map[string][]int64{}, map[string][]int64{}))
	_ = userController.register(*NewUser("zjk", "123456", 1 /*1 man 2 woman 3 unknown*/, "白京"))
	userController.login("zjk", "123456")
}
