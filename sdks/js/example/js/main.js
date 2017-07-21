$('#result').hide();

$(document).ready(function () {
    $("#form").submit(function (e) {
        e.preventDefault();

        const sdk = new Ratings({
            api: 'http://127.0.0.1:5000',
            token: 'e10adc3949ba59abbe56e057f20f883e',
            app: 'e10adc3949ba59abbe56e057f20f883e',
            version: '3.0',
            range: 'e10adc3949ba59abbe56e057f20f883e',
            platform: 'e10adc3949ba59abbe56e057f20f883e'
        });

        sdk.create({
            rating: parseInt($('#rating').val()),
            description: $('#description').val(),
            comment: $('#comment').val()
        })
        .then((response) => {
            $('#result').show();

            response.json().then((json) => {
                $('#result-code span').text(json.meta.code);
                $('#result-message span').text(json.meta.message);

                if (json.errors) {
                    $('#result-errors').show();
                    $('#result-errors span').text(json.errors.join('\n'));
                }
                else $('#result-errors').hide();
            });
        })
        .catch((error) => {
            console.error(error);
        });

        return false;
    });
});