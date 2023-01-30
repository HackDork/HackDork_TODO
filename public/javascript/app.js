// Add click event to fact to show answer
(() => {
    const answerWrapper = document.querySelectorAll('.assignee-wrapper');
    const toggleBtns = document.querySelectorAll('.assignee-toggle')

    for (const ans of answerWrapper) {
        ans.style.display = 'none';
    }

    for (const btn of toggleBtns) {
        btn.addEventListener('click', (e) => {
            const answer = e.target.parentElement.nextElementSibling;
            answer.style.display = answer.style.display === 'none' ? 'block' : 'none';
        } );
    }
})()