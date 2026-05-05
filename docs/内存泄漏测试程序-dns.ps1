



go run ./cmd/ -enable-pprof  '-dohurl' 'https://61919494499.security.cloudflare-dns.com/dns-query' '-dohip' '104.21.14.41' '--dohalpn=h2' '-port' '28340' '-username' 'admin' '-password' 'iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58' '-dohalpn=h3' '-dohurl' 'https://61919494499.security.cloudflare-dns.com/dns-query' '--dohip=104.21.9.230' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h3' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h3' -cache-file ./data/dns_cache.json -cache-aof-file ./data/dns_cache.aof  --ip-priority ipv6 "-upstream-address" "socks5://127.0.0.1:20808" "-upstream-type" "socks5" "--upstream-resolve-ips=true"  --upstream-username "hp40emnw108got6a67p2isj1x65qwjtz60fh5dtl7nhjhor3va" --upstream-password "i7esr1nwxcil034gslw4sdzjyfejfvf5xiaagx4x286nw6l3ff" 






# 定义生成随机英文字符串的函数，默认长度6位
function Get-RandomAlpha($length = 16) {
    $chars = 65..90 + 97..122 | ForEach-Object { [char]$_ }
    -join ($chars | Get-Random -Count $length)
}

for (; ; ) {
    # 生成随机前缀
    $prefix = Get-RandomAlpha 20
    # 拼接目标URL
    $targetUrl = "https://${prefix}.encryptedsni.com/cdn-cgi/trace"
    
    curl --proxy-user admin:iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58 -v -I $targetUrl -L -x http://127.0.0.1:28340 --doh-url https://pngwczx94z.cloudflare-gateway.com/dns-query  --connect-timeout 10 --max-time 10 -X GET

    Start-Sleep 5
}