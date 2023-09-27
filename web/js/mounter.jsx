import React from 'react';
import ReactDOM from 'react-dom';
import ComponentA from '../jsx/ComponentA/ComponentA';
import ComponentB from '../jsx/ComponentB/ComponentB';

const components = {
    ComponentA,
    ComponentB
};

// Expose the mount function globally so Go can call it
window.mountReactComponents = function() {
    console.log("Looking for component entry");
    const reactPlaceholders = document.querySelectorAll('[data-react-component]');

    reactPlaceholders.forEach(placeholder => {
        const componentName = placeholder.getAttribute('data-react-component');
        const encodedProps = placeholder.getAttribute("data-react-props");
        const propsString = atob(encodedProps);
        try {
            const componentProps = JSON.parse(propsString);
            const ReactComponent = components[componentName];
            if (ReactComponent) {
                console.log("rendering component with props ", componentProps);
                ReactDOM.render(<ReactComponent {...componentProps} />, placeholder);
            }
        } catch(e) {
            console.error("Failed to parse props:", propsString);
            return; // skip this iteration
        }
    });
};

// Call the function once when the DOM is fully loaded
document.addEventListener('DOMContentLoaded', window.mountReactComponents);
