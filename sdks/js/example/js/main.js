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

        sdk.user = {
            name: $('#name').val(),
            email: $('#email').val()
        }

        sdk.create({
            rating: parseInt($('#rating').val()),
            description: $('#description').val(),
            comment: $('#comment').val()
        })
        .then((response) => {
            $('#result').show();

            $('#result-code span').text(response.meta.code);
            $('#result-message span').text(response.meta.message);

            if (response.errors) {
                $('#result-errors').show();
                $('#result-errors span').text(response.errors.join('\n'));
            }
            else $('#result-errors').hide();
        })
        .catch((errors) => {
            console.error(errors);
        });

        return false;
    });
});