/*global KeyValueApp _config*/

var KeyValueApp = window.KeyValueApp || {};

(function keyvalueScopeWrapper($) {
    var authToken;
    KeyValueApp.authToken.then(function setAuthToken(token) {
        if (token) {
            authToken = token;
        } else {
            window.location.href = '/signin.html';
        }
    }).catch(function handleTokenError(error) {
        alert(error);
        window.location.href = '/signin.html';
    });
    
    
    function generatekeyvalue() {
        $.ajax({
            method: 'POST',
            url: _config.api.invokeUrl + '/keys',
            headers: {
                Authorization: authToken
            },
            data:  $('form.form-horizontal').serialize(),
            contentType: 'application/json',
            success: getKeyValues,
            error: function ajaxError(jqXHR, textStatus, errorThrown) {
                console.error('Error calling backend: ', textStatus, ', Details: ', errorThrown);
                console.error('Response: ', jqXHR.responseText);
                alert('An error occured when inserting keyvalue in backend:\n' + jqXHR.responseText);
            }
        });
    }

    function completeRequest(result) {
    	var table = document.getElementById("keyvalue-table");
    	//or use :  var table = document.all.tableid;
    	for(var i = table.rows.length - 1; i > 0; i--)
    	{
    	    table.deleteRow(i);
    	}
    	  $.each(result, function(i, item) {
              $('<tr>').append(
                  $('<td>').text(i+1),
                  $('<td>').text(item.key),
                  $('<td>').text(item.value),
                  $('<td>').text(item.userName)
              ).appendTo('#keyvalue-table');
          });
    }
    
    
    function getKeyValues() {
        $.ajax({
            method: 'GET',
            url: _config.api.invokeUrl + '/keys',
            headers: {
                Authorization: authToken
            },
            success: completeRequest,
            error: function ajaxError(jqXHR, textStatus, errorThrown) {
                console.error('Error calling backend: ', textStatus, ', Details: ', errorThrown);
                console.error('Response: ', jqXHR.responseText);
                alert('An error occured when inserting keyvalue in backend:\n' + jqXHR.responseText);
            }
        });
    }

    // Register click handler for #request button
    $(function onDocReady() {
        $('#postRequest').click(handlePostRequestClick);
        $('#postRequest2').click(handlePostRequestClick);
        $('#getRequest').click(handleGetRequestClick);
        $('#signOut').click(function() {
            KeyValueApp.signOut();
            alert("You have been signed out.");
            window.location = "signin.html";
        });

        KeyValueApp.authToken.then(function updateAuthMessage(token) {
            if (token) {
                displayUpdate('You are authenticated. Click to see your <a href="#authTokenModal" data-toggle="modal">auth token</a>.');
                $('.authToken').text(token);
            }
        });

        if (!_config.api.invokeUrl) {
            $('#noApiMessage').show();
        }
    });


    function handlePostRequestClick(event) {
        event.preventDefault();
        generatekeyvalue();
    }
    
    function handleGetRequestClick(event) {
        event.preventDefault();
        getKeyValues();
    }
    
    
    function displayUpdate(text) {
        $('#updates').append($('<li>' + text + '</li>'));
    }

}(jQuery));
