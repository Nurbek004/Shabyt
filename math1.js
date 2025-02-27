function openVideo(url) {
    // Open the YouTube video in a new tab
    window.open(url, '_blank');
}
// Filter lessons based on search input
function filterLessons() {
    const searchInput = document.getElementById('searchInput').value.toLowerCase();
    const lessonCards = document.querySelectorAll('.lesson-card');

    lessonCards.forEach(card => {
        const title = card.querySelector('h3').innerText.toLowerCase();
        if (title.includes(searchInput)) {
            card.style.display = 'block';
        } else {
            card.style.display = 'none';
        }
    });
}