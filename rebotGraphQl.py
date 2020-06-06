import requests

url = "https://api.github.com/graphql"
params = {'query': '''{
  viewer {
    name
  }
}'''}
res = requests.post(url=url, params=params)
res.close