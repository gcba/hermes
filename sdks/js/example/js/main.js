$(document).ready(function() {
    const sdk = new Ratings({
        api: 'http://127.0.0.1:5000',
        token: 'e10adc3949ba59abbe56e057f20f883e',
        app: 'e10adc3949ba59abbe56e057f20f883e',
        version: '3.0',
        range: 'e10adc3949ba59abbe56e057f20f883e',
        platform: 'e10adc3949ba59abbe56e057f20f883e'
    });

    sdk.create({
        rating: 5,
        description: 'Regular',
        comment: 'Lorem ipsum'
    })
    .then(() => {
        console.info('SUCCESS');
    })
    .catch((error) => {
        console.error(error);
    });
});