package gcba.ratings.sdk;

import java.util.HashMap;

import android.content.res.Resources;
import android.os.Build;

/**
 * Created by ritazerrizuela on 7/17/17.
 */

final class RatingsRequest {
    RatingsRequest() {
        this.app = getApp();
        this.platform = getPlatform();
        this.device = getDevice();
    }

    int rating;
    String description;
    String comment;
    String range;
    HashMap<String, String> app;
    HashMap<String, String> platform;
    HashMap<String, Object> device;
    HashMap<String, String> user;

    private HashMap<String, String> getApp() {
        HashMap<String, String> app = new HashMap<String, String>();

        app.put("version", String.valueOf(BuildConfig.VERSION_CODE));

        return app;
    }

    private HashMap<String, String> getPlatform() {
        HashMap<String, String> platform = new HashMap<String, String>();

        platform.put("version", String.valueOf(Build.VERSION.RELEASE));

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