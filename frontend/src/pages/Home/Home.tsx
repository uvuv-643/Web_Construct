import './Home.css';
import {ArrowUpRightFromSquare, At, BroomMotion} from '@gravity-ui/icons';

export const Home = () => {
    return (
        <div className="homepage">
            <div className="page">
                <header className="homepage-header">
                    <h1 className="homepage-title">Добро пожаловать в Web Construct</h1>
                    <p className="homepage-subtitle">Создавайте потрясающие сайты с легкостью</p>
                    <button className="homepage-cta">Начать</button>
                </header>

                <section className="features-section">
                    <h2 className="section-title">Наши особенности</h2>
                    <div className="features-container">
                        <div className="feature-card">
                            <div className="feature-icon">
                                <ArrowUpRightFromSquare
                                    width={36}
                                    height={36}
                                ></ArrowUpRightFromSquare>
                            </div>
                            <h3>Простота использования</h3>
                            <p>Интерфейс перетаскивания с гибкими возможностями настройки.</p>
                        </div>
                        <div className="feature-card">
                            <div className="feature-icon">
                                <At width={36} height={36}></At>
                            </div>
                            <h3>Адаптивные дизайны</h3>
                            <p>
                                Дизайны, которые прекрасно смотрятся на любом устройстве и экране.
                            </p>
                        </div>
                        <div className="feature-card">
                            <div className="feature-icon">
                                <BroomMotion width={36} height={36}></BroomMotion>
                            </div>
                            <h3>Встроенные инструменты SEO</h3>
                            <p>Оптимизируйте ваш сайт для поисковых систем без усилий.</p>
                        </div>
                    </div>
                </section>
            </div>
        </div>
    );
};
