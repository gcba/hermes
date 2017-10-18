'use strict';

import 'whatwg-fetch'

import isMobile from 'ismobilejs';
import platform from 'platform';

const fail = (message) => {
    console.error(message);
}

const check = {
    isString: (thing) => {
        return typeof thing === 'string' || thing instanceof String;
    },
    isBool: (thing) => {
        return typeof (thing) === 'boolean';
    },
    isInteger: (thing) => {
        return typeof thing === 'number' && isFinite(thing) && Math.floor(thing) === thing;
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

        fail('invalid options object');
    },
    rating: (value) => {
        if (check.isInteger(value) && value >= -127 && value <= 127) return value;

        fail('invalid rating');
    },
    description: (value) => {
        if (check.isString(value)) {
            let trimmedValue = value.trim();

            if (trimmedValue.length < 3) fail('description too short');
            if (trimmedValue.length > 30) fail('description too long');

            return trimmedValue;
        }

        fail('invalid description');
    },
    comment: (value) => {
        if (check.isString(value)) {
            let trimmedValue = value.trim();

            if (trimmedValue.length < 3) fail('comment too short');
            if (trimmedValue.length > 1000) fail('comment too long');

            return trimmedValue;
        }

        fail('invalid comment');
    },
    key: (value, name) => {
        if (value && check.isString(value.trim()) && value.trim().length === 32) return value.trim();

        fail('invalid ' + name);
    },
    token: (value) => {
        if (value && check.isString(value) && value.trim().length > 0) return value.trim();

        fail('invalid token');
    },
    url: (value) => {
        const url = new RegExp(/^(ftp|http|https):\/\/[^ "]+$/);

        if (value && check.isString(value) && url.test(value.trim())) {
            let baseUrl = value.trim();

            return baseUrl[baseUrl.length - 1] === '/' ? baseUrl + 'ratings' : baseUrl + '/ratings';
        }

        fail('invalid api');
    },
    appVersion: (value) => {
        if (check.isString(value)) {
            let trimmedValue = value.trim();

            if (trimmedValue.length < 1) fail('version too short');
            if (trimmedValue.length > 15) fail('version too long');

            return trimmedValue;
        }

        fail('invalid version');
    },
    isMobile: (value) => {
        if (value === undefined || value === null || check.isBool(value)) return value;

        fail('invalid isMobile');
    },
    userAgent: (value) => {
        if (check.isString(value) && value.trim().length > 0) return value.trim();

        fail('invalid userAgent');
    },
    name: (value) => {
        if (check.isString(value)) {
            let trimmedValue = value.trim();

            if (trimmedValue.length < 3) fail('name too short');
            if (trimmedValue.length > 70) fail('name too long');

            return trimmedValue;
        }

        fail('invalid name');
    },
    email: (value) => {
        const email = new RegExp(/^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/);

        if (check.isString(value) && email.test(value.trim())) {
            let trimmedValue = value.trim();

            if (trimmedValue.length < 3) fail('email too short');
            if (trimmedValue.length > 100) fail('email too long');

            return trimmedValue;
        }

        fail('invalid email');
    },
    mibaId: (value) => {
        if (check.isString(value) && value.length === 36) return value.trim();

        fail('invalid mibaId');
    }
};

class RatingError {
    constructor(message) {
        this.message = message;
        this.name = 'RatingError'
    }

    toString() {
        return this.name + ': ' + this.message
    }
}

class Rating {
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
        return (this._isMobile === undefined || this._isMobile === null) ? isMobile.any : this._isMobile;
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
            version: this._platform.os.version
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
        if (this._user) {
            const result = {};

            if (this._user.name) result.name = this._user.name
            if (this._user.email) result.email = this._user.email
            if (this._user.mibaId) result.mibaId = this._user.mibaId

            return result;
        }

        return;
    }

    get browser() {
        return {
            name: this._platform.name,
            version: this._platform.version
        };
    }

    set user(value) {
        const isPlainObject = check.isPlainObject(value);
        const hasName = check.isString(value.name) && value.name.trim().length > 0;
        const hasEmail = check.isString(value.email) && value.email.trim().length > 0;
        const hasMibaId = check.isString(value.mibaId) && value.mibaId.trim().length > 0;
        const name = validate.name(value.name);
        const email = validate.email(value.email);
        const mibaId = validate.mibaId(value.mibaId);
        const user = {};

        if (!(isPlainObject && (hasName || hasEmail || hasMibaId))) fail('user object is invalid');
        if (hasName) user.name = name;
        if (hasEmail) user.email = email;
        if (hasMibaId) user.mibaId = mibaId;

        this._user = user;
    }

    set screen(value) {
        const isPlainObject = check.isPlainObject(value);
        const hasValidWidth = check.isInteger(value.width) && value > 0;
        const hasValidHeight = check.isInteger(value.width) && value > 0;

        if (!(isPlainObject && hasValidWidth &&  hasValidHeight)) fail('screen object is invalid');

        this._screen = value;
    }

    create(data) {
        const complaint = {
            rating: validate.rating(data.rating),
            range: this._keys.range,
            app: this.app,
            platform: this.platform,
            device: this.device,
            browser: this.browser
        };

        if (data.description) data.description = validate.description(data.description);
        if (data.comment) data.comment = validate.comment(data.comment);
        if (this.user) complaint.user = this.user;

        return this.send(complaint);
    }

    send(complaint) {
        const options = {
            method: 'POST',
            headers: new Headers({
                'Content-Type': 'application/json; charset=UTF-8',
                'Accept': 'application/json',
                'Authorization': 'Bearer ' + this._token
            }),
            body: JSON.stringify(complaint),
        };

        return fetch(this._url, options);
    }
}

export default Rating;