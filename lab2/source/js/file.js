$(function() {

const bnt = document.querySelector('.btn-toggle');

bnt.addEventListener('click', function() {
    document.body.classList.toggle('dark-theme');
})


/* Slider слайдер цитат*/
$("[data-slider]").slick({
    infinite: true,
    fade: false,
    slidesToShow: 1,
    slidesToScroll: 1
});

 /* Collapse выпадающий список*/
 $("[data-collapse]").on("click", function(event) {
    event.preventDefault();

    var $this = $(this),
        blockId = $this.data('collapse');

    $this.toggleClass("active-item");
    // $(blockId).sliderToggle();
    // 
});

});