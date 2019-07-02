## Edit your environment

```shell
% cp settings.toml.sample config/settings.toml
% emacs settings.toml
```

## build

```shell
% make install 
```

or

```
% statik -src=config
% go build -o BINNAME main.go
```
