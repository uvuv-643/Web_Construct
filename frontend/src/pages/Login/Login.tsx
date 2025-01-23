import {Alert, Button, TextInput} from '@gravity-ui/uikit';
import './Login.css';
import {SyntheticEvent, useState} from 'react';
import axios from 'axios';
import {useNavigate} from 'react-router';

type LoginProps = {
    triggerRerenderHeader: () => void;
};

export const Login = ({triggerRerenderHeader}: LoginProps) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loading, setLoading] = useState(false);

    const [success, setSuccess] = useState(false);

    const navigate = useNavigate();

    const onSubmit = (e: SyntheticEvent) => {
        e.preventDefault();
        e.stopPropagation();

        setLoading(true);
        setTimeout(() => {
            axios
                .post(
                    'http://localhost:8031/api/auth/login',
                    {
                        email,
                        password,
                    },
                    {withCredentials: false},
                )
                .then((response) => {
                    setSuccess(true);
                    setTimeout(() => {
                        setSuccess(false);
                        navigate('/');
                        triggerRerenderHeader();
                    }, 1500);
                    localStorage.setItem('jwt', response.data);
                })
                .catch((err) => {
                    if (err.response.status === 401) {
                        alert('Некорректный логин или пароль');
                    } else {
                        alert('Ошибка сервера');
                    }
                    // eslint-disable-next-line no-console
                    console.log(err.response);
                })
                .finally(() => {
                    setLoading(false);
                });
        }, 500);
    };

    return (
        <div>
            {success && <Alert theme="success" message="Вы успешно авторизовались" />}
            <div className="page">
                <section className="form">
                    <form onSubmit={onSubmit}>
                        <div className="form-element">
                            <h1>Войти в аккаунт</h1>
                        </div>
                        <div className="form-element">
                            <TextInput
                                placeholder="Email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
                        </div>
                        <div className="form-element">
                            <TextInput
                                placeholder="Пароль"
                                value={password}
                                type="password"
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>
                        <div className="form-element">
                            <Button view="action" type="submit" loading={loading}>
                                Войти
                            </Button>
                        </div>
                    </form>
                </section>
            </div>
        </div>
    );
};
