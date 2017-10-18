 require.config(
    {
        paths: {
                'jquery': '/static/bootstrap/js/jquery.min',
                'bootstrap':'/static/bootstrap/js/bootstrap.min'
        },

        shim: {
            jquery:{exports: 'jquery'},
            bootstrap: {deps: ['jquery']}
        }
    }
);
require(['jquery','bootstrap'],function ($) {
    $(document).on('click','#contentBtn',function(){
        $('#messagebox').html('You have access Jquery by using require()');
    });
});