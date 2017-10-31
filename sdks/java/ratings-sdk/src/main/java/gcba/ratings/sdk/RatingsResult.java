package gcba.ratings.sdk;

import org.json.JSONObject;

/**
 * Created by ritazerrizuela on 10/26/17.
 */

public final class RatingsResult {
    public RatingsResult(JSONObject response, RatingsError error) {
        this.response = response;
        this.error = error;
    }

    public JSONObject response;
    public RatingsError error;
}
