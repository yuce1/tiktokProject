
# base api

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestFeed|TestUserAction|TestPublish)$ tiktok-go/test -v -count=1

=== RUN   TestFeed
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 8.972044ms
        Content-Length: 384
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"video_list":[{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676986858}
--- PASS: TestFeed (0.01s)
=== RUN   TestUserAction
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyin47354&username=douyin47354 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyin47354&username=douyin47354
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 2.699258ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":2,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoyLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.3zqSoKgIBRDUNcXl8YaSjgmGjdB8CS8pjnIOb_3HP4I"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyin47354&username=douyin47354 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyin47354&username=douyin47354
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.405326ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":2,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjozLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.Px3odJDUuezxYg7cUlJrCX7W25-2m6PrnU01OXQQw2s"}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/user/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjozLCJ1c2VyIjoyLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.Px3odJDUuezxYg7cUlJrCX7W25-2m6PrnU01OXQQw2s HTTP/1.1
        Host: localhost:8080

    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 3.540688ms
        Content-Length: 54
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user":{"id":2,"name":"douyin47354"}}
--- PASS: TestUserAction (0.01s)
=== RUN   TestPublish
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 3.646741ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:58 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/publish/action/ HTTP/1.1
        Host: localhost:8080
        Content-Type: multipart/form-data; boundary=6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9

        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9
        Content-Disposition: form-data; name="data"; filename="../public/bear.mp4"
        Content-Type: application/octet-stream

        [file data content, skip]

        Content-Disposition: form-data; name="token"

        eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ
        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9
        Content-Disposition: form-data; name="title"

        Bear
        --6f3a1ef51cb3e1e0989a5e7bf0050155a90c91577a7c20ce46d0415293f9--
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 220.060574ms
        Content-Length: 125
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:59 GMT

        {"status_code":0,"status_msg":"3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4 uploaded successfully"}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/publish/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo0LCJ1c2VyIjozLCJleHAiOjE2NzcwNzMyNTgsImlhdCI6MTY3Njk4Njg1OH0.BjPjXohAGE1fRHTG_EDpXH-jyMMqVy-tERIAzXXV3OQ&user_id=3 HTTP/1.1
        Host: localhost:8080

    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 2.656585ms
        Content-Length: 371
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:40:59 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","is_favorite":false,"title":"Bear"}]}
--- PASS: TestPublish (0.35s)
PASS
ok  	tiktok-go/test	0.372s

```

# interact_api

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestFavorite|TestComment)$ tiktok-go/test -v -count=1

=== RUN   TestFavorite
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 5.531507ms
        Content-Length: 723
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","is_favorite":false,"title":"Bear"},{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676987088}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 849.755µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.518259ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/favorite/action/?action_type=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&video_id=2
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 11.475292ms
        Content-Length: 40
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"status_msg":"Success"}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/favorite/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&user_id=3 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo1LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.RIQwDRo_Q11u-Av0vH6GTDCCou6qBHzVZgM7gCJjYYY&user_id=3
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.449418ms
        Content-Length: 409
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","favorite_count":1,"is_favorite":false,"title":"Bear"}]}
--- PASS: TestFavorite (0.02s)
=== RUN   TestComment
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/feed/ HTTP/1.1
        Host: localhost:8080

    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.21459ms
        Content-Length: 761
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"video_list":[{"id":2,"author":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"play_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.mp4","cover_url":"http://47.113.179.123:8080/static/3_3472a10a299cfd24a062a2f9415e57e3fd3fad14df59e9182571fef592e4eb5f.jpg","favorite_count":1,"is_favorite":false,"title":"Bear"},{"id":1,"author":{"id":1,"name":"test","work_count":1},"play_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.mp4","cover_url":"http://47.113.179.123:8080/static/1_003198fdd79363536c7c6db516d174df06ed9993879b7eb2234269c51444f833.jpg","is_favorite":false,"title":"heusb"}],"next_time":1676987088}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 706.034µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 777.666µs
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/comment/action/?action_type=1&comment_text=%E6%B5%8B%E8%AF%95%E8%AF%84%E8%AE%BA&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&comment_text=%E6%B5%8B%E8%AF%95%E8%AF%84%E8%AE%BA&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 71.941649ms
        Content-Length: 208
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"status_msg":"publish successfully","comment":{"id":1,"user":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"content":"测试评论","create_date":"2023-02-21 13:44:48"}}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/comment/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.60325ms
        Content-Length: 179
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0,"comment_list":[{"id":1,"user":{"id":3,"name":"douyinTestUserA","work_count":1,"favorite_count":1},"content":"测试评论","create_date":"2023-02-21 13:44:48"}]}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/comment/action/?action_type=2&comment_id=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=2&comment_id=1&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo2LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM0ODgsImlhdCI6MTY3Njk4NzA4OH0.yzpNYC7smLGSuBB4UgybhDgU08dVmnIoboY931vlfX0&video_id=2
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 3.21199ms
        Content-Length: 17
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:44:48 GMT

        {"status_code":0}
--- PASS: TestComment (0.08s)
PASS
ok  	tiktok-go/test	0.115s
```

# social api

```
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestRelation|TestChat)$ tiktok-go/test -v -count=1

=== RUN   TestRelation
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 3.298312ms
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.225268ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo3LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.9B9TX6A1WSRhyTa3eG-BkQuonUGwnFxFEhN0tAdPcwY"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserB&username=douyinTestUserB HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserB&username=douyinTestUserB
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 2.558839ms
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_id":4,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo4LCJ1c2VyIjo0LCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.pzkz1GIi2eKHgbURTw_pnvI5EBfilkicfCwHYDR9XI4"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/relation/action/?action_type=1&to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo3LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.9B9TX6A1WSRhyTa3eG-BkQuonUGwnFxFEhN0tAdPcwY HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo3LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.9B9TX6A1WSRhyTa3eG-BkQuonUGwnFxFEhN0tAdPcwY
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 5.590203ms
        Content-Length: 17
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/relation/follow/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo3LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.9B9TX6A1WSRhyTa3eG-BkQuonUGwnFxFEhN0tAdPcwY&user_id=3 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo3LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.9B9TX6A1WSRhyTa3eG-BkQuonUGwnFxFEhN0tAdPcwY&user_id=3
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.121876ms
        Content-Length: 101
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_list":[{"id":4,"name":"douyinTestUserB","follower_count":1,"is_follow":true}]}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/relation/follower/list/?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo4LCJ1c2VyIjo0LCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.pzkz1GIi2eKHgbURTw_pnvI5EBfilkicfCwHYDR9XI4&user_id=4 HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo4LCJ1c2VyIjo0LCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.pzkz1GIi2eKHgbURTw_pnvI5EBfilkicfCwHYDR9XI4&user_id=4
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 1.105995ms
        Content-Length: 116
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_list":[{"id":3,"name":"douyinTestUserA","follow_count":1,"work_count":1,"favorite_count":1}]}
--- PASS: TestRelation (0.02s)
=== RUN   TestChat
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 450.4µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserA&username=douyinTestUserA HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserA&username=douyinTestUserA
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 710.527µs
        Content-Length: 196
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_id":3,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo5LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.sDJLrU_L7syumGc2FSLBnIefdOH7_u-w3v0c83J9J-w"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/register/?password=douyinTestUserB&username=douyinTestUserB HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserB&username=douyinTestUserB
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 522.38µs
        Content-Length: 62
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":1,"status_msg":"User already exist","token":""}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/user/login/?password=douyinTestUserB&username=douyinTestUserB HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        password=douyinTestUserB&username=douyinTestUserB
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 761.387µs
        Content-Length: 197
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"user_id":4,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoxMCwidXNlciI6NCwiZXhwIjoxNjc3MDczNTIzLCJpYXQiOjE2NzY5ODcxMjN9.kJ6v_xV6SjJcQGJotLHq-yvq8alv3QbibZwho841-Lg"}
    /home/go/tiktokProject/test/printer.go:116: POST /douyin/message/action/?action_type=1&content=Send+to+UserB&content=Send+to+UserB&to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo5LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.sDJLrU_L7syumGc2FSLBnIefdOH7_u-w3v0c83J9J-w HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        action_type=1&to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo5LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.sDJLrU_L7syumGc2FSLBnIefdOH7_u-w3v0c83J9J-w
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 2.086667ms
        Content-Length: 17
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/message/chat/?to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo5LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.sDJLrU_L7syumGc2FSLBnIefdOH7_u-w3v0c83J9J-w HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        to_user_id=4&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjo5LCJ1c2VyIjozLCJleHAiOjE2NzcwNzM1MjMsImlhdCI6MTY3Njk4NzEyM30.sDJLrU_L7syumGc2FSLBnIefdOH7_u-w3v0c83J9J-w
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 873.767µs
        Content-Length: 126
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"message_list":[{"id":1,"from_user_id":3,"to_user_id":4,"content":"Send to UserB","create_time":1676987123}]}
    /home/go/tiktokProject/test/printer.go:116: GET /douyin/message/chat/?to_user_id=3&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoxMCwidXNlciI6NCwiZXhwIjoxNjc3MDczNTIzLCJpYXQiOjE2NzY5ODcxMjN9.kJ6v_xV6SjJcQGJotLHq-yvq8alv3QbibZwho841-Lg HTTP/1.1
        Host: localhost:8080
        Content-Type: application/x-www-form-urlencoded

        to_user_id=3&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoxMCwidXNlciI6NCwiZXhwIjoxNjc3MDczNTIzLCJpYXQiOjE2NzY5ODcxMjN9.kJ6v_xV6SjJcQGJotLHq-yvq8alv3QbibZwho841-Lg
    /home/go/tiktokProject/test/printer.go:133: HTTP/1.1 200 OK 598.724µs
        Content-Length: 126
        Content-Type: application/json; charset=utf-8
        Date: Tue, 21 Feb 2023 13:45:23 GMT

        {"status_code":0,"message_list":[{"id":1,"from_user_id":3,"to_user_id":4,"content":"Send to UserB","create_time":1676987123}]}
--- PASS: TestChat (0.01s)
PASS
ok  	tiktok-go/test	0.030s
```

# ALL PASS