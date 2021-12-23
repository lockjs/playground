# PHP Experiments / Tutorials

A collection of completed tutorials / experimentsusing PHP

### Run

```shell
# Build the container
docker build -t playground:php .

# Run a CLI project
docker_run playground:php php <path_to_project>

# Run a HTTP project
docker_run -p 8080:8080 playground:php php -S 0.0.0.0:8080 -t <path_to_project>
```

## Tutorials

- Slim 4 Tutorial by Danial Opitz (source: https://odan.github.io/2019/11/05/slim4-tutorial.html)