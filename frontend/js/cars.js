async function fetchCars(filters = {}) {
  let url = '/api/cars';

  const params = new URLSearchParams();
  for (let key in filters) {
    if (filters[key]) params.append(key, filters[key]);
  }

  if (params.toString()) url += '?' + params;

  const response = await fetch(url, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`,
    },
  });

  return await response.json();
}

function renderCars(cars) {
  const container = document.getElementById('carList');
  container.innerHTML = '';
  cars.forEach(car => {
    const div = document.createElement('div');
    div.className = 'car-card';
    div.innerHTML = `<h3>${getModelName(car.model)}</h3>
                     <p>Цвет: ${car.color}</p>
                     <p>Год: ${car.year}</p>
                     <p>Цена: ₽${car.price}</p>`;
    container.appendChild(div);
  });
}