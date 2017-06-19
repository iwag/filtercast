import React from 'react';
import Layout from '../../components/Layout';
import s from './styles.css';
import {title, html} from './index.md';
import config from '../../components/Config';
import 'whatwg-fetch';
import history from '../history';

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

class Input extends React.Component {

  constructor(props) {
    super(props);
    this.state = {};
    this.handleSubmit = this.handleSubmit.bind(this);
    this.url = config.host + "/api/rss/new";
  }

  componentDidMount() {
    document.title = "タスクを登録";
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

  handleSubmit(e) {
    e.preventDefault();
    var v = {
      url: this.refs.url.value
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
      <form onSubmit={this.handleSubmit.bind(this)}>
        <div className="mdl-textfield mdl-js-textfield" style={{
          display: "table-cell",
          padding: "5px 0px"
        }}>
          <textarea className="mdl-textfield__input" type="text" rows="1" ref="url" name="url" style={{
            width: "320pt",
            "font-size": 2 + "em",
            border: "1px solid rgba(0,0,0,.12)"
          }}></textarea>
        </div>
        <button type="submit" className="mdl-button mdl-js-button" style={{
          width: 100 + "pt"
        }}>Enter</button>
      </form>
    );
  }

}

export default TopPage;
