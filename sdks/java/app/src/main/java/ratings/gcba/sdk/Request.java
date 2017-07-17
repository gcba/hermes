package ratings.gcba.sdk;

import java.util.HashMap;
import android.os.Build;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

class Request {
    public Byte rating;
    public String description;
    public String comment;
    public String range;
    public String app;
    public String platform;
    public HashMap<String, String> device;
    public HashMap<String, Short> screen;
    public HashMap<String, String> user;
}