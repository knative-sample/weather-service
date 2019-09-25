# weather-service
- 提供根据城市、日期，查询天气预报信息

request:
```
http://127.0.0.1:8080/api/weather/query?cityCode=110000&date=2019-09-26
```
response:
```
{
    "code":200,
    "message":"",
    "data":{
        "adcode":"110000",
        "city":"北京市",
        "date":"2019-09-26",
        "daypower":"≤3",
        "daytemp":"30",
        "dayweather":"晴",
        "daywind":"东南",
        "nightpower":"≤3",
        "nighttemp":"15",
        "nightweather":"晴",
        "nightwind":"东南",
        "province":"北京",
        "reporttime":"2019-09-25 14:50:46",
        "week":"4"
    }
}
```
