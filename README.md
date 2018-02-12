
Very best tool for podcast
====

Allow to rebroadcast/filter episodes of podcast by modifying RSS feed.

## how to use

* Register the podcast feed's RSS. You can choose a way of publishing at random or reverse(old to new).
* Subscribe the RSS into podcast app you get from this app then user can listen podcasts rebroadcasted.


Library/Languages
===
* Java8
* SpringBoot
* App Engine Standard Environment
* [react-starter-kit](https://github.com/kriasoft/react-starter-kit) as frontend

System overview
===
<img src="https://i.gyazo.com/55f3295e619ef684ca45f982ddd81734.png" width=520px />

CI strategy
====
(TBD)

API Documentation
===

## basis response

if success, http status code must be 200.

```json
{"status":200}
```

When errors, status would be 4xx, 5xx and it has additional info.

```json
{
"status":400,
"message":"rss id doesn't exist"
"debug":"...."
}
```


## GET /api/rss/:rid.xml

Returns a RSS feed of rss id.
RSS returned follows [the specification]()


## POST /api/rss

Create new RSS feed by Entity json.

## DELETE /api/rss/:rid

Delete RSS.

## GET /api/rss/:rid

Returns a status of rss by id which format is json.
This api is not open and permits to use by only admin.

## GET /api/rss/_all

Returns all rss by json.
This api is not open and permits to use by only admin.

## GET /api/rss/:rid/_refresh

Force to update a rss of rss id.
This api is not open and permits to use by only admin.

## GET /user/:uid/rss

Returns a user's rss.


how to run
===

```bash
# google appengine sdk is neccesary. follow [this instruction](https://cloud.google.com/appengine/docs/standard/go/download_
$ goapp serve . # run locally
$ appcfg.py -A YOUR_PROJECT_ID -V v1 update .
# open YOUR_PROJECT_ID.appspot.com
```
