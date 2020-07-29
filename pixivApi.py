from pixivpy3 import *

api = AppPixivAPI()
api.login("loveloliii", "")   # Not required
json_result = api.illust_detail(59580629)
illust = json_result.illust
print(">>> origin url: %s" % illust.image_urls['large'])


# get ranking: 1-30
# mode: [day, week, month, day_male, day_female, week_original, week_rookie, day_manga]
json_result = api.illust_ranking('day')
for illust in json_result.illusts:
    print(" p1 [%s] %s" % (illust.title, illust.image_urls.large))
    api.download(illust.image_urls.large)

# next page: 31-60
next_qs = api.parse_qs(json_result.next_url)
json_result = api.illust_ranking(**next_qs)
for illust in json_result.illusts:
    print(" p2 [%s] %s" % (illust.title, illust.image_urls.medium))