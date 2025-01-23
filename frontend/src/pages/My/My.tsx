import {useEffect, useState} from 'react';
import {ComponentCard, ComponentDto} from '../../components/ComponentCard/ComponentCard';
import axios from 'axios';
import {Loader} from '@gravity-ui/uikit';

export const My = () => {
    const [components, setComponents] = useState<ComponentDto[]>([]);

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios
            .get('http://localhost:8031/api/order', {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('jwt')}`,
                },
            })
            .then((response) => {
                console.log(response.data);
                setComponents(response.data ?? []);
            })
            .finally(() => {
                setLoading(false);
            });
    }, []);

    return (
        <div>
            <div className="page">
                <section className="components">
                    {loading && <Loader></Loader>}
                    {components.map((component) => (
                        <ComponentCard component={component} />
                    ))}
                </section>
            </div>
        </div>
    );
};
