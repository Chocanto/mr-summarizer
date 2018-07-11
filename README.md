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
language: en-US
```

**Run**
```
mr-summarizer config.yml
```

Configuration
-------------

### Mattermost configuration

| Parameter | Type   | Description                                 | Example                                          |
|-----------|--------|---------------------------------------------|--------------------------------------------------|
| webhook   | string | Webhook url for this bot                    | https://mattermost.com/hooks/sfdsqsdf3q88gj83bs7 |
| channel   | string | Channel name where to send messages         | my-cool-channel                                  |
| username  | string | Username of the bot (used on sent messages) | MR Summarizer                                    |

Example :
```
mattermost:
  webhook: https://mattermost.com/hooks/sfdsqsdf3q88gj83bs7
  channel: my-channel
  username: MR Summarizer
```


### Gitlab configuration

| Parameter | Type   | Description                                              | Example                   |
|-----------|--------|----------------------------------------------------------|---------------------------|
| url       | string | API endpoint of Gitlab instance                          | https://gitlab.com/api/v4 |
| token     | string | Private token used for Gitlab API Authentication         | DFsd5fsF8C1df1            |
| group     | string | Name of the group that will be scanned for merge request | my-group                  |

Example :
```
gitlab:
  url: https://gitlab.com/api/v4
  token: DFsd5fsF8C1df1
  group: my-group
```

### Other configuration

| Parameter | Type   | Description                                                  | Example          |
|-----------|--------|--------------------------------------------------------------|------------------|
| threshold | number | Number of days before the merge request is marked as warning | 7                |
| language  | string | BCP 47 tag language for translation and localization         | en-US            |

Example :
```
threshold: 7
language: en-US
```

Output example
--------------


### MERGE REQUESTS WAITING FOR APPROVAL - 02/03/18
|                                                  PROJECT                                                   |                                                                    TITLE                                  |    SINCE     | :+1: |           ASSIGNED         |
|-----------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------|--------------|------|----------------------------|
| [MR Summarizer](https://gitlab.kazan.priv.atos.fr/A643410/mr-summarizer)                                  | [Replace all bugs with features](https://gitlab.kazan.priv.atos.fr/A643410/mr-summarizer/merge_requests/1) | 8D :warning: |    1 | DUPOND MARC (@dupond.marc) |
| [Cool project](https://gitlab.kazan.priv.atos.fr/A643410/cool-project)                                    | [New name for this cool project](https://gitlab.kazan.priv.atos.fr/A643410/cool-project/merge_requests/9)  | 1D           |    0 | Anthony Granger (@A643410) |
