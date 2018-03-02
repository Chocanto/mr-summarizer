MR Summarizer
=============

MR Summarizer is a Go Mattermost bot that scan all merge requests of all projects from a Gitlab group
that need to be reviewed and send a nice summary table in the chosen channel.

Usage
------------

**Requirements** :
- Go toolchain
- `$GOPATH/bin` in $PATH

**Get and build**
```
go get gitlab.kazan.priv.atos.fr/A643410/mr-summarizer
```

**Create a new configuration file**
config.yml
```yml
mattermost:
  webhook: https://mattermost.com/hooks/sfdsqsdf3q88gj83bs7
  channel: my-channel
  username: MR Summarizer

gitlab:
  url: https://gitlab.com/api/v4
  token: DFsd5fsF8C1df1
  group: my-group

threshold: 7
title: MRs FOR APPROVAL
```

**Run**
```
mr-summarizer config.yml
```