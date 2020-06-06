import requests
import bs4

url = 'https://github.com/zytx121/je/issues/'
i = 1
total=''
while i<1132:
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
        total = total + ' \t\n OxO '+str(code[0].text)
    i = i + 1

fh = open('code.txt','a')
fh.write(total)
fh.close()
