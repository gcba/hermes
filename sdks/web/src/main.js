'use strict';

import 'whatwg-fetch'

import MobileDetect from 'mobile-detect';
import Promise from 'promise-polyfill';
import platform from 'platform';

if (!window.Promise) {
  window.Promise = Promise;
}

const isString = (thing) => {
    return typeof thing === 'string' || thing instanceof String;
}

const isBool = (thing) => {
    return typeof (thing) === "boolean";
}

class Complaint {
    constructor(options) {
        // Should fail when:
        // - Required things are missing
        // - Things are invalid: format, type

        this.keys = {}; // TODO: Make private modifying the descriptor
        this.versions = {}; // TODO: Make private modifying the descriptor
        this.screen = {}; // TODO: Make private modifying the descriptor

        this.keys.app = options.app; // TODO: Validate (maybe in a proxy?) Required
        this.keys.platform = options.platform; // TODO: Validate (maybe in a proxy?) Required
        this.keys.range = options.range; // TODO: Validate (maybe in a proxy?) Required

        this.versions.app = options.appVersion; // TODO: Validate (maybe in a proxy?) Required

        this.url = options.api; // TODO: Validate (maybe in a proxy?) Required
        this.token = options.token; // TODO: Validate (maybe in a proxy?) Required
        this._userAgent = options.userAgent; // TODO: Make private
        this._isMobile = options.isMobile; // TODO: Validate (maybe in a proxy?) Make private

        this.mobileDetect = new MobileDetect(this.userAgent || window.navigator.userAgent); // TODO: Validate
        this.platform = platform.parse(this.userAgent || window.navigator.userAgent); // TODO: Validate
    }

    get isMobile() {
        return (this._isMobile === undefined || this._isMobile === null) ?
            mobileDetect.mobile() !== null :
            this._isMobile;
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

    get device() {
        const result = {
            name: this.isMobile ? platform.product : 'Desktop',
            screen: this.screen
        };

        if (this.isMobile && platform.manufacturer) result.brand = platform.manufacturer;

        return result;
    }

    get screen() {
        const mobile = mobileDetect.mobile();

        return {
            width: self.screen.width || window.screen.width,
            height: self.screen.height || window.screen.height
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

    get browser() {
        return {
            name: platform.name,
            version: platform.version
        };
    }

    set url(value) {
        const urlRegex = new RegExp(/^(?:(?:https?|ftp):\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,}))\.?)(?::\d{2,5})?(?:[/?#]\S*)?$/i);

        if (isString(value) && urlRegex.test(trim(value))) this.url = value;
        else throw new Error({ name: 'RatingError', message: 'Invalid URL' });
    }

    set token(value) {
        if (isString(value)) this.token = trim(value);
    }

    set _isMobile(value) {
        return isBool(value);
    }

    set _userAgent(value) {
        if (isString(value)) this._userAgent = trim(value);
    }

    set user(value) { // TODO: Validate keys / consider converting into proxy
        this.user = value;
    }

    set screen(value) { // TODO: Validate keys / consider converting into proxy
        this.screen = value;
    }

    create(data) {
        const complaint = {
            rating: data.rating, // TODO: Validate (is int?) / consider converting into proxy
            description: data.description, // TODO: Validate keys / consider converting into proxy
            comment: data.comment, // TODO: Validate keys / consider converting into proxy
            range: this.keys.range,
            app: this.app,
            platform: this.platform,
            device: this.device
        };

        if (this.user) complaint.user = this.user;
        if (this.browser) complaint.browser = this.browser;

        return this.send(complaint); //  TODO: Return promise
    }

    send(complaint) { // TODO: make private
        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=UTF-8',
                'Accept': 'application/json',
                'Accept-Charset': 'utf-8'
            },
            body: JSON.stringify(complaint)
        };

        return fetch(this.url, options);
    }

    isValidKey(key) { // TODO: Consider converting into proxy (for each key)
        if (key.length != 32) throw new Error({ name: 'RatingError', message: 'Invalid key' });

        return true;
    }
}

export default Complaint;