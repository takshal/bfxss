# bfxss
this tool help you to find fast blind  cross site scripting.


description: i belive you can find blind cross site scripting vlnerablity in user-agent,host,Referer header. so i have created this go base cli tool to find blind cross site scripting vulnerablity

what require?
so you need xss hunter paylod or account.

<h5>how to use?</h5>

``` assetfinder --subs-only example.com | httprobe | bfxss -p "your xsshunter payload" ```

<h5>how to install?</h5>

``` go install github.com/takshal/bfxss@latest ```
