package check

import (
    "encoding/json"
    "log"
    "os"
    "sensor-online-check/esclient"
    "sensor-online-check/query"
    "sensor-online-check/utils"
    "time"
)

// CheckDeviceStatus 检查单个设备的状态
func CheckDeviceStatus(esClient *esclient.ElasticsearchClient, deviceSN string) (map[string]interface{}, error) {
    // 获取设备查询体
    queryStr := query.GetDeviceQuery(deviceSN)

    // 发送请求
    esResponse, err := esClient.SendRequest(queryStr)
    if err != nil {
        log.Println("Error sending request:", err)
        return nil, err
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
            return nil, err
        }

        // 将时间调整为 UTC 时间，减去 8 小时
        adjustedTime := localTime.Add(-8 * time.Hour)

        // 获取当前时间
        currentTime := utils.GetCurrentTime()

        // 计算时间差并判断是否大于5分钟
        timeDiff := currentTime.Sub(adjustedTime)
        log.Println("时间差为：", timeDiff)

        status := "在线"
        if utils.IsTimeDifferenceGreaterThanFiveMinutes(timeDiff) {
            log.Printf("传感器%s掉线\n", deviceSN)
            status = "掉线"
        } else {
            log.Printf("传感器%s在线\n", deviceSN)
        }

        // 创建日志条目
        logEntry := map[string]interface{}{
            "deviceSN":  deviceSN,
            "status":    status,
            "timestamp": time.Now().Format(time.RFC3339),
        }

        return logEntry, nil
    } else {
        log.Println("没有找到数据")
        return nil, nil
    }
}

// CheckMultipleDeviceStatus 检查多个设备的状态并将结果写入日志文件和 JSON 文件
func CheckMultipleDeviceStatus(logFile *os.File, jsonLogFile *os.File, esClient *esclient.ElasticsearchClient, deviceSNs []string) {
    // 设置日志输出到文件
    log.SetOutput(logFile)

    // 创建一个日志条目切片
    var logEntries []map[string]interface{}

    // 遍历设备列表并检查每个设备的状态
    for _, deviceSN := range deviceSNs {
        logEntry, err := CheckDeviceStatus(esClient, deviceSN)
        if err != nil {
            log.Printf("Error checking device %s: %v", deviceSN, err)
            continue
        }
        if logEntry != nil {
            logEntries = append(logEntries, logEntry)
        }
    }

    // 输出分隔符
    log.Println("#################################################################################")

    // 将日志条目切片转换为 JSON 字符串
    jsonLog, err := json.MarshalIndent(logEntries, "", "  ")
    if err != nil {
        log.Printf("Error marshaling log entries to JSON: %v", err)
        return
    }

    // 将 JSON 字符串写入文件
    if _, err := jsonLogFile.Write(append(jsonLog, '\n')); err != nil {
        log.Printf("Error writing JSON log to file: %v", err)
    }

    // 手动刷新缓冲区
    logFile.Sync()
    jsonLogFile.Sync()
}