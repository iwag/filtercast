import React from 'react';
import Layout from '../../components/Layout';
import s from './styles.css';
import {title, html} from './index.md';
import config from '../../components/Config';
import 'whatwg-fetch';
import history from '../history';

import {grpc, Code, Metadata} from "grpc-web-client";
import {RssService} from "../../js/_proto/library/book_service_pb_service";
import {QueryRssRequest, Rss, PutRssRequest} from "../../js/_proto/library/book_service_pb";

class TopPage extends React.Component {

  componentDidMount() {
    document.title = title;
  }

  render() {
    return (
      <Layout className={s.content}>
        <Input/>
        <div className="mdl-layout-spacer"></div>
        <div dangerouslySetInnerHTML={{
          __html: html
        }}/>
      </Layout>
    );
  }

}

class Combobox extends React.Component {
  render() {
	  var items = this.props.combolist.map((item, i) => {
		  return (<option key={i} value={item.key}>
		  {item.name}
		  </option>)
	  });

    return (
      <select name={this.props.name} className="form-control">
	  {items}
      </select>
    );
  }
}

class Input extends React.Component {


  constructor(props) {
    super(props);
    this.state = {};
    this.handleSubmit = this.handleSubmit.bind(this);
    this.url = config.host + "/api/rss/new";
    this.comboList = [{name: 'Random', key: 'random'},{name: 'Old to New', key: 'firstout'}];
  }

  componentDidMount() {
    document.title = "rssを登録";
    this.refs.url.focus();
  }

  checkStatus(r) {
    if (r.status >= 200 && r.status < 300) {
      return r
    } else {
      var error = new Error(r.statusText)
      error.response = r
      throw error
    }
  }


  putRss(e) {
    e.preventDefault();
    
    const host = "http://localhost:8080";
    const putRssRequest = new PutRssRequest();
    putRssRequest.setUrl(this.refs.url.value);
    grpc.unary(RssService.PutRss, {
      request: putRssRequest,
      host: host,
      onEnd: res => {
        const { status, statusMessage, headers, message, trailers } = res;
        console.log("getRss.onEnd.status", status, statusMessage);
        console.log("getRss.onEnd.headers", headers);
        if (status === Code.OK && message) {
          console.log("getRss.onEnd.message", message.toObject());
        }
        console.log("getRss.onEnd.trailers", trailers);
        // queryRss();
      }
    });
  }


  handleSubmit(e) {
    e.preventDefault();
    var way = this.comboList[0].key;
    for (let i of this.comboList) {
      if (i.name == this.refs.form.publishway.value)
        way = i.key;
    }
    var v = {
      url: this.refs.url.value,
      publish_way: way,
      duration: this.refs.time.value,
    };
    fetch(this.url, {
      method: "POST",
      body: JSON.stringify(v),
      headers: {
        "Content-Type": "application/json"
      },
      credentials: 'same-origin'
    }).then(this.checkStatus).then(r => r.json()).then(r => {
      history.push({pathname: "rss/" + r.id});
    }).catch(e => {
      console.log('request failed', e);
    });
  }

  render() {
    return (
      <form onSubmit={this.putRss.bind(this)} ref="form">
        <div className="mdl-textfield mdl-js-textfield" style={{
          display: "table-cell",
          padding: "5px 0px"
        }}>
          RSS's URL:<textarea className="mdl-textfield__input" type="text" rows="1" ref="url" name="url"
          defaultValue="http://example.com/feed.rss" style={{
            width: "520pt",
            "font-size": 3 + "em",
            border: "1px solid rgba(0,0,0,.12)"
          }}></textarea>
        </div><br />
        <div className="mdl-textfield mdl-js-textfield" style={{
          display: "table-cell",
          padding: "5px 0px 1px 1px"
        }}>
        Duration:
        <input defaultValue="12h" type="text" ref="time" style={{
                  width: "180pt",
                  "font-size": 3 + "em",
                  border: "1px solid rgba(0,0,0,.12)"
                }}/>
        </div>
        <Combobox name="publishway" ref="way" combolist={this.comboList} />
        <button type="submit" className="mdl-button mdl-js-button" style={{
          width: 120 + "pt"
        }}>Enter</button>
      </form>
    );
  }

}

export default TopPage;
