package query

// GetDeviceQuery 返回查询 Elasticsearch 的查询体
func GetDeviceQuery(deviceSN string) string {
    return `{
        "query": {
            "bool": {
                "must": [
                    {
                        "term": {
                            "device_sn": "` + deviceSN + `"
                        }
                    }
                ]
            }
        },
        "sort": [
            {
                "@timestamp": {
                    "order": "desc"
                }
            }
        ]
    }`
}
