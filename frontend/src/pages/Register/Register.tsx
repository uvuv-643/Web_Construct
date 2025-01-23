import {Alert, Button, TextInput} from '@gravity-ui/uikit';
import axios from 'axios';
import {SyntheticEvent, useState} from 'react';
import {useNavigate} from 'react-router';

type RegisterProps = {
    triggerRerenderHeader: () => void;
};

export const Register = ({triggerRerenderHeader}: RegisterProps) => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordConfirm, setPasswordConfirm] = useState('');
    const [loading, setLoading] = useState(false);

    const [success, setSuccess] = useState(false);

    const navigate = useNavigate();

    const onSubmit = (e: SyntheticEvent) => {
        e.preventDefault();
        e.stopPropagation();

        if (password !== passwordConfirm) {
            alert('Пароли не совпадают');
            return;
        }

        setLoading(true);
        setTimeout(() => {
            axios
                .post(
                    'http://localhost:8031/api/auth/register',
                    {
                        email,
                        password,
                        full_name: name,
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
            {success && <Alert theme="success" message="Вы успешно зарегистрировались" />}
            <div className="page">
                <section className="form">
                    <form onSubmit={onSubmit}>
                        <div className="form-element">
                            <h1>Создать аккаунт</h1>
                        </div>
                        <div className="form-element">
                            <TextInput
                                placeholder="Ваше имя"
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                            />
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
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>
                        <div className="form-element">
                            <TextInput
                                placeholder="Повторите пароль"
                                type="password"
                                value={passwordConfirm}
                                onChange={(e) => setPasswordConfirm(e.target.value)}
                            />
                        </div>
                        <div className="form-element">
                            <Button view="action" loading={loading} type="submit">
                                Зарегистрироваться
                            </Button>
                        </div>
                    </form>
                </section>
            </div>
        </div>
    );
};
