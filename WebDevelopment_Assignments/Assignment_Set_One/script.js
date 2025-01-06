
document.addEventListener('DOMContentLoaded', () => {
    const taskInput = document.getElementById('task-input');
    const addTaskBtn = document.getElementById('add-task-btn');
    const taskList = document.getElementById('task-list');
    const pendingCountEl = document.getElementById('pending-count');
  
    let tasks = JSON.parse(localStorage.getItem('tasks')) || [];
  
    const renderTasks = () => {
      taskList.innerHTML = '';
      tasks.forEach((task, index) => {
        const taskItem = document.createElement('li');
        taskItem.className = task.completed ? 'completed' : '';
        taskItem.innerHTML = `
          <span>${task.name}</span>
          <div class="actions">
            <button class="edit" onclick="editTask(${index})">Edit</button>
            <button class="delete" onclick="deleteTask(${index})">Delete</button>
            <button class="complete" onclick="toggleComplete(${index})">
              ${task.completed ? 'Undo' : 'Complete'}
            </button>
          </div>
        `;
        taskList.appendChild(taskItem);
      });
      updatePendingCount();
    };
  
    const updatePendingCount = () => {
      const pendingTasks = tasks.filter(task => !task.completed).length;
      pendingCountEl.textContent = `Pending Tasks: ${pendingTasks}`;
    };
  
    const saveTasks = () => {
      localStorage.setItem('tasks', JSON.stringify(tasks));
    };
  
    addTaskBtn.addEventListener('click', () => {
      const taskName = taskInput.value.trim();
      if (taskName) {
        tasks.push({ name: taskName, completed: false });
        taskInput.value = '';
        saveTasks();
        renderTasks();
      }
    });
  
    window.editTask = (index) => {
      const newName = prompt('Edit Task', tasks[index].name);
      if (newName) {
        tasks[index].name = newName.trim();
        saveTasks();
        renderTasks();
      }
    };
  
    window.deleteTask = (index) => {
      tasks.splice(index, 1);
      saveTasks();
      renderTasks();
    };
  
    window.toggleComplete = (index) => {
      tasks[index].completed = !tasks[index].completed;
      saveTasks();
      renderTasks();
    };
  
    
    renderTasks();
  });
  