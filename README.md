
A tool for podcast
====

Allow to rerun/filter an episode of podcast.

## how to use

* Register the podcast feed's RSS. You can choose a way of publishing at random or reverse(old to new).
* Subscribe the RSS into podcast app you get from this app.

System overview
===
<img src="https://i.gyazo.com/55f3295e619ef684ca45f982ddd81734.png" width=520px />

Server side is written in Golang and frontend part in Javascript(ES6), ReactJS.
Whole app is run in Google Appengine Standard Environment, working with cloud datastore and memcached(cloud DNS).
[goon](https://github.com/mjibson/goon) is used for a utility for memcached and cloud datasore. WEB frontend is with [react-static-boilerplate](https://github.com/kriasoft/react-static-boilerplate).

Library/Languages
===
* golang
* [labstack/echo](https://echo.labstack.com/) as http framework
* [goon](https://github.com/mjibson/goon) as datastore client
* App Engine Standard Environment
* [react-starter-kit](https://github.com/kriasoft/react-starter-kit) as frontend

how to run
===

```bash
# google appengine sdk is neccesary. follow [this instruction](https://cloud.google.com/appengine/docs/standard/go/download_
$ goapp serve . # run locally
$ appcfg.py -A YOUR_PROJECT_ID -V v1 update .
# open YOUR_PROJECT_ID.appspot.com
```
