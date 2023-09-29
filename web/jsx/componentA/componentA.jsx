import React from 'react';
import JSONPretty from 'react-json-pretty';
import 'react-json-pretty/themes/monikai.css';


function ComponentA(props) {
    console.log("ComponentA is rendering with props:", props);

    const styles = {
        width: '200px',
        height: '100px',
        backgroundColor: 'lightblue',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        border: '1px solid darkblue',
        borderRadius: '5px',
        padding: '10px',
        margin: '5px 0'
    };

    return <JSONPretty id="json-pretty" data={props}></JSONPretty>

    // return <div style={styles}>Hello from ComponentA with data: {props.data}</div>;
}

export default ComponentA;
