//21,94     define variable,function      jQuery = function()
//96,283    define method,attribute
//285,347   extend 
//349,817   jQuery,extend
//877,2856  Sizzle: complexed selector
//2880,3042 callbacks
//3043,3183 deferred


//8826      window.jQuery = window.$ = jQuery  

(function(){
    jQuery = function(selector,context){
        return new jQuery.fn.init(selector,context,rootjQuery);
    };



    window.jQuery = window.$ = jQuery；


})()

