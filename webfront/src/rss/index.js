import React, { PropTypes } from 'react';
import Layout from '../../components/Layout';
import $ from 'jquery';
import {Button, Grid, Icon, IconButton, Checkbox,Cell  } from 'react-mdl';
import config from '../../components/Config';

class UserTaskPage extends React.Component {

  componentDidMount() {
    document.title = this.props.route.params.id + "さんのタスク一覧";
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
            </Cell>
            <Cell col={12}>
              <h3><a href={config.host + "/rss/" + this.props.data.id + "/feed.rss"}>Go!</a></h3>
            </Cell>
        </Grid>
      </div>
      </Layout>
    );
  }

}

class Task extends React.Component {

  constructor(props) {
    super(props);
  }

  render() {
      return (
        <tr>
         <td></td>
         <td className="mdl-data-table__cell--non-numeric" ><strong>{this.props.w.text}</strong></td>
        </tr>
      )
  }
}

export default UserTaskPage;
