import redis
client = redis.Redis()

i = 0
while True:
    client.set('id', i)
    user = client.get('id')
    print(user) 
    i = i + 1
