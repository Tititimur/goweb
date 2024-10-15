var page_id = 0;
var mouse_pos = {};
var anime_interval = 3000;    //  Интервал плеера

var lik_time = Math.floor(Date.now()/1000);
var shift_minute = (new Date()).getTimezoneOffset();
var tick_delta = 0;
var tick_second = 0;

function start_run(id) {
    page_id = Number(id);
    jQuery('body').mousemove(function(e){
        mouse_pos = getPosition(e);
    });
    setInterval(function(){ step_run(); }, 100);
}

function getPosition(e){
	var x = y = 0;
	if (!e) {
		var e = window.event;
	}
	if (e.pageX || e.pageY){
		x = e.pageX;
		y = e.pageY;
	} else if (e.clientX || e.clientY){
		x = e.clientX + document.body.scrollLeft + document.documentElement.scrollLeft;
		y = e.clientY + document.body.scrollTop + document.documentElement.scrollTop;
	}
	return {left: x, top: y};
}

function step_run() {
    var now = Date.now();
    refresh_anime();
}

function refresh_anime() {
    jQuery(".anime").each(function(idx, dom) {
        var elm = jQuery(dom);
        var nextAt = elm.attr("next");
        var orgSrc = elm.attr("org");
        if (!orgSrc) {
            orgSrc = elm.attr("src");
            elm.attr("org", orgSrc);
            nextAt = Date.now();
            elm.attr("next", nextAt);
            elm.on('load', function(){
                nextAt = Date.now();
                elm.attr("next", nextAt);
            });
            elm.on('change', function(){
                nextAt = Date.now();
                elm.attr("next", nextAt);
            });
        }
        if (nextAt <= Date.now()) {
            nextAt = Date.now()+anime_interval;
            elm.attr("next", nextAt);
            elm.attr("src", orgSrc + `&ra=${Math.random()}`);
        }
    });
}

