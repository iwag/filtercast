import React, { PropTypes } from 'react';
import Layout from '../../components/Layout';
import {Button, Grid, Icon, IconButton, Cell} from 'react-mdl';

class UserRssPage extends React.Component {

  componentDidMount() {
    document.title = "admin";
  }

  constructor(props) {
    super(props);
  }

  render() {
    var items = this.props.data.map((item, i) => {
      return (<Cell col={12}>
      <a href={"/rss/" + item.Id}>{item.Id}</a> <a href={item.Url}> {item.Url}</a>
      </Cell>)
    });

    return (
      <Layout name={this.props.route.params.id} >
      <div >
          <h3>{this.props.data.url}</h3>
          <Grid className="demo-grid-1">
          {items}
        </Grid>
      </div>
      </Layout>
    );
  }

}

export default UserRssPage;
