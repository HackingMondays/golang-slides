# Exercices (2)

## Service Web

Code source en Node.js :

~~~ {.nodejs include="src/05-web/server.js" }
~~~

Exécuter avec : `node server.js`


## httptest

05-web/web_test.go

~~~ {.go include="src/05-web/main_test.go" }
~~~


## http

05-web/web.go

~~~ {.go include="src/05-web/main.go.sol" }
~~~


## GO vs Node.js

~~~
GO                                         Node.js
----------------------------------------   ----------------------------
Concurrency Level:     500                   500
Time taken for tests:  0.910 seconds         2.617 seconds
Complete requests:     500                   500
Failed requests:       0                     0
Total transferred:     64000 bytes           56000 bytes
HTML transferred:      5500 bytes            5500 bytes
Requests per second:   549.42 [#/sec]        191.07 [#/sec]
Time per request:      910.052 [ms] (mean)   2616.873 [ms] (mean)
Time per request:      1.820 [ms]            5.234 [ms]
Transfer rate:         68.68 [Kbytes/sec]    20.90 [Kbytes/sec]

Connection Times (ms)                      Connection Times (ms)
              min  mean[+/-sd] med  max      min  mean[+/-sd] med   max
Connect:        0    1  22.3     0  498        0    1  23.6     0   528
Processing:   304  491 151.4   556  848      559 1460 290.5  1302  2041
Waiting:      303  485 148.6   539  847      559 1460 290.6  1302  2041
Total:        304  493 152.0   556  849      559 1461 291.7  1302  2041

Percentage of the requests served          Percentage of the requests
  50%    556                                 50%   1302
  66%    561                                 66%   1662
  75%    564                                 75%   1719
  80%    569                                 80%   1771
  90%    585                                 90%   1823
  95%    831                                 95%   1827
  98%    836                                 98%   1831
  99%    838                                 99%   2026
 100%    849 (longest request)              100%   2041 (longest)
~~~

`ab -c 500 -n 500 http://localhost:8080/`

## Grep en GO

on a :

~~~
$ grep --regexp=^root passwd
root:*:0:0:System Administrator:/var/root:/bin/sh
~~~

on veut :

~~~
$ go run main.go --regexp=^root passwd
root:*:0:0:System Administrator:/var/root:/bin/sh
~~~

## flag

06-io/main.go

~~~ {.go include="src/06-io/flag.src" }
~~~

## os / bufio

06-io/main.go

~~~ {.go include="src/06-io/fileopen.src" }
~~~
