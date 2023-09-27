import React from 'react';

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

    return <div style={styles}>Hello from ComponentA with data: {props.data}</div>;
}

export default ComponentA;
