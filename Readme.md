# redirectgen

This tool is used to generate the friendly event redirect URLs in the form `http://belfastgopher.com/dec2020` with the
necessary tags to allow link previews in Twitter, LinkedIn and Slack. 

### Example from Twitter:
![Linnk preview on Twitter](images/twitter.png)

### Example from Slack:
![Linnk preview on LinkedIn](images/slack.png)


## Installation

```shell
go get github.com/belfastgophers/cmd/redirectgen
```

## Usage
You should run this from the root of our website repo

```shell
redirectgen {Event} {Description} {Meetup_ID} {Card}
```