$(document).ready(function(){

  // customize link methods
  $('a').click(function(event) {
    method = $(this).data('method')
    url = $(this).attr('href')
    methods = ['post', 'put', 'patch', 'delete']
    if (method && $.inArray(method.toLowerCase(), methods) != -1) {
      event.preventDefault()
      $.ajax({
        url: url,
        dataType: 'html',
        method: method
      })
    }
  })
});
