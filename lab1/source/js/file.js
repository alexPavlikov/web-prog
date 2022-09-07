$(function() {

    var header = $("#header"),
    introH = $("#intro").innerHeight(),
    scrollOffSet =  $(window).scrollTop();



    /* Header fixed мини шапка*/
    checkScroll(scrollOffSet);

    $(window).on("scroll", function() {

        scrollOffSet = $(this).scrollTop();
        
        checkScroll(scrollOffSet);

    });


    function checkScroll(scrollOffSet) {
        if( scrollOffSet >= introH ) {
            header.addClass("fixed");
        } else {
            header.removeClass("fixed");
        }
    };


    /* Smooth scroll плавный скрол*/
    $("[data-scroll]").on("click", function(event) {
        event.preventDefault();

        var $this = $(this),
            blockId = $this.data('scroll'),
            blockOffset = $(blockId).offset().top;

        $("#nav a").removeClass("active");
        $this.addClass("active");

        $("html, body").animate({
            scrollTop:  blockOffset
        }, 700);
    });



});