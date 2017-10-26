package gcba.ratings.sdk;

/**
 * Created by ritazerrizuela on 10/26/17.
 */

enum RatingsValidationError implements RatingsError {
    NAME_TOO_SHORT(100, "name too short"),
    NAME_TOO_LONG(101, "name too long"),
    INVALID_EMAIL(102, "invalid email"),
    EMAIL_TOO_SHORT(103, "email too short"),
    EMAIL_TOO_LONG(104, "email too long"),
    INVALID_MIBAID(105, "invalid mibaId"),
    MISSING_EMAIL_AND_MIBAID(106, "user has no valid email or mibaId"),
    DESCRIPTION_TOO_SHORT(107, "description too short"),
    DESCRIPTION_TOO_LONG(108, "description too long"),
    COMMENT_TOO_SHORT(109, "comment too short"),
    COMMENT_TOO_LONG(110, "comment too long"),
    INVALID_RATING(111, "invalid rating");

    private final int code;
    private final String description;

    RatingsValidationError(int code, String description) {
        this.code = code;
        this.description = description;
    }

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
