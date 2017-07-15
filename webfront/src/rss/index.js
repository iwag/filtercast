import React, { PropTypes } from 'react';
import Layout from '../../components/Layout';
import $ from 'jquery';
import {Button, Grid, Icon, IconButton, Checkbox,Cell  } from 'react-mdl';
import config from '../../components/Config';

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
