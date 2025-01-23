/* eslint-disable @typescript-eslint/no-redeclare */
import './Header.css';
import {Button, Text} from '@gravity-ui/uikit';
import {useEffect, useState} from 'react';
import {Link} from 'react-router';

type HeaderProps = {
    renderState: number;
    setLoginMessage: (s: string) => void;
};
const Header = ({renderState, setLoginMessage}: HeaderProps) => {
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>();

    useEffect(() => {
        setIsLoggedIn(Boolean(localStorage.getItem('jwt')));
    }, [renderState]);

    const handleLogout = () => {
        localStorage.setItem('jwt', '');
        setLoginMessage('Вы успешно вышли из аккаунта');
        setTimeout(() => {
            setLoginMessage('');
        }, 1500);
        setIsLoggedIn(false);
    };

    return (
        <header className="header">
            <div className="header-left">
                <h1 className="header-logo">
                    <Link to="/">
                        <Text variant="display-1">Web Contruct</Text>
                    </Link>
                </h1>
                <nav className="nav">
                    <ul className="nav-list">
                        <li className="nav-item">
                            <Link to={isLoggedIn ? '/create' : '/login'} className="nav-link">
                                <Text variant="body-2">Создать компонент</Text>
                            </Link>
                        </li>
                        {isLoggedIn && (
                            <li className="nav-item">
                                <Link to="/my" className="nav-link">
                                    <Text variant="body-2">Мои компоненты</Text>
                                </Link>
                            </li>
                        )}
                    </ul>
                </nav>
            </div>
            <div className="header-right">
                {isLoggedIn ? (
                    <>
                        <Button view="outlined-contrast" size="xl" onClick={handleLogout}>
                            Выйти из аккаунта
                        </Button>
                    </>
                ) : (
                    <>
                        <Link to="/register">
                            <Button view="flat-contrast" size="xl">
                                Регистрация
                            </Button>
                        </Link>
                        <Link to="/login">
                            <Button view="outlined-contrast" size="xl">
                                Войти
                            </Button>
                        </Link>
                    </>
                )}
            </div>
        </header>
    );
};

export default Header;
