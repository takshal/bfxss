# kal-xss
this tool help you to find blind  cross site scripting.


description: i belive you can find blind cross site scripting vlnerablity in user-agent,host,Referer header. so i have created this go base cli tool to find blind cross site scripting vulnerablity

what require?
so you need xss hunter paylod or account.

how to use?

``` assetfinder --subs-only example.com | httprobe | kal-xss -p "your xsshunter payload" ```

how to install?
