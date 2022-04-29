package main

import (
	"math"
	"sort"
	"time"
)

type MetricsCollector struct {
	metricsRepo MetricsRepo
}

func NewMetricsCollector(metricsRepo MetricsRepo) *MetricsCollector {
	return &MetricsCollector{metricsRepo: metricsRepo}
}

func (mc *MetricsCollector) RecordRequest(requestInfo *RequestInfo) {
	if requestInfo == nil || requestInfo.apiName == "" {
		return
	}
	mc.metricsRepo.SaveRequestInfo(requestInfo)
}

func (mc *MetricsCollector) GetRequestInfos(apiName string, startTs, endTs int64) []RequestInfo {

	return mc.metricsRepo.GetRequestInfos(apiName, startTs, endTs)
}

type RequestInfo struct {
	apiName      string
	responseTime float64
	timestamps   int64
}

type MetricsRepo interface {
	SaveRequestInfo(requestInfo *RequestInfo)
	GetRequestInfos(apiName string, startTimeInMillis int64, endTimeInMillis int64) []RequestInfo
	GetAllRequestInfos(startTimeInMillis int64, endTimeInMillis int64) map[string][]RequestInfo
}

type RedisMetricsRepository struct {
}

func NewRedisMetricsRepository() MetricsRepo {
	return &RedisMetricsRepository{}
}

type RequestStat struct {
	MaxResponseTime  float64
	MinResponseTime  float64
	AvgResponseTime  float64
	P999ResponseTime float64
	P99ResponseTime  float64
	Count            int64
	Tps              int64
}

// copy from https://github.com/dgqypl/BeautyOfDesignPatterns/blob/main/lesson25and26/metrics.go#L57
// and do some chores
func Aggregate(requestInfos []RequestInfo, durationInMillis int64) RequestStat {
	maxRespTime := math.SmallestNonzeroFloat64
	minRespTime := math.MaxFloat64
	avgRespTime := -1.0
	p999RespTime := -1.0
	p99RespTime := -1.0
	sumRespTime := 0.0
	var count int64 = 0
	for _, requestInfo := range requestInfos {
		count++
		respTime := requestInfo.responseTime
		if maxRespTime < respTime {
			maxRespTime = respTime
		}
		if minRespTime > respTime {
			minRespTime = respTime
		}
		sumRespTime += respTime
	}
	if count != 0 {
		avgRespTime = sumRespTime / float64(count)
	}
	tps := count / durationInMillis * 1000
	sort.SliceStable(requestInfos, func(i, j int) bool {
		return requestInfos[i].responseTime-requestInfos[j].responseTime < 0.0
	})
	idx999 := int(float64(count) * 0.999)
	idx99 := int(float64(count) * 0.99)
	if count != 0 {
		p999RespTime = requestInfos[idx999].responseTime
		p99RespTime = requestInfos[idx99].responseTime
	}
	return RequestStat{
		MaxResponseTime:  maxRespTime,
		MinResponseTime:  minRespTime,
		AvgResponseTime:  avgRespTime,
		P999ResponseTime: p999RespTime,
		P99ResponseTime:  p99RespTime,
		Count:            count,
		Tps:              tps,
	}
}

func (r *RedisMetricsRepository) SaveRequestInfo(requestInfo *RequestInfo) {
	panic("implement me")
}

func (r *RedisMetricsRepository) GetRequestInfos(apiName string, startTimeInMillis int64, endTimeInMillis int64) []RequestInfo {
	panic("implement me")
}
func (r *RedisMetricsRepository) GetAllRequestInfos(startTimeInMillis int64, endTimeInMillis int64) map[string][]RequestInfo {
	panic("implement me")
}

type ConsoleReporter struct {
	metricsRepo MetricsRepo
}

func NewConsoleReporter(metricsRepo MetricsRepo) *ConsoleReporter {
	return &ConsoleReporter{metricsRepo: metricsRepo}
}

func (cr *ConsoleReporter) StartRepeatedReport(periodInSeconds int64, durationInSeconds int64) {
	// todo 输出代码
}

func main() {
	repo := NewRedisMetricsRepository()
	collector := NewMetricsCollector(repo)
	collector.RecordRequest(&RequestInfo{
		apiName:      "register",
		responseTime: float64(time.Now().Unix()),
		timestamps:   time.Now().Unix(),
	})

	reporter := NewConsoleReporter(repo)
	reporter.StartRepeatedReport(0, 60)
}
