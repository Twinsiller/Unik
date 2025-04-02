<template>
  <div id="app">
        <button @click="toggleTheme" class="theme-toggle">
            {{ darkMode ? 'Светлая тема' : 'Тёмная тема' }}
        </button>
        
        <div class="container" :class="{ 'dark-mode': darkMode }">
            <header>
                <img :src="profile.image" alt="Фото профиля" class="profile-img">
                <h1>{{ profile.name }}</h1>
                <p>{{ profile.title }}</p>
            </header>
            
            <section class="section">
                <h2>Обо мне</h2>
                <p>{{ profile.about }}</p>
            </section>
            
            <section class="section">
                <h2>Навыки</h2>
                <div class="skills-list">
                    <span class="skill-tag" v-for="(skill, index) in profile.skills" :key="index">
                        {{ skill }}
                    </span>
                </div>
            </section>
            
            <section class="section">
                <h2>Опыт работы</h2>
                <div v-for="(job, index) in profile.experience" :key="index" class="job">
                    <h3>{{ job.position }}</h3>
                    <p><strong>{{ job.company }}</strong> | {{ job.period }}</p>
                    <p>{{ job.description }}</p>
                </div>
            </section>
            
            <section class="section">
                <h2>Контакты</h2>
                <ul class="contact-list">
                    <li v-for="(contact, key) in profile.contacts" :key="key">
                        <i>{{ getContactIcon(key) }}</i>
                        <span v-if="key !== 'email' && key !== 'phone'">
                            <a :href="contact.link" target="_blank">{{ contact.text }}</a>
                        </span>
                        <span v-else>
                            {{ contact }}
                        </span>
                    </li>
                </ul>
            </section>
            
            <footer>
                <p>© {{ currentYear }} {{ profile.name }}. Все права защищены.</p>
            </footer>
        </div>
    </div>
</template>

<script>
const { createApp, ref, computed } = Vue;
        
        createApp({
            setup() {
                // Данные профиля
                const profile = ref({
                    name: "Иван Иванов",
                    title: "Frontend разработчик",
                    image: "https://via.placeholder.com/150",
                    about: "Я frontend разработчик с 3-летним опытом работы с Vue.js, React и современными веб-технологиями. Увлекаюсь созданием интуитивно понятных и отзывчивых пользовательских интерфейсов.",
                    skills: ["Vue.js", "JavaScript", "HTML5", "CSS3", "Git", "Webpack", "API Integration"],
                    experience: [
                        {
                            position: "Frontend Developer",
                            company: "WebTech Solutions",
                            period: "2020 - настоящее время",
                            description: "Разработка и поддержка клиентской части веб-приложений с использованием Vue.js."
                        },
                        {
                            position: "Junior Web Developer",
                            company: "Digital Agency",
                            period: "2018 - 2020",
                            description: "Верстка макетов, создание простых веб-сайтов, участие в командных проектах."
                        }
                    ],
                    contacts: {
                        email: "ivan.ivanov@example.com",
                        phone: "+7 (123) 456-78-90",
                        github: {
                            text: "github.com/ivanov",
                            link: "https://github.com/ivanov"
                        },
                        linkedin: {
                            text: "linkedin.com/in/ivanov",
                            link: "https://linkedin.com/in/ivanov"
                        }
                    }
                });
                
                // Текущий год для футера
                const currentYear = computed(() => new Date().getFullYear());
                
                // Темная тема
                const darkMode = ref(false);
                const toggleTheme = () => {
                    darkMode.value = !darkMode.value;
                };
                
                // Иконки для контактов
                const getContactIcon = (contactType) => {
                    const icons = {
                        email: "✉️",
                        phone: "📱",
                        github: "💻",
                        linkedin: "🔗"
                    };
                    return icons[contactType] || "📌";
                };
                
                return {
                    profile,
                    currentYear,
                    darkMode,
                    toggleTheme,
                    getContactIcon
                };
            }
        }).mount('#app');
</script>

<style scoped>
    body {
            font-family: 'Arial', sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 0;
            background-color: #f5f5f5;
            color: #333;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background-color: white;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
            border-radius: 8px;
            margin-top: 30px;
            margin-bottom: 30px;
        }
        header {
            text-align: center;
            padding: 20px 0;
            border-bottom: 1px solid #eee;
        }
        h1 {
            color: #2c3e50;
            margin-bottom: 10px;
        }
        .profile-img {
            width: 150px;
            height: 150px;
            border-radius: 50%;
            object-fit: cover;
            border: 3px solid #42b983;
            margin-bottom: 20px;
        }
        .section {
            margin: 30px 0;
        }
        h2 {
            color: #42b983;
            border-bottom: 2px solid #42b983;
            padding-bottom: 5px;
            display: inline-block;
        }
        .skills-list {
            display: flex;
            flex-wrap: wrap;
            gap: 10px;
        }
        .skill-tag {
            background-color: #42b983;
            color: white;
            padding: 5px 10px;
            border-radius: 20px;
            font-size: 14px;
        }
        .contact-list {
            list-style: none;
            padding: 0;
        }
        .contact-list li {
            margin: 10px 0;
            display: flex;
            align-items: center;
        }
        .contact-list i {
            margin-right: 10px;
            color: #42b983;
            width: 20px;
            text-align: center;
        }
        footer {
            text-align: center;
            padding: 20px 0;
            border-top: 1px solid #eee;
            color: #777;
        }
        .theme-toggle {
            position: fixed;
            top: 20px;
            right: 20px;
            background: #42b983;
            color: white;
            border: none;
            padding: 8px 15px;
            border-radius: 20px;
            cursor: pointer;
            font-weight: bold;
        }
        .dark-mode {
            background-color: #1a1a1a;
            color: #f0f0f0;
        }
        .dark-mode .container {
            background-color: #2d2d2d;
            color: #f0f0f0;
            box-shadow: 0 0 10px rgba(255,255,255,0.1);
        }
        .dark-mode h1, .dark-mode h2 {
            color: #42b983;
        }
        .dark-mode header, .dark-mode footer {
            border-color: #444;
        }
</style>