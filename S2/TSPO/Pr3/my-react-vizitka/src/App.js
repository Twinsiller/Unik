import React, { useState, useMemo } from 'react';
import './App.css';

const App = () => {
  const [darkMode, setDarkMode] = useState(false);

  const profile = {
    name: "Болдинов Алексей Валерьевич",
    title: "Чел",
    image: "https://i.pinimg.com/originals/83/1a/7c/831a7cf39e76484168b529d4486f97db.jpg",
    about:
      "Вы бы не смогли смотреть на то как я плох в этом. Но разачарованы не Вы здесь, а я (от самого себя). Позор - это моё второе имя. Зато я немного умею играю в волейбол. ЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫЫ",
    skills: ["C#", "C++", "HTML5", "CSS3", "Git", "Webpack", "API Integration"],
    experience: [
      {
        position: "Учитель информатики",
        company: "ГБОУ Школа №1400",
        period: "2022 - 2024",
        description: "Обучение учеников навыкам программирования.",
      },
    ],
    contacts: {
      email: "boldinov.leha@yandex.ru",
      phone: "+7 (919) 965-22-30",
      github: {
        text: "github.com/Twinsiller",
        link: "https://github.com/Twinsiller",
      },
    },
  };

  const currentYear = useMemo(() => new Date().getFullYear(), []);

  const toggleTheme = () => {
    setDarkMode(!darkMode);
  };

  const getContactIcon = (type) => {
    const icons = {
      email: "✉️",
      phone: "📱",
      github: "💻",
      linkedin: "🔗",
    };
    return icons[type] || "📌";
  };

  return (
    <div id="app">
      <button onClick={toggleTheme} className="theme-toggle">
        {darkMode ? 'Светлая тема' : 'Тёмная тема'}
      </button>

      <div className={`container ${darkMode ? 'dark-mode' : ''}`}>
        <header>
          <img src={profile.image} alt="Фото профиля" className="profile-img" />
          <h1>{profile.name}</h1>
          <p>{profile.title}</p>
        </header>

        <section className="section">
          <h2>Обо мне</h2>
          <p>{profile.about}</p>
        </section>

        <section className="section">
          <h2>Навыки</h2>
          <div className="skills-list">
            {profile.skills.map((skill, index) => (
              <span className="skill-tag" key={index}>{skill}</span>
            ))}
          </div>
        </section>

        <section className="section">
          <h2>Опыт работы</h2>
          {profile.experience.map((job, index) => (
            <div key={index} className="job">
              <h3>{job.position}</h3>
              <p><strong>{job.company}</strong> | {job.period}</p>
              <p>{job.description}</p>
            </div>
          ))}
        </section>

        <section className="section">
          <h2>Контакты</h2>
          <ul className="contact-list">
            {Object.entries(profile.contacts).map(([key, value]) => (
              <li key={key}>
                <i>{getContactIcon(key)}</i>
                {typeof value === 'object' ? (
                  <a href={value.link} target="_blank" rel="noopener noreferrer">
                    {value.text}
                  </a>
                ) : (
                  <span>{value}</span>
                )}
              </li>
            ))}
          </ul>
        </section>

        <footer>
          <p>© {currentYear} {profile.name}. Все права защищены.</p>
        </footer>
      </div>
    </div>
  );
};

export default App;
