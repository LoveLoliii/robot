import requests
import bs4
import json
import random
import time
#ac匿名版 novel爬取
#一些配置
#此url是将要爬取的url
#url = 'https://tnmb.org/t/163880?page='
#t为串号
#t='1376846'
#url = 'https://tnmb.org/t/'+t+'?page='
t='34139581'
url = 'https://adnmb3.com/t/'+t+'?page='
#起始页
i = 1
#总页数
pages=13
#过滤字数
filter_words=50
#保存爬取内容
total=""
#user_agent 集合
USER_AGENTS = [
    "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
    "Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
    "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
    "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
    "Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
    "Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
    "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
    "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
    "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
    "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
]
#此i代表将要爬取的page
while i<=pages:
    urln = url + str(i)
    #print(urln)
    headers={'User-Agent':random.choice(USER_AGENTS)}
    print(headers)
    response = requests.get(urln,headers=headers)
    status_code = response.status_code
    content = bs4.BeautifulSoup(response.content.decode("utf-8"),"lxml")
    #print(content)
    code = content.find_all(name='div',attrs={"class":'h-threads-content'})
    #print(code)
    for contents in code:
        strs = str(contents.text);
        #print(len(strs));
        #粗暴的过滤评论（存在误伤或过滤不全的问题
        if  len(strs) > filter_words:
            #print(strs);
            if strs in total:
                #过滤每页的主贴重复
                print("已存在");
            else:    
                total+="\n";
                total+=strs;
            
    i+=1
fh = open('novel_'+t+'_'+str(time.time())+'.txt','a',encoding='utf-8')
fh.write(total)
fh.close()
