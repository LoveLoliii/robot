import requests
import bs4
import json
url = 'https://tnmb.org/t/163880?page='
i = 1
while i<=68:
    urln = url + str(i)
    print(urln)
    response = requests.get(urln)
    status_code = response.status_code
    content = bs4.BeautifulSoup(response.content.decode("utf-8"),"lxml")
    print(content)
    code = content.find_all(name='div',attrs={"class":'h-threads-content'})
    #print(code)
    for content in code:
        print(content);
        #if(content.len()<100)
    i+=1
    # for(int i=0;i<code.size()):
    #     if(code[i].length>100):
    #         print(code[i])