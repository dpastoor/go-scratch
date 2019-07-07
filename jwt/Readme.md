```
curl -i -X POST \
-H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI1MTQ5NzIsImlhdCI6MTU2MjUxNDM3MiwiaXNzIjoiMSJ9.WpxcYowv_53Mhz3JTYNw3kR0HBK939tWgDDFI4dotLW8Gr0pYWX4Xjwcx0QuV-QJr9BO4k8Okkf_ch_Dc4hwih-GQoglJ1i42gL4ziHuB5Dvxyw6dy9dD7KX74yZYOKFmj3QXE5XjJ8ydtgoyey_wMossePCnz9i0LN9NGn86uga8nXL3VIZopCDRe_YsYBGs71uvqkl2k8McT_mStu2v89Wbup8bJolZT_LImTXA4hHuB8sD-Kzt59f0FnPrZztcL7XyagDDedltyY7b1A-7SoNLUCpYDJu_Y88G3FjksdgrV-A4hx3tJIrFjK3qTFzL5no4zv6yrlhLSZ8S-c37Q" \
-H "Accept: application/vnd.github.machine-man-preview+json" \
https://ghe.metrumrg.dev/api/v3/app/installations/:installation_id/access_tokens
```

```
HTTP/1.1 201 Created
Server: GitHub.com
Date: Sun, 07 Jul 2019 15:46:58 GMT
Content-Type: application/json; charset=utf-8
Content-Length: 101
Status: 201 Created
Cache-Control: public, max-age=60, s-maxage=60
Vary: Accept
ETag: "59e2b0475a56c7efa02c44a3239c401f"
X-GitHub-Enterprise-Version: 2.17.3
X-GitHub-Media-Type: github.machine-man-preview; format=json
Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
Access-Control-Allow-Origin: *
X-GitHub-Request-Id: d90e8ced-ca4c-4d1d-b31d-e11405fe8a01
Strict-Transport-Security: max-age=31536000; includeSubdomains
X-Frame-Options: deny
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
Content-Security-Policy: default-src 'none'
X-Runtime-rack: 0.379679

{
  "token": "v1.0c3f8cc48217422434556c19e4a8524bf226f92c",
  "expires_at": "2019-07-07T16:46:58Z"
}
```

This token can be used


```
git clone https://x-access-token:v1.0c3f8cc48217422434556c19e4a8524bf226f92c@ghe.metrumrg.dev/mrg/IT_Help
```