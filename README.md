# pr0bot

A simple Go bot that posts content from pr0gram to a telegram channel

## Getting started

```bash
go get -u

```

## Setting up config file

<details>
    <summary><b>Click here for more details</b></summary>

Fill up rest of the fields. Meaning of each fields are described below:

- **tags**: write here your custom filter like I did in config_example.yaml
- **mongodb_url**: Your MongoDB-URL
- **bot_token**: Your telegram bot token
- **cookie**: Your pr0gramm cookie
- **channel**: Your channel/group id

</details>

## Start the bot

```bash
go build . && ./pr0.bot

```

## Check out my channel

https://t.me/geistigeerguesse
