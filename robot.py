#encoding:utf-8
import requests
import bs4
import json
import time
import random
import json
import numpy as np
# list 转成Json格式数据
def listToJson(lst):
    
    keys = [str(x) for x in np.arange(len(lst))]
    list_json = dict(zip(keys, lst))
    str_json = json.dumps(list_json, indent=2, ensure_ascii=False)  # json转为string
    return str_json
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
headers={'User-Agent':random.choice(USER_AGENTS)}
url = 'https://github.com/zytx121/je/issues/'
i = 1
#total=''
r = requests.get("https://github.com/zytx121/je/issues",headers=headers)
rc = bs4.BeautifulSoup(r.content.decode("utf-8"),"lxml")
ic = rc.find(class_='Box-row Box-row--focus-gray p-0 mt-0 js-navigation-item js-issue-row')
it = ic.text
start = int(it.find("#"))
# 截取4位 如果issue超过9999 或者小于1000 会出错
end = start + 4
xs = str(it)
x = 1166#int(xs[(start+1):(end+1)])
print(x)
raw_list = []
#x =int(x)
while i<=x:#x:#1132
    print(i)
    urln = url + str(i)
    #print(urln)
    headers={'User-Agent':random.choice(USER_AGENTS)}
    try:
        response = requests.get(urln,headers=headers)
    except ConnectionError:
        print("connectionError:"+urln)
    else:
        status_code = response.status_code
        content = bs4.BeautifulSoup(response.content.decode("utf-8"),"lxml")
        code = content.find('table')
        if(status_code == 404): 
            print('404_')
        else:
            try:
                raw  = code.prettify()
                raw_list.append(raw)
            except AttributeError:
                print("AttributeError")
        #print(status_code)
        #获取h1
        #h1 = code.find('h1')
        #title = h1
        #pic = code.find('p')
        #if(code.find('li') is None):
            #singer=""    
        #else:
            #singer = code.find('li').text  
        #score = code.find('code')
        #total = total + ' \t\n OxO '+str(code[0].text)
        #try use api save data to server
        #res = {"title":title,"pic":pic,"singer":singer,"score":score,"issue":i}
        #print(str(code.text)) 
        #_=requests.post("http://localhost:8888/addSong",data = res )
        i = i + 1
raw_json = listToJson(raw_list)
time = time.localtime(time.time())
t_str = str(time.tm_year)+'-'+str(time.tm_mon)+'-'+str(time.tm_mday)+"-"+str(time.tm_hour)+"_"+str(time.tm_min)+'_'+str(time.tm_sec)
fh = open('raw_json'+t_str+'.json','w',encoding="utf-8")
fh.write(raw_json)
fh.close()



