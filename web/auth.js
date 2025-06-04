console.log('auth.js загружен');

document.addEventListener('DOMContentLoaded', () => {
  const userId = localStorage.getItem('user_id');

  if (!userId) {
    document.getElementById('loginFormContainer').style.display = 'none';
    document.getElementById('registerFormContainer').style.display = 'block';
    document.querySelectorAll('.auth-tab').forEach(tab => tab.classList.remove('active'));
    document.querySelector('.auth-tab[data-tab="register"]').classList.add('active');
  }

  initAuthTabs();
  initLoginForm();
  initRegisterForm();
});

function initAuthTabs() {
  document.querySelectorAll('.auth-tab').forEach(tab => {
    tab.addEventListener('click', () => {
      document.querySelectorAll('.auth-tab').forEach(t => t.classList.remove('active'));
      tab.classList.add('active');

      const tabName = tab.dataset.tab;
      document.getElementById('loginFormContainer').style.display =
        tabName === 'login' ? 'block' : 'none';
      document.getElementById('registerFormContainer').style.display =
        tabName === 'register' ? 'block' : 'none';
    });
  });
}

function initLoginForm() {
  const loginForm = document.getElementById('loginForm');
  const loginMessage = document.getElementById('loginMessage');

  loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const identifier = document.getElementById('loginInput').value.trim();

    if (!identifier) {
      showMessage(loginMessage, 'Введите email или ник', 'error');
      return;
    }

    showMessage(loginMessage, 'Входим...', 'loading');

    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ identifier })
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('user_id', data.user_id);
        showMessage(loginMessage, 'Успешный вход!', 'success');
        window.location.href = 'films.html'; // редирект
      } else {
        const error = await response.text();
        throw new Error(error);
      }
    } catch (error) {
      showMessage(loginMessage, 'Ошибка входа: ' + error.message, 'error');
    }
  });
}

function initRegisterForm() {
  const registerForm = document.getElementById('registerForm');
  const registerMessage = document.getElementById('registerMessage');

  registerForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const user = {
      name: document.getElementById('name').value.trim(),
      email: document.getElementById('email').value.trim(),
      nickname: document.getElementById('nickname').value.trim()
    };

    if (!user.name || !user.email || !user.nickname) {
      showMessage(registerMessage, 'Заполните все поля', 'error');
      return;
    }

    showMessage(registerMessage, 'Регистрируем...', 'loading');

    try {
      const response = await fetch('/models/users', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('user_id', data.user_id);
        showMessage(registerMessage, 'Регистрация успешна!', 'success');

        window.location.href = '/films.html';
      } else {
        const error = await response.text();
        throw new Error(error);
      }
    } catch (error) {
      showMessage(registerMessage, 'Ошибка регистрации: ' + error.message, 'error');
    }
  });
}

function showMessage(element, text, type) {
  element.textContent = text;
  element.className = 'message';
  element.classList.add(type);
}
