package check

import (
    "log"
    "sensor-online-check/esclient"
    "sensor-online-check/query"
    "sensor-online-check/utils"
    "time"
    "os"
)

// CheckDeviceStatus 执行设备状态检查并输出状态
func CheckDeviceStatus(esClient *esclient.ElasticsearchClient, deviceSN string) {
    // 打开日志文件，文件不存在时创建，追加写入
    logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    defer logFile.Close()

    // 设置日志输出到文件
    log.SetOutput(logFile)

    // 获取设备查询体
    queryStr := query.GetDeviceQuery(deviceSN)

    // 发送请求
    esResponse, err := esClient.SendRequest(queryStr)
    if err != nil {
        log.Println("Error sending request:", err)
        return
    }

    // 获取 @timestamp 并进行比较
    if len(esResponse.Hits.Hits) > 0 {
        // 假设 timestamp 是以 UTC 时间格式给出的
        timestamp := esResponse.Hits.Hits[0].Source.Timestamp
        log.Println("接收到的时间：", timestamp)

        // 将时间解析为 UTC 时间
        localTime, err := utils.ConvertTimestampToLocalTime(timestamp)
        if err != nil {
            log.Println("时间解析错误：", err)
            return
        }

        // 将时间调整为 UTC 时间，减去 8 小时
        adjustedTime := localTime.Add(-8 * time.Hour)

        // 获取当前时间
        currentTime := utils.GetCurrentTime()

        // 计算时间差并判断是否大于5分钟
        timeDiff := currentTime.Sub(adjustedTime)
        log.Println("时间差为：", timeDiff)

        if utils.IsTimeDifferenceGreaterThanFiveMinutes(timeDiff) {
            log.Println("掉线")
        } else {
            log.Println("在线")
        }
    } else {
        log.Println("没有找到数据")
    }
	log.Println("#################################################################################")
}
