# docker-ui

UI for docker. A great example of a full stack on go + vue.
 
## Up an running

Is very simple! 

```
docker run -v /var/run/docker.sock:/var/run/docker.sock -p 9000:9000 --restart always -it docker-ui
```

open at http://localhost:9000.

See `examples` dir.

