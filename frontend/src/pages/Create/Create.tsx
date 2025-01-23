import {Alert, Button, TextArea} from '@gravity-ui/uikit';
import axios from 'axios';
import {SyntheticEvent, useState} from 'react';
import {useNavigate} from 'react-router';

export const Create = () => {
    const [loading, setLoading] = useState(false);
    const [success, setSuccess] = useState(false);
    const [uuid, setUuid] = useState('');

    const [request, setRequest] = useState('');

    const navigate = useNavigate();

    const onSubmit = (e: SyntheticEvent) => {
        e.preventDefault();
        e.stopPropagation();

        setLoading(true);
        setTimeout(() => {
            axios
                .post(
                    'http://localhost:8031/api/order',
                    {
                        request,
                    },
                    {
                        withCredentials: false,
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem('jwt')}`,
                        },
                    },
                )
                .then((response) => {
                    setSuccess(true);
                    setTimeout(() => {
                        setSuccess(false);
                        navigate('/my');
                    }, 2500);
                    setUuid(response.data);
                })
                .catch((err) => {
                    if (err.response.status === 401) {
                        alert('Вы не авторизованы');
                        localStorage.setItem('jwt', '');
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

    const message = `Вы успешно отправили запрос на создание компонента, ожидайте его появления в своем личном кабинете. uuid: ${uuid}`;
    return (
        <div>
            {success && <Alert theme="success" message={message} />}
            <div className="page">
                <section className="form">
                    <form onSubmit={onSubmit}>
                        <div className="form-element">
                            <h1>Создать новый компонент</h1>
                        </div>
                        <div className="form-element">
                            <TextArea
                                placeholder="Введите описание вашего компонента"
                                value={request}
                                rows={10}
                                onChange={(e) => setRequest(e.target.value)}
                            />
                        </div>
                        <div className="form-element">
                            <Button view="action" type="submit" loading={loading}>
                                Создать
                            </Button>
                        </div>
                    </form>
                </section>
            </div>
        </div>
    );
};
