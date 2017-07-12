import MobileDetect from 'mobile-detect';
import platform from 'platform';

class Complaint {
    constructor(options) {
        this.keys = {}; // TODO: Make private modifying the descriptor
        this.versions = {}; // TODO: Make private modifying the descriptor
        this.screen = {}; // TODO: Make private modifying the descriptor

        this.keys.app = !isValidKey(options.app) || options.app;
        this.keys.platform = !isValidKey(options.platform) || options.platform;
        this.keys.range = !isValidKey(options.range) || options.range;

        this.versions.app = options.appVersion; // TODO: Validate (maybe in a proxy?)
        this.mobileDetect = MobileDetect(options.userAgent || window.navigator.userAgent); // TODO: Validate
        this.platform = platform.parse(options.userAgent || window.navigator.userAgent);
    }

    get isMobile() { // TODO: Consider converting into proxy
        const mobile = mobileDetect.mobile();

        return mobile != null;
    }

    get app() {
        return {
            key: this.keys.app,
            version: this.versions.app
        };
    }

    get platform() {
        return {
            key: this.keys.platform,
            version: platform.os.split(' ').pop()
        };
    }

    get user() {
        const result = {};

        if (this.user.name) result.name = this.user.name
        if (this.user.email) result.email = this.result.email
        if (this.user.mibaId) result.mibaId = this.user.mibaId

        if (!result.email && !result.mibaId) throw new Error({
            name: 'RatingError',
            message: 'User has no email/mibaId set'
        });

        return result;
    }

    get device() { // TODO: Consider converting into proxy
        return {
            name: this.isMobile ? platform.product : 'Desktop',
            brand: platform.manufacturer,
            screen: this.screen
        };
    }

    get screen() { // TODO: Consider converting into proxy
        const mobile = mobileDetect.mobile();

        return {
            width: self.screen.width || window.screen.width,
            height: self.screen.height || window.screen.height
        };
    }

    get browser() { // TODO: Consider converting into proxy
        return {
            name: platform.name,
            version: platform.version
        };
    }

    set token(value) { // TODO: Validate / consider converting into proxy
        this.token = value;
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

    set user(value) { // TODO: Validate keys / consider converting into proxy
        this.user = value;
    }

    set screen(value) { // TODO: Validate keys / consider converting into proxy
        this.screen = value;
    }

    build() {
        const result = {
            rating: this.rating,
            description: this.comment,
            comment: this.comment,
            range: this.keys.range,
            app: this.app,
            platform: this.platform
        };

        if (this.user) result.user = this.user;
        if (this.device) result.device = this.device;
        if (this.browser) result.browser = this.browser;

        return result;
    }

    send() { // TODO: Return success/error codes
        const complaint = this.build();

    }

    isValidKey(key) { // TODO: Consider converting into proxy (for each key)
        if (key.length != 32) throw new Error({ name: 'RatingError', message: 'Invalid key' });

        return true;
    }
}