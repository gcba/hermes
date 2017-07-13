'use strict';

import 'whatwg-fetch'

import isMobile from 'ismobilejs';
import platform from 'platform';

const fail = (message) => {
    throw new Error({ name: 'RatingError', message: message });
}

const check = {
    isString: (thing) => {
        return typeof thing === 'string' || thing instanceof String;
    },
    isBool: (thing) => {
        return typeof (thing) === 'boolean';
    },
    isInteger: (thing) => {
        return typeof val === 'number' && isFinite(val) && Math.floor(val) === val;
    },
    isPlainObject: (thing) => {
        return typeof thing === 'object' &&
                thing !== null &&
                thing.constructor === Object &&
                thing.hasOwnProperty('isPrototypeOf') === false &&
                thing.toString() === '[object Object]';
    }
};

const validate = {
    options: (value) => {
        if (check.isPlainObject(value)) return true;

        fail('Invalid options object');
    },
    rating: (value) => {
        if (value && check.isInteger(value) && value >= -127 && value <= 127) return value;

        fail('Invalid or missing rating');
    },
    description: (value) => {
        if (check.isString(value) && value.trim().length >= 3 && value.trim().length <= 30) return value.trim();

        fail('Invalid description');
    },
    comment: (value) => {
        if (check.isString(value) && value.trim().length >= 3 && value.trim().length <= 1000) return value.trim();

        fail('Invalid comment');
    },
    key: (value, name) => {
        if (value && check.isString(value.trim()) && value.trim().length === 32) return value.trim();

        fail('Invalid or missing ' + name);
    },
    token: (value) => {
        if (value && check.isString(value) && value.trim().length > 0) return value.trim();

        fail('Invalid or missing token');
    },
    url: (value) => {
        const url = new RegExp(/^(?:(?:https?|ftp):\/\/)(?:\S+(?::\S*)?@)?(?:(?!(?:10|127)(?:\.\d{1,3}){3})(?!(?:169\.254|192\.168)(?:\.\d{1,3}){2})(?!172\.(?:1[6-9]|2\d|3[0-1])(?:\.\d{1,3}){2})(?:[1-9]\d?|1\d\d|2[01]\d|22[0-3])(?:\.(?:1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.(?:[1-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(?:(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)(?:\.(?:[a-z\u00a1-\uffff0-9]-*)*[a-z\u00a1-\uffff0-9]+)*(?:\.(?:[a-z\u00a1-\uffff]{2,}))\.?)(?::\d{2,5})?(?:[/?#]\S*)?$/i);

        if (value && check.isString(value) && value.trim().length > 0 && url.test(value.trim())) return value.trim();

        fail('Invalid or missing api');
    },
    appVersion: (value) => {
        if (value && check.isString(value) && value.trim().length >= 1 && value.trim().length <= 15)
            return value.trim();

        fail('Invalid or missing version');
    },
    isMobile: (value) => {
        if (value === undefined || value === null || check.isBool(value)) return value;

        fail('Invalid isMobile');
    },
    userAgent: (value) => {
        if (check.isString(value) && value.trim().length > 0) return value.trim();

        fail('Invalid userAgent');
    }
};

class Complaint {
    constructor(options) {
        validate.options(options);

        this._keys = {};
        this._versions = {};
        this._screen = {};

        this._keys.app = validate.key(options.app, 'app');
        this._keys.platform = validate.key(options.platform, 'platform');
        this._keys.range = validate.key(options.range, 'range');

        this._appVersion = validate.appVersion(options.version);
        this._url = validate.url(options.api);
        this._token = validate.token(options.token);
        this._userAgent = options.userAgent ? validate.userAgent(options.userAgent) : window.navigator.userAgent;
        this._isMobile = validate.isMobile(options.isMobile);

        this._platform = platform.parse(this._userAgent);
    }

    get isMobile() {
        return (this._isMobile === undefined || this._isMobile === null) ? isMobile.any() : this._isMobile;
    }

    get app() {
        return {
            key: this._keys.app,
            version: this._appVersion
        };
    }

    get platform() {
        return {
            key: this._keys.platform,
            version: this._platform.os.split(' ').pop()
        };
    }

    get device() {
        const result = {
            name: this.isMobile ? this._platform.product : 'Desktop',
            screen: this.screen
        };

        if (this.isMobile && this._platform.manufacturer) result.brand = this._platform.manufacturer;

        return result;
    }

    get screen() {
        return {
            width: self.screen.width || window.screen.width,
            height: self.screen.height || window.screen.height
        };
    }

    get user() {
        const result = {};

        if (this_.user.name) result.name = this._user.name
        if (this_.user.email) result.email = this._user.email
        if (this_.user.mibaId) result.mibaId = this._user.mibaId

        return result;
    }

    get browser() {
        return {
            name: this._platform.name,
            version: this._platform.version
        };
    }

    set user(value) {
        const email = new RegExp(/^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/);

        const isPlainObject = check.isPlainObject(value);
        const hasName = check.isString(value.name) && value.name.trim().length > 0;
        const hasEmail = check.isString(value.email) && value.email.trim().length > 0;
        const hasMibaId = check.isString(value.mibaId) && value.mibaId.trim().length > 0;
        const nameIsValid = value.name.trim().length >= 3 && value.name.trim().length <= 70;
        const emailIsValid = email.test(value.trim()) &&
            value.email.trim().length >= 3 && value.email.trim().length <= 100;
        const mibaIdIsValid = value.mibaId.trim().length >= 1;

        if (!(isPlainObject && (hasName || hasEmail || hasMibaId))) fail('User object is invalid');
        if (!((hasEmail && emailIsValid) || (hasMibaId && mibaIdIsValid))) fail('User has no valid email or mibaId');
        else this._user = value;
    }

    set screen(value) {
        const isPlainObject = check.isPlainObject(value);
        const hasValidWidth = check.isInteger(value.width) && value > 0;
        const hasValidHeight = check.isInteger(value.width) && value > 0;

        if (!(isPlainObject && hasValidWidth &&  hasValidHeight)) fail('Screen object is invalid');
        else this._screen = value;
    }

    create(data) {
        const complaint = {
            rating: validate.rating(data.rating),
            range: this._keys.range,
            app: this.app,
            platform: this.platform,
            device: this.device
        };

        if (data.description) data.description = validate.description(data.description);
        if (data.comment) data.comment = validate.comment(data.comment);
        if (this.user) complaint.user = this.user;
        if (this.browser) complaint.browser = this.browser;

        return this.send(complaint);
    }

    send(complaint) {
        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json; charset=UTF-8',
                'Accept': 'application/json',
                'Accept-Charset': 'utf-8',
                'Authorization': 'Bearer ' + this.token
            },
            body: JSON.stringify(complaint)
        };

        return fetch(this.url, options);
    }
}

export default Complaint;