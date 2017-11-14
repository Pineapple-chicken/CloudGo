# CloudGo
Assignment for service computing : CloudGo

## 框架
按照老师文档所给说明，使用了gorilla/mux、codegangsta/negroni、unrolled/render。

## 设计思路
基本上是按照老师所给例子进行设计，将一些部分在研究了几个框架后作出了相应的修改。

## 测试
### curl
* 命令：
```shell
curl -v http://localhost:8000/hello/pineapple-chicken
```

* 测试结果：
```shell
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET /hello/pineapple-chicken HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.52.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Thu, 09 Nov 2017 10:18:08 GMT
< Content-Length: 23
< 
{
  "name": "pineapple-chicken"
}
* Curl_http_done: called premature == 0
* Connection #0 to host localhost left intact
```
可看出测试结果正确显示了传入的name。

### ab
* 命令：
```shell
ab -n 100 -c 10 http://localhost:8000/hello/pineapple-chicken
```

* 测试结果：
```shell
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /hello/pineapple-chicken
Document Length:        31 bytes

Concurrency Level:      10
Time taken for tests:   0.072 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      18400 bytes
HTML transferred:       3100 bytes
Requests per second:    1623.54 [#/sec] (mean)
Time per request:       7.132 [ms] (mean)
Time per request:       0.702 [ms] (mean, across all concurrent requests)
Transfer rate:          230.21 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   3.7      1      21
Processing:     0    4   3.8      2      20
Waiting:        0    3   4.6      2      18
Total:          0    6   6.7      4      19

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      8
  75%      8
  80%      9
  90%     20
  95%     20
  98%     20
  99%     20
 100%     20 (longest request)
```
从结果可看出成功完成了测试。
