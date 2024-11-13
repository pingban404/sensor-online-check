package main

import (
    "log"
    "os"
    "sensor-online-check/config"
    "sensor-online-check/esclient"
    "sensor-online-check/check"
    "time"
)

func main() {
    // 打开日志文件，文件不存在时创建，追加写入
    logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }
    defer logFile.Close()

    // 打开 JSON 日志文件，文件不存在时创建，追加写入
    jsonLogFile, err := os.OpenFile("app.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Error opening JSON log file: %v", err)
    }
    defer jsonLogFile.Close()

    // 加载配置文件
    conf, err := config.LoadConfig("config.yml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // 配置 Elasticsearch 请求和认证信息
    esClient := esclient.NewElasticsearchClient(conf.Elasticsearch.URL, conf.Elasticsearch.AuthToken)
    
    // 配置设备SN
    deviceSNs := []string{"2407101", "2407104", "2407106"}

    // // 立即执行
    // check.CheckMultipleDeviceStatus(logFile, jsonLogFile, esClient, deviceSNs)
    
    // 定义 ticker，每隔 30 分钟运行一次
    ticker := time.NewTicker(30 * time.Minute)
    defer ticker.Stop()

    for {
        // 执行设备状态检查
        check.CheckMultipleDeviceStatus(logFile, jsonLogFile, esClient, deviceSNs)
        // 等待 30 分钟
        <-ticker.C
    }
}