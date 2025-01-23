import Header from './components/Header/Header';
import {BrowserRouter, Route, Routes} from 'react-router';
import {Home} from './pages/Home/Home';
import {Login} from './pages/Login/Login';
import {Register} from './pages/Register/Register';
import {useState} from 'react';
import {Alert} from '@gravity-ui/uikit';
import {Create} from './pages/Create/Create';
import {My} from './pages/My/My';

export const App = () => {
    const [renderState, setRenderState] = useState(1);

    const triggerRerenderHeader = () => {
        setRenderState(renderState + 1);
    };

    const [loginMessage, setLoginMessage] = useState('');

    return (
        <BrowserRouter>
            <Header renderState={renderState} setLoginMessage={setLoginMessage} />
            {loginMessage.length ? <Alert message={loginMessage} theme="success"></Alert> : <></>}
            <Routes>
                <Route path="/" element={<Home />}></Route>
                <Route
                    path="/login"
                    element={<Login triggerRerenderHeader={triggerRerenderHeader} />}
                ></Route>
                <Route
                    path="/register"
                    element={<Register triggerRerenderHeader={triggerRerenderHeader} />}
                ></Route>
                <Route path="/create" element={<Create />}></Route>
                <Route path="/my" element={<My />}></Route>
            </Routes>
            <footer className="footer-section">
                <p>&copy; 2025 Web Construct. Все права защищены.</p>
            </footer>
        </BrowserRouter>
    );
};

export default App;
