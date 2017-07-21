package gcba.ratings.sdk;

import com.goebl.david.Response;
import com.goebl.david.Webb;
import com.google.gson.Gson;
import org.json.JSONObject;

import java.util.HashMap;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

public final class Rating {
    public Rating(String api, String app, String platform, String range, String token) {
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

    private void validateRating(int rating) {
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

    private Request getRequest(int rating) {
        Request request;

        request = new Request();
        request.rating = rating;
        request.range = range;

        request.app.put("key", app);
        request.platform.put("key", platform);

        return request;
    }

    private JSONObject send(Request request) throws Error {
        Webb webb;
        Gson gson;
        String json;
        Response<JSONObject> response;

        webb = Webb.create();
        gson = new Gson();
        json = gson.toJson(request);

        webb.setBaseUri(url);

        try{
            response = webb
                    .post("/ratings")
                    .header("Content-Type", "application/json; charset=UTF-8")
                    .header("Accept", "application/json")
                    .header("Authorization", "Bearer " + token)
                    .body(json)
                    .retry(3, false)
                    .asJsonObject();
        } catch(Error e) {
            throw e;
        }

        return response.getBody();
    }

    public void setUser(String name, String mibaId) {
        setUser(name, mibaId, null);
    }

    public void setUser(String name, String mibaId, String email) {
        HashMap<String, String> newUser;

        if (!(name != null || email != null || mibaId != null)) throw new IllegalArgumentException("user parameters can't all be null");
        if (email == null && mibaId == null) throw new IllegalArgumentException("user has no valid email or mibaId");

        newUser = new HashMap<String, String>();

        if (name != null) {
            validateName(name);
            newUser.put("name", name.trim());
        }

        if (mibaId != null) {
            validateMibaId(mibaId);
            newUser.put("mibaId", mibaId.trim());
        }

        if (email != null) {
            validateEmail(email);
            newUser.put("email", email.trim());
        }

        user = newUser;
    }

    public JSONObject create(int rating) {
        return create(rating, null, null);
    }

    public JSONObject create(int rating, String description) {
        return create(rating, description, null);
    }

    public JSONObject create(int rating, String description, String comment) {
        Request request;

        validateRating(rating);

        request = getRequest(rating);

        if (description != null) {
            validateDescription(description);

            request.description = description.trim();
        }

        if (comment != null) {
            validateComment(comment);

            request.comment = comment.trim();
        }

        if (user != null) {
            request.user = user;
        }

        return send(request);
    }
}