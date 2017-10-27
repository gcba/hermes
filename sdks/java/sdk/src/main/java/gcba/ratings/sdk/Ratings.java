package gcba.ratings.sdk;

import com.goebl.david.Response;
import com.goebl.david.Webb;
import com.goebl.david.WebbException;
import com.google.gson.Gson;
import org.json.JSONObject;

import java.util.HashMap;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

public final class Ratings {
    public Ratings(String api, String app, String platform, String range, String token) {
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
        this.timeout = 3000;
    }

    private String url;
    private String app;
    private String platform;
    private String range;
    private String token;
    private Integer timeout;
    private HashMap<String, String> user;

    private void validateUrl(String url) {
        if (url == null) throw new IllegalArgumentException("endpoint can't be null");
        if (!url.trim().matches("^(ftp|http|https):\\/\\/[^ \"]+$")) throw new IllegalArgumentException("invalid endpoint");
    }

    private void validateKey(String key, String description) {
        if (key == null) throw new IllegalArgumentException(description + " can't be null");
        if (key.trim().length() != 32) throw new IllegalArgumentException(description + " is not a valid key");
    }

    private void validateToken(String token) {
        if (token == null) throw new IllegalArgumentException("token can't be null");
        if (token.trim().length() < 1) throw new IllegalArgumentException("invalid token");
    }

    private RatingsError validateName(String name) {
        if (name.trim().length() < 3) return RatingsValidationError.NAME_TOO_SHORT;
        if (name.trim().length() > 70) return RatingsValidationError.NAME_TOO_LONG;

        return null;
    }

    private RatingsError validateEmail(String email) {
        if (!email.trim().matches("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")) return RatingsValidationError.INVALID_EMAIL;
        if (email.trim().length() < 3) return RatingsValidationError.EMAIL_TOO_SHORT;
        if (email.trim().length() > 100) return RatingsValidationError.EMAIL_TOO_LONG;

        return null;
    }

    private RatingsError validateMibaId(String mibaId) {
        if (mibaId.trim().length() != 36) return RatingsValidationError.INVALID_MIBAID;

        return null;
    }

    private RatingsError validateUser(HashMap<String, String> user) {
        String email, mibaId, name;

        name = user.get("name");
        email = user.get("email");
        mibaId = user.get("mibaId");

        if (email == null && mibaId == null) return RatingsValidationError.MISSING_EMAIL_AND_MIBAID;
        if (name != null) return validateName(name);
        if (email != null) return validateEmail(email);

        return validateMibaId(mibaId);
    }

    private RatingsError validateRating(int rating) {
        if (rating < -127 || rating > 127) return RatingsValidationError.INVALID_RATING;

        return null;
    }

    private RatingsError validateDescription(String description) {
        if (description.trim().length() < 3) return RatingsValidationError.DESCRIPTION_TOO_SHORT;
        if (description.trim().length() > 30) return RatingsValidationError.DESCRIPTION_TOO_LONG;

        return null;
    }

    private RatingsError validateComment(String comment) {
        if (comment.trim().length() < 3) return RatingsValidationError.COMMENT_TOO_SHORT;
        if (comment.trim().length() > 1000) return RatingsValidationError.COMMENT_TOO_LONG;

        return null;
    }

    private RatingsRequest getRequest(int rating) {
        RatingsRequest request;

        request = new RatingsRequest();
        request.rating = rating;
        request.range = range;

        request.app.put("key", app);
        request.platform.put("key", platform);

        return request;
    }

    private RatingsResult send(RatingsRequest request) {
        Webb webb;
        Gson gson;
        String json;
        Response<JSONObject> response;
        Response errorResponse;
        RatingsHTTPError error;

        webb = Webb.create();
        gson = new Gson();
        json = gson.toJson(request);

        webb.setBaseUri(url);

        try {
            response = webb
                .post("/ratings")
                .header("Content-Type", "application/json; charset=UTF-8")
                .header("Accept", "application/json")
                .header("Authorization", "Bearer " + token)
                .body(json)
                .connectTimeout(timeout)
                .retry(3, false)
                .asJsonObject();
        } catch (WebbException e) {
            errorResponse = e.getResponse();

            if (errorResponse != null) error = new RatingsHTTPError(errorResponse.getStatusCode(), errorResponse.getStatusLine());
            else error = new RatingsHTTPError(112, e.getMessage());

            return new RatingsResult(null, error);
        }

        return new RatingsResult(response.getBody(), null);
    }

    public void setTimeout(Integer value) {
        timeout = value * 1000;
    }

    public void setUser(String name, String mibaId) {
        setUser(name, mibaId, null);
    }

    public void setUser(String name, String mibaId, String email) {
        HashMap<String, String> newUser;

        newUser = new HashMap<String, String>();

        if (name != null) newUser.put("name", name.trim());
        if (mibaId != null) newUser.put("mibaId", mibaId.trim());
        if (email != null) newUser.put("email", email.trim());

        user = newUser;
    }

    public RatingsResult create(int rating) {
        return create(rating, null, null);
    }

    public RatingsResult create(int rating, String description) {
        return create(rating, description, null);
    }

    public RatingsResult create(int rating, String description, String comment) {
        RatingsRequest request;
        RatingsError error;

        validateRating(rating);

        request = getRequest(rating);

        if (description != null) {
            error = validateDescription(description);

            if (error != null) return new RatingsResult(null, error);

            request.description = description.trim();
        }

        if (comment != null) {
            error = validateComment(comment);

            if (error != null) return new RatingsResult(null, error);

            request.comment = comment.trim();
        }

        if (user != null) {
            error = validateUser(user);

            if (error != null) return new RatingsResult(null, error);

            request.user = user;
        }

        return send(request);
    }
}