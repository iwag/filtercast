/**
 * React Static Boilerplate
 * https://github.com/kriasoft/react-static-boilerplate
 *
 * Copyright © 2015-present Kriasoft, LLC. All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE.txt file in the root directory of this source tree.
 */

import React from 'react';
import Link from '../Link';
import {Icon, Cell  } from 'react-mdl';
import 'react-mdl/extra/material.css';
import 'react-mdl/extra/material.js';
import $ from 'jquery';
import config from '../Config';

class Navigation extends React.Component {

  constructor() {
    super();
    this.state = {profile:null};
  }

  componentDidMount() {
    window.componentHandler.upgradeElement(this.root);
  }

  componentWillUnmount() {
    window.componentHandler.downgradeElements(this.root);
  }

  render() {
    var icon, name;
      return (
        <nav className="mdl-navigation" ref={node => (this.root = node)}>
          <Link className="mdl-navigation__link" to="/">Home</Link>
        </nav>
      );
  }

}

export default Navigation;
