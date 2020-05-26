import requests
import bs4

url = 'https://github.com/zytx121/je/issues/'
i = 1;
total=''
while i<4:
    url = url + str(i)
    response = requests.get(url)
    status_code = response.status_code

    content = bs4.BeautifulSoup(response.content.decode("utf-8"),"lxml")
    code = content.find_all('table')
    if(status_code == '404'): 
        print('404')
    else:
        print(status_code)
        total = total + ' /t '+str(code[0].text.split()[0])
        i = i + 1

fh = open('code.txt','w')
fh.write(total)
fh.close()
