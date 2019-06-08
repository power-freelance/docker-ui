# docker-ui

UI for docker. A great example of a full stack on go + vue.
 
## Up an running

Is very simple! 

```
docker run \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -p 9000:9000 \
    --restart always \
    -it powerfreelance/docker-ui
```

open at http://localhost:9000.

See `examples` dir.


## TODO

Containers:

- Prune stopped containers
- See containers log
- Stop/Start/Restart container
- Group actions (start/stop/restart) for containers
- Inspect container
- Exec command in container and get output

Images:

- Prune dangling images
- Remove individual image

Compose:

- Show available compose projects
- Detail of projects, list of services

Networks:

- List 
- Prune

Volumes:

- List
- Prune

Common:

- System prune (prune networks, volumes, containers and images)
