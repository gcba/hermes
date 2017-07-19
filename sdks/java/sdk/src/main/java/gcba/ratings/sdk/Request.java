package gcba.ratings.sdk;

import java.util.HashMap;
import android.os.Build;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

class Request {
    Request(String app, String platform, String range) {
        this.range = range;

        this.app = new HashMap<String, String>();
        this.app.put("key", app);
        this.app.put("version", String.valueOf(Build.VERSION.RELEASE));

        this.platform = new HashMap<String, String>();
        this.platform.put("key", platform);
        this.platform.put("version", String.valueOf(Build.VERSION.SDK_INT));

        this.device = new HashMap<String, Object>();

        this.user = new HashMap<String, String>();
    }

    public Byte rating;
    public String description;
    public String comment;
    public String range;
    public HashMap<String, String> app;
    public HashMap<String, String> platform;
    public HashMap<String, Object> device;
    public HashMap<String, String> user;

    public String toJSON() {
        return "";
    }
}