<p align="center">
  What did he do?
  <br>
  Extract commit messages and create a prompt for ChatGPT
</p>
<p align="center">
  <img alt="Commit messages extractor" src="https://imgur.com/Qnxbuxe.png">
</p>
<p align="center">
  <img alt="Ask GPT" src="https://imgur.com/ocLOA16.png">
</p>


## Install

```sh
curl -sL https://raw.githubusercontent.com/kevincobain2000/wdhd/master/install.sh | sh
```

or via go

```sh
go install github.com/kevincobain2000/wdhd@latest
```

## Basic Usage

```sh
# for Github
wdhd --token=$GITHUB_TOKEN --user=kevincobain2000

# for Github Enterprise
wdhd --base-url=ghe.enterprise-me.com --token=$GHE_TOKEN --user=kevin.cobain

# Custom prompt
wdhd --prompt="What did he do today?" --days=2 --token=$GHE_TOKEN --user=kevin.cobain
```


## CHANGELOG

- **v1.0.0** - Initial release
