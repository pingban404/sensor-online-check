package main

import (
    "log"
    "sensor-online-check/config"
    "sensor-online-check/esclient"
    "sensor-online-check/check"
    "time"
)

func main() {
    // 加载配置文件
    conf, err := config.LoadConfig("config.yml")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // 配置 Elasticsearch 请求和认证信息
    esClient := esclient.NewElasticsearchClient(conf.Elasticsearch.URL, conf.Elasticsearch.AuthToken)

    // 定义 ticker，每隔 30 分钟运行一次
    ticker := time.NewTicker(30 * time.Minute)
    defer ticker.Stop()

    for {
        // 执行设备状态检查
        check.CheckDeviceStatus(esClient, "2407104")

        // 等待 30 分钟
        <-ticker.C
    }
}
