package gcba.ratings.sdk;

import java.util.HashMap;

import android.content.res.Resources;
import android.os.Build;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

final class Request {
    Request(String app, String platform, String range) {
        if (app == null) throw new IllegalArgumentException("app can't be null");
        if (platform == null) throw new IllegalArgumentException("platform can't be null");
        if (range == null) throw new IllegalArgumentException("range can't be null");

        this.range = range;

        this.app = new HashMap<String, String>();
        this.app.put("key", app);
        this.app.put("version", String.valueOf(Build.VERSION.RELEASE));

        this.platform = new HashMap<String, String>();
        this.platform.put("key", platform);
        this.platform.put("version", String.valueOf(Build.VERSION.SDK_INT));

        HashMap<String, Integer> screen = new HashMap<String, Integer>();
        screen.put("width", Resources.getSystem().getDisplayMetrics().widthPixels);
        screen.put("height", Resources.getSystem().getDisplayMetrics().heightPixels);
        screen.put("ppi", Math.round(Resources.getSystem().getDisplayMetrics().densityDpi));

        this.device = new HashMap<String, Object>();
        this.device.put("name", Build.MODEL);
        this.device.put("brand", Build.BRAND);
        this.device.put("screen", screen);

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
}