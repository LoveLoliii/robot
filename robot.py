import requests
import bs4
import json

url = 'https://github.com/zytx121/je/issues/'
i = 473
total=''
while i<1132:#1132
    urln = url + str(i)
    print(urln)
    response = requests.get(urln)
    status_code = response.status_code

    content = bs4.BeautifulSoup(response.content.decode("utf-8"),"lxml")
    code = content.find_all('table')
    if(status_code == 404): 
        print('404_')
    else:
        print(status_code)
        #获取h1
        h1 = code[0].find('h1')
        title = h1
        pic = code[0].find('p')
        if(code[0].find('li') is None):
            singer=""    
        else:
            singer = code[0].find('li').text  
        score = code[0].find('code')
        #total = total + ' \t\n OxO '+str(code[0].text)
        #try use api save data to server
        res = {"title":title,"pic":pic,"singer":singer,"score":score}
        print(str(code[0].text)) 
        _=requests.post("http://localhost:8888/addSong",data = res )
    i = i + 1

fh = open('code.txt','a')
fh.write(total)
fh.close()
