go run ./cmd/ -enable-pprof  '-dohurl' 'https://doh-server.masx200.ddns-ip.net' '-dohip' '104.21.14.41' '--dohalpn=h2' '-port' '28340' '-username' 'admin' '-password' 'iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58' '-dohalpn=h3' '-dohurl' 'https://doh-server.masx200.ddns-ip.net' '--dohip=104.21.9.230' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h2' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h3' '-dohurl' 'https://pngwczx94z.cloudflare-gateway.com/dns-query' '-dohip' '104.21.9.230' '--dohalpn=h3' -cache-file ./data/dns_cache.json -cache-aof-file ./data/dns_cache.aof  --ip-priority ipv6







for (; ; ) {

    curl  --proxy-user admin:iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58 -v -I https://dash.cloudflare.com/ -L -x http://127.0.0.1:28340 --doh-url https://doh.opendns.com/dns-query --connect-timeout 10 --max-time 10
    
    sleep 2
    curl --proxy-user admin:iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58 -v  https://ipv4.ipleak.net/?mode=json -L -x http://127.0.0.1:28340 --doh-url https://doh.opendns.com/dns-query  --connect-timeout 10 --max-time 10
    sleep 5
    curl --proxy-user admin:iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58 -v  https://ifconfig.co/json -L -x http://127.0.0.1:28340 --doh-url https://doh.opendns.com/dns-query  --connect-timeout 10 --max-time 10
    sleep 5
    curl --proxy-user admin:iy3w0rqwftfb1z7jr2nd4c894rc8t3pxhtw1qj94bxnjvioq58 -v -I https://note.wps.cn/ -L -x http://127.0.0.1:28340 --doh-url https://doh.opendns.com/dns-query  --connect-timeout 10 --max-time 10;



    sleep 5;
}