let search = document.querySelector('.search-box');

document.querySelector('#search-icon').onclick = () => {
    search.classList.toggle('active');
    menu.classList.remove('active');
}

let menu = document.querySelector('.navbar');

document.querySelector('#menu-icon').onclick = () => {
    menu.classList.toggle('active');
    search.classList.remove('active');
}
window.onscroll = () => {
    menu.classList.remove('active');
    search.classList.remove('active');
}
// Header
let header = document.querySelector('header');

window.addEventListener('scroll' , () => {
    header.classList.toggle('shadow', window.scrollY > 0);
});

document.getElementById('search-input').addEventListener('input', function () {
    const query = this.value.toLowerCase(); // Get the search query and convert to lowercase
    const sections = document.querySelectorAll('.searchable'); // Select all searchable sections

    sections.forEach(section => {
        const sectionText = section.textContent.toLowerCase(); // Get the text content of the section
        if (sectionText.includes(query)) {
            section.style.display = 'block'; // Show the section if it matches the query
        } else {
            section.style.display = 'none'; // Hide the section if it doesn't match
        }
    });
});