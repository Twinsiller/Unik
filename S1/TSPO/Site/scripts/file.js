document.addEventListener('DOMContentLoaded', function() {
    const postMini = document.querySelector('.post-mini');
    const briefInfo = postMini.querySelector('.brief-information');

    postMini.addEventListener('mouseover', function() {
        briefInfo.style.display = 'flex';
        briefInfo.style.animation = 'moveUp 2s forwards';
    });

    postMini.addEventListener('mouseout', function() {
        briefInfo.style.display = 'none';
    });
});


// document.addEventListener('DOMContentLoaded', function() { 
//     const postMini = document.querySelector('.post-mini'); 
//     const briefInfo = postMini.querySelector('.brief-information'); 
//     postMini.addEventListener('mouseover', function() { 
//         briefInfo.style.display = 'flex'; 
//     }); 
//     postMini.addEventListener('mouseout', function(){ 
//         briefInfo.style.display = 'none'; 
//     }); 
// });

//alert('Hello, World!');