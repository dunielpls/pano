import React from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

import { Navbar } from "react-bootstrap";

export default class App extends React.Component {
    constructor() {
        super();

        this.state = {};
    }

    render() {
        return (<>
            <div>Hello, World!</div>
        </>);
    }
}
