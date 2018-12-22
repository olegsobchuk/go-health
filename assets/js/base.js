$(document).ready(function(){
  // customize link methods
  $('a').click(function(event) {
    method = $(this).data('method')
    url = $(this).attr('href')
    methods = ['post', 'put', 'patch', 'delete']
    if (method && $.inArray(method.toLowerCase(), methods) != -1) {
      event.preventDefault()
      form = $('<form/>').attr('action', url).attr( 'method', method)
      console.log(form);
      form.appendTo('body').submit()
    }
  })
});
