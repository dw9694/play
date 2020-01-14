# Play video

Get token vkontakte
```
https://oauth.vk.com/authorize?client_id=CLIENT_ID&scope=1073737727&redirect_uri=https://oauth.vk.com/blank.html&display=page&response_type=token&revoke=1
```

Build
```bash
docker build -t play .
```
Run
```bash
docker run -d -p 7531:7531 --env-file .env --restart unless-stopped \
  --name play play
```
