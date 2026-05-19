const button = document.getElementById('fetch');

button.addEventListener('click', (event)=>{
    button.textContent = "Fetching";
    button.classList.add('active');
});