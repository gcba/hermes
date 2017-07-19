package gcba.ratings.sdk;

import java.util.HashMap;
import java.util.StringTokenizer;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

public final class Rating {
    Rating(String api, String app, String platform, String range, String token) {
        validateUrl(api);
        validateKey(app, "app");
        validateKey(platform, "platform");
        validateKey(range, "range");
        validateToken(token);

        this.url = api.trim();
        this.app = app.trim();
        this.platform = platform.trim();
        this.range = range.trim();
        this.token = token.trim();
        this.user = new HashMap<String, String>();
    }

    private String url;
    private String app;
    private String platform;
    private String range;
    private String token;
    private HashMap<String, String> user;

    private void validateUrl(String url) {
        if (url == null) throw new IllegalArgumentException("api can't be null");
        if (!url.trim().matches("^(ftp|http|https):\\/\\/[^ \"]+$")) throw new IllegalArgumentException("invalid api");
    }

    private void validateKey(String key, String description) {
        if (key == null) throw new IllegalArgumentException(description + " can't be null");
        if (key.trim().length() != 32) throw new IllegalArgumentException(description + " is not a valid key");
    }

    private void validateToken(String token) {
        if (token == null) throw new IllegalArgumentException("token can't be null");
        if (token.trim().length() < 1) throw new IllegalArgumentException("invalid token");
    }

    private void validateName(String name) {
        if (name.trim().length() < 3) throw new IllegalArgumentException("name too short");
        if (name.trim().length() > 70) throw new IllegalArgumentException("name too long");
    }

    private void validateEmail(String email) {
        if (!email.trim().matches("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")) throw new IllegalArgumentException("invalid email");
        if (email.trim().length() < 3) throw new IllegalArgumentException("email too short");
        if (email.trim().length() > 100) throw new IllegalArgumentException("email too long");
    }

    private void validateMibaId(String mibaId) {
        if (mibaId.trim().length() < 1) throw new IllegalArgumentException("mibaId too short");
    }

    public void setUser(String name, String email, String mibaId) {
        if (!(name != null || email != null || mibaId != null)) throw new IllegalArgumentException("user parameters can't all be null");
        if (email == null && mibaId == null) throw new IllegalArgumentException("user has no valid email or mibaId");

        if (name != null) {
            validateName(name);
            user.put("name", name.trim());
        }

        if (email != null) {
            validateEmail(email);
            user.put("email", email.trim());
        }

        if (mibaId != null) {
            validateMibaId(mibaId);
            user.put("mibaId", mibaId.trim());
        }
    }

    public void create() {

    }

    public void send() {

    }
}