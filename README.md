## Install

[Download the latest binary](https://github.com/silviomm/gorabbit/releases) and place it under a folder in your PATH;

## Roadmap:
- Different output formats;
- Default values rabbitmqconfig prompt;
- Accept flags to fill prompts;
- create queue;
- send msgs;
- consume msgs;
- get/delete based on regex;
- shovel:
    - when consumed all msgs, send notification to channel to stop consuming (option by flag) 
    - specify the rate
    - dispose connections on ^C
    - declare queue on IN rabbit if it doesn't exists (option by flag)
    - better output to show messages consumed/sent
    - ignore if queue to consume does not exist but send a warning;