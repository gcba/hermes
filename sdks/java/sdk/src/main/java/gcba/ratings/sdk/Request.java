package gcba.ratings.sdk;

import java.util.HashMap;

import android.content.res.Resources;
import android.os.Build;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

final class Request {
    Request() {
        this.app = getApp();
        this.platform = getPlatform();
        this.device = getDevice();
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

    private HashMap<String, String> getApp() {
        HashMap<String, String> app = new HashMap<String, String>();

        app.put("version", String.valueOf(Build.VERSION.RELEASE));

        return app;
    }

    private HashMap<String, String> getPlatform() {
        HashMap<String, String> platform = new HashMap<String, String>();

        platform.put("version", String.valueOf(Build.VERSION.SDK_INT));

        return platform;
    }

    private HashMap<String, Object> getDevice() {
        HashMap<String, Object> device = new HashMap<String, Object>();
        HashMap<String, Integer> screen = new HashMap<String, Integer>();

        screen.put("width", Resources.getSystem().getDisplayMetrics().widthPixels);
        screen.put("height", Resources.getSystem().getDisplayMetrics().heightPixels);
        screen.put("ppi", Math.round(Resources.getSystem().getDisplayMetrics().densityDpi));

        device.put("name", Build.MODEL);
        device.put("brand", Build.BRAND);
        device.put("screen", screen);

        return device;
    }
}