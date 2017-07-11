import MobileDetect from 'mobile-detect';

class Complaint {
    constructor(options) {
        this.keys = {};

        this.keys.app = !isValidKey(options.app) || options.app;
        this.keys.platform = !isValidKey(options.platform) || options.platform;
        this.keys.range = !isValidKey(options.range) || options.range;

        this.md = new MobileDetect(window.navigator.userAgent);
    }

    get isMobile() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile != null;
    }

    get app() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    get platform() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    get user() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    get device() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    get screen() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    get browser() { // TODO: Consider converting into proxy
        let mobile = this.md.mobile();

        return mobile || null;
    }

    set rating(value) { // TODO: Validate (is int?) / consider converting into proxy
        if (value >= -127 || value <= 127) this.rating = value;
    }

    set description(value) { // TODO: Validate / consider converting into proxy
        this.description = value;
    }

    set comment(value) { // TODO: Validate / consider converting into proxy
        this.comment = value;
    }

    set user(value) { // TODO: Validate / consider converting into proxy
        this.user = value;
    }

    build() {

    }

    send() {

    }

    isValidKey(key) { // TODO: Consider converting into proxy (for each key)
        if (key.length != 32) throw new Error({ name: 'RatingError', message: 'Invalid key' });

        return true;
    }
}