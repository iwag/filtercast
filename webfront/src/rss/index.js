import React, { PropTypes } from 'react';
import Layout from '../../components/Layout';
import {Button, Grid, Icon, IconButton, Checkbox,Cell  } from 'react-mdl';
import config from '../../components/Config';

import {grpc, Code, Metadata} from "grpc-web-client";
import {RssService} from "../../js/_proto/library/book_service_pb_service";
import {QueryRssRequest, Rss, GetRssRequest} from "../../js/_proto/library/book_service_pb";


const host = "http://localhost:8080";

function getRss() {
  const getRssRequest = new GetRssRequest();
  getRssRequest.setId("111");
  grpc.unary(RssService.GetRss, {
    request: getRssRequest,
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
getRss();

class UserRssPage extends React.Component {

  componentDidMount() {
    document.title = this.props.route.params.id + "'s RSS'";
  }

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Layout name={this.props.route.params.id} >
      <div >
          <h3>{this.props.data.url}</h3>
          <Grid className="demo-grid-1">
            <Cell col={12}>
              <a href={config.host + "/rss/" + this.props.data.id + "/feed.rss"}>
              Go to feed.rss!
              </a>
            </Cell>
            <Cell col={12}>
              <a  href={config.host + "/api/rss/" + this.props.data.id + "/publish"}>
                Publish one episode now!
              </a>
            </Cell>
        </Grid>
      </div>
      </Layout>
    );
  }

}

export default UserRssPage;
