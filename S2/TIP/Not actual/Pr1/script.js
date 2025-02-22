document.addEventListener('DOMContentLoaded', () => {
    const taskInput = document.getElementById('taskInput');
    const addButton = document.getElementById('addButton');
    const taskList = document.getElementById('taskList');

    let clickTimeout = null;

    // ✅ Функция для сохранения списка задач в localStorage
    function saveTasks() {
        const tasks = [];
        document.querySelectorAll('#taskList li').forEach(li => tasks.push(li.textContent));
        localStorage.setItem('tasks', JSON.stringify(tasks));
    }

    // ✅ Функция для загрузки задач из localStorage
    function loadTasks() {
        const tasks = JSON.parse(localStorage.getItem('tasks')) || [];
        tasks.forEach(taskText => addTask(taskText));
    }

    // ✅ Функция для создания и добавления новой задачи
    function addTask(taskText, isChecked = false) {
        const divLine = document.createElement('div');
        
        const listItem = document.createElement('li');

        listItem.addEventListener('click', (event) => {
            if (clickTimeout) {
                clearTimeout(clickTimeout);
                clickTimeout = null;
                return; // Остановить одиночный клик, если произошёл двойной
            }

            clickTimeout = setTimeout(() => {
                console.log('Одиночный клик:', event.target.textSpan);
                taskList.removeChild(listItem);
                saveTasks(); // Сохранить после удаления
                clickTimeout = null;
            }, 300);
        });

        listItem.addEventListener('dblclick', (event) => {
            clearTimeout(clickTimeout);
            clickTimeout = null;

            console.log('Двойной клик:', event.target.textSpan);
            const newText = prompt('Введите новый текст:', event.target.textSpan);
            if (newText) {
                event.target.textSpan = newText;
                saveTasks(); // Сохранить после изменения
            }
        });

        const radio = document.createElement('input');
        radio.type = 'radio';
        radio.checked = isChecked;
        radio.addEventListener('change', saveTasks);

        const textSpan = document.createElement('span');
        textSpan.textContent = taskText;

        listItem.appendChild(radio);
        listItem.appendChild(textSpan);

        taskList.appendChild(listItem);
        taskList.appendChild(buttonDiv);
        saveTasks(); // Сохранить после добавления
    }

    // ✅ Добавление новой задачи при клике на кнопку
    addButton.addEventListener('click', () => {
        const taskText = taskInput.value.trim();
        if (taskText === '') return;
        addTask(taskText);
        taskInput.value = '';
    });

    // ✅ Загрузка задач при загрузке страницы
    loadTasks();
});
