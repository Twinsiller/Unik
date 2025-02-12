document.addEventListener('DOMContentLoaded', () => {
    const taskInput = document.getElementById('taskInput');
    const addButton = document.getElementById('addButton');
    const taskList = document.getElementById('taskList');
    
    let clickTimeout = null; // Делаем переменную глобальной для всех listItem

    addButton.addEventListener('click', () => {
        const taskText = taskInput.value.trim();
        if (taskText === '') return;

        const listItem = document.createElement('li');
        listItem.textContent = taskText;

        listItem.addEventListener('click', (event) => {
            if (clickTimeout) {
                clearTimeout(clickTimeout);
                clickTimeout = null;
                return; // Остановить одиночный клик, если произошёл двойной
            }
    
            clickTimeout = setTimeout(() => {
                console.log('Одиночный клик:', event.target.textContent);
                taskList.removeChild(listItem);
                clickTimeout = null;
            }, 300); // Задержка для определения двойного клика
        });

        listItem.addEventListener('dblclick', (event) => {
            clearTimeout(clickTimeout); // Остановить таймер одиночного клика
            clickTimeout = null;

            console.log('Двойной клик:', event.target.textContent);
            const newText = prompt('Введите новый текст:', event.target.textContent);
            if (newText) event.target.textContent = newText;
        });

        taskList.appendChild(listItem);
        taskInput.value = '';
    });
});
