package gcba.ratings.sdk;

/**
 * Created by ritazerrizuela on 10/26/17.
 */

// From https://stackoverflow.com/questions/446663/best-way-to-define-error-codes-strings-in-java

final class RatingsHTTPError implements RatingsError {
    RatingsHTTPError(int code, String description) {
        this.code = code;
        this.description = description;
    }

    private final int code;
    private final String description;

    public String getDescription() {
        return description;
    }

    public int getCode() {
        return code;
    }

    @Override
    public String toString() {
        return code + ": " + description;
    }
}
