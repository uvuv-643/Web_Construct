import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

import '@gravity-ui/uikit/styles/fonts.css';
import '@gravity-ui/uikit/styles/styles.css';
import './global.css';
import {ThemeProvider} from '@gravity-ui/uikit';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
    <React.StrictMode>
        <ThemeProvider theme="light">
            <App />
        </ThemeProvider>
    </React.StrictMode>,
);
