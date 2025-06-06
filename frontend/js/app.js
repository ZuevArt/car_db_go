// app.js — Инициализация страницы

document.addEventListener("DOMContentLoaded", function () {
  // Подгрузка header и footer
  if (document.getElementById("header")) {
    utils.loadComponent("header", "../components/header.html");
  }
  if (document.getElementById("footer")) {
    utils.loadComponent("footer", "../components/footer.html");
  }

  // Автоматическая загрузка автомобилей на главной странице
  if (window.location.pathname.endsWith("catalog.html")) {
    cars.fetchCars().then(cars.renderCars);
  }
});