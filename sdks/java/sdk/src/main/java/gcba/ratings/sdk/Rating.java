package gcba.ratings.sdk;

import android.util.Log;

import com.google.gson.Gson;

import java.util.HashMap;

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

        this.url = getUrl(api.trim());
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

    private void validateRating(byte rating) {
        if (rating < -127) throw new IllegalArgumentException("invalid rating");
    }

    private void validateDescription(String description) {
        if (description.trim().length() < 3) throw new IllegalArgumentException("description too short");
        if (description.trim().length() > 30) throw new IllegalArgumentException("description too long");
    }

    private void validateComment(String comment) {
        if (comment.trim().length() < 3) throw new IllegalArgumentException("comment too short");
        if (comment.trim().length() > 1000) throw new IllegalArgumentException("comment too long");
    }

    private String getUrl(String url) {
        return url.charAt(url.length() - 1) == '/' ? url + "ratings"  : url + "/ratings";
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

    public void create(byte rating, String description, String comment) {
        Request request;

        validateRating(rating);

        request = new Request();
        request.rating = rating;
        request.range = range;

        request.app.put("key", app);
        request.platform.put("key", platform);

        if (description != null) {
            validateDescription(description);

            request.description = description.trim();
        }

        if (comment != null) {
            validateComment(comment);

            request.comment = comment.trim();
        }

        send(request);
    }

    private void send(Request request) {
        Gson gson;
        String json;

        gson = new Gson();
        json = gson.toJson(request);

        Log.d("JSON", json);
    }
}