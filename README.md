## Description

A small note taking web-app which I am building for myself just for fun and learning purposes.
My goal is too use as much of the standard library as possible without using too many frameworks.

## Setup

- change the middle part of filename of `.env.example.json` into the appropriate environment name (e.g.: `.env.dev.json`) and set your own values to the json fields.
- if you change the name to anything other than `.env.dev.json`, make sure to set the environment variable `ENVIRONMENT` equal to the middle part of the file name (e.g.: `.env.prod.json` => `ENVIRONMENT=prod`)
- run `start.sh`
