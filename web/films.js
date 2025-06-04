document.addEventListener('DOMContentLoaded', () => {
  const userId = localStorage.getItem('user_id');
  if (!userId) {
    window.location.href = 'auth.html'; // если не авторизован — редирект
    return;
  }

  initFilmForm();
  initFilmFilter();
});

function initFilmForm() {
  const addFilmForm = document.getElementById('addFilmForm');
  const addFilmMessage = document.getElementById('addFilmMessage');

  addFilmForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    console.log('Форма отправлена'); // Лог для проверки срабатывания события

    const title = document.getElementById('filmTitle').value.trim();
    const genre = document.getElementById('filmGenre').value;
    const yearInput = document.getElementById('filmYear').value.trim();
    const year = parseInt(yearInput, 10);

    if (!title || !genre || isNaN(year)) {
      showMessage(addFilmMessage, 'Заполните все поля корректно', 'error');
      return;
    }

    if (year < 1900 || year > 2025){
      showMessage(addFilmMessage, 'Введите год в диапазоне от 1900 до 2025', 'error')
      return;
    }

    showMessage(addFilmMessage, 'Добавляем фильм...', 'loading');

    try {
      console.log('Отправляем фильм:', { title, genre, year });

      const response = await fetch('/models/films', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, genre, year })
      });

      if (response.ok) {
        const data = await response.json();
        showMessage(addFilmMessage, `Фильм добавлен!`, 'success');
        addFilmForm.reset();

        const selectedGenre = document.getElementById('genreSelect').value;
        if (selectedGenre && selectedGenre === genre) {
          document.getElementById('showFilmsBtn').click();
        }
      } else {
        const error = await response.text();
        throw new Error(error);
      }
    } catch (error) {
      showMessage(addFilmMessage, 'Ошибка: ' + error.message, 'error');
    }
  });
}

function initFilmFilter() {
  const showFilmsBtn = document.getElementById('showFilmsBtn');
  const filmsList = document.getElementById('filmsList');

  showFilmsBtn.addEventListener('click', async () => {
    const genre = document.getElementById('genreSelect').value;

    if (!genre) {
      alert('Выберите жанр');
      return;
    }

    filmsList.innerHTML = '<p class="empty-message">Загрузка...</p>';

    try {
      const response = await fetch(`/models/films/bygenre?genre=${encodeURIComponent(genre)}`);

      if (!response.ok) {
        throw new Error('Ошибка загрузки');
      }

      const films = await response.json();

      if (films.length === 0) {
        filmsList.innerHTML = '<p class="empty-message">Фильмы не найдены</p>';
        return;
      }

      filmsList.innerHTML = films.map(film => `
        <div class="film-item">
          <div class="film-title">${film.title}</div>
          <div class="film-genre">Жанр: ${film.genre}</div>
          <div class="film-year">Год выпуска: ${film.year}</div>
        </div>
      `).join('');
    } catch (error) {
      filmsList.innerHTML = '<p class="empty-message">Ошибка загрузки: ' + error.message + '</p>';
    }
  });
}

function showMessage(element, text, type) {
  element.textContent = text;
  element.className = 'message';
  element.classList.add(type);
}
