function login(email, password) {
  fetch('/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  })
    .then(res => res.json())
    .then(data => {
      localStorage.setItem('token', data.token);
      window.location.href = 'catalog.html';
    });
}

document.getElementById('loginForm')?.addEventListener('submit', e => {
  e.preventDefault();
  const email = e.target[0].value;
  const password = e.target[1].value;
  login(email, password);
});