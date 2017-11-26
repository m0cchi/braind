import 'jquery';
// import $ from 'jquery';
import 'tether';
import 'bootstrap';



function cleanArticle(){
    $('#articleTitle').val('');
    $('#articleBody').val('');
}

function postArticle(){
    var postArticleButton = $('#postArticleButton');
    postArticleButton.attr('disabled', true);
    var articleForm = $('#articleForm');
    var fd = new FormData(articleForm[0]);
    
    $.ajax({
        type: 'POST',
        url: articleForm.attr('action'),
        data: fd,
        processData: false,
        contentType: false,
        dataType: 'json'
    }).done(function(data) {
        postArticleButton.attr('disabled', false);
        cleanArticle();
    }).fail(function(data) {
        alert('送信に失敗しました。');
        postArticleButton.attr('disabled', false);
    });
}

function init() {
    var postArticleButton = $('#postArticleButton');
    postArticleButton[0].onclick = postArticle;
    var cleanArticleButton = $('#cleanArticleButton');
    console.log(cleanArticleButton);
    cleanArticleButton[0].onclick = cleanArticle;
    //    cleanArticleButton.addEventListener('click', () => alert('hhhh'), false);
    //    postArticleButton.addEventListener('click', () => alert('hhhh'), false);
    console.log('load header');
    document.body.jquery = $;
}

document.addEventListener("DOMContentLoaded", init, false);
