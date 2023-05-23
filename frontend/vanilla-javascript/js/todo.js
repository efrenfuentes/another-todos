const baseUrl = 'http://192.168.0.149:4000/v1';

window.onload = function() {
  getTodos();
};

function getTodos() {
  const url_parts = document.URL.split('/')
  const filter = url_parts[url_parts.length - 1];

  switch(filter) {
    case 'active':
      document.getElementById('filter_active').classList.add('text-sky-600');
      document.getElementById('filter_completed').classList.remove('text-sky-600');
      document.getElementById('filter_all').classList.remove('text-sky-600');
      break;
    case 'completed':
      document.getElementById('filter_active').classList.remove('text-sky-600');
      document.getElementById('filter_completed').classList.add('text-sky-600');
      document.getElementById('filter_all').classList.remove('text-sky-600');
      break;
    default:
      document.getElementById('filter_active').classList.remove('text-sky-600');
      document.getElementById('filter_completed').classList.remove('text-sky-600');
      document.getElementById('filter_all').classList.add('text-sky-600');
  }

  const url = `${baseUrl}/todos`;
  fetch(url).then((response) => {
    return response.json();
  } ).then((data) => {
    todoList = document.getElementById('todo-list');
    todoList.innerHTML = '';

    for(let i=0; i < data.todos.length; i++) {
      let checked = '';
      if(data.todos[i].completed) {
        checked = 'checked';
      }

      if (filter == 'active' && data.todos[i].completed) {
        continue;
      }

      if (filter == 'completed' && !data.todos[i].completed) {
        continue;
      }

      todoList.innerHTML += `
      <div id="container-todo-${data.todos[i].id}" class="relative flex items-start py-4 text-xl">
        <div class="min-w-0 flex-1">
          <label for="todo-${data.todos[i].id}" class="select-none font-medium ${data.todos[i].completed ? 'line-through text-gray-300' : 'text-gray-700'}">${data.todos[i].title}</label>
        </div>
        <div class="ml-3 flex h-5 items-center">
          <input id="todo-${data.todos[i].id}" name="todo-${data.todos[i].id}" type="checkbox" ${checked} onclick="handleCompletedClick(this);" class="hidden h-4 w-4 rounded border-gray-300 text-sky-600 focus:ring-sky-500">
          <button id="delete-${data.todos[i].id}" type="button" onclick="handleDeleteClick(this);" class="ml-2 px-2.5 py-1.5 font-semibold text-red-500  hover:text-red-700">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" />
            </svg>
          </button>
        </div>
      </div>`
    }
  });
}

function handleCompletedClick(completeControl) {
  const todoId = completeControl.id.replace('todo-', '');
  const url = `${baseUrl}/todos/${todoId}`;

  fetch(url, {
    method: 'PUT',
    headers:{
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      completed: completeControl.checked
    }),
  }).then((response) => {
    return response.json();
  }).then((data) => {
    getTodos();
  });
}

function handleDeleteClick(deleteControl) {
  result = confirm("Are you sure you want to delete this todo?");
  if(!result) {
    return;
  }

  const todoId = deleteControl.id.replace('delete-', '');
  const url = `${baseUrl}/todos/${todoId}`;

  fetch(url, {
    method: 'DELETE',
    headers:{
      'Content-Type':'application/json'
    },
  }).then((response) => {
    return response.json();
  }
  ).then((data) => {
    getTodos();
  });
}

function handleNewClick() {
  const title = document.getElementById('new-todo').value;
  const url = `${baseUrl}/todos`;

  fetch(url, {
    method: 'POST',
    headers:{
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      title: title,
      completed: false
    }),
  }).then((response) => {
    return response.json();
  }).then((data) => {
    getTodos();
  });
}

function handleFilterClick(filterLink) {
  const filter = filterLink.href;

  window.history.pushState({}, '', filter);

  getTodos();
}
