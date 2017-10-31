package gcba.ratings.example;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;

import gcba.ratings.sdk.Ratings;
import gcba.ratings.sdk.RatingsResult;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_main);

        Thread t = new Thread(new Runnable() {
            public void run() {
                String app = "c33367701511b4f6020ec61ded352059";
                String platform = app;
                String range = app;
                String token = app;

                Ratings sdk = new Ratings("https://75100f11.ngrok.io", token, app, platform, range);
                RatingsResult result;

                // Rating only; no user

                result = sdk.create(5);

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }

                // Rating and description only; no user

                result = sdk.create(4, "Bueno");

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }

                // Rating, description and comment; no user

                result = sdk.create(3, "Regular", "Lorem ipsum dolor...");

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }

                // Rating, description and comment; user name and mibaId only

                sdk.setUser("Juan Pérez", null, "dc62591b-1cd3-4c6c-a943-f682e8860e08");

                result = sdk.create(2, "Regular", "Lorem ipsum dolor...");

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }

                // Rating, description and comment; user name and email only

                sdk.setUser("Juan Pérez", "juan@example.com",null);

                result = sdk.create(1, "Muy Malo", "Lorem ipsum dolor...");

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }

                // Rating, description and comment; user name, email and mibaId

                sdk.setUser("Juan Pérez","juan@perez.com", "ae0bfd64-7b37-4bb5-a628-b6cfe28a68af");

                result = sdk.create(5, "Muy Bueno", "Lorem ipsum dolor...");

                if (result.error != null) {
                    Log.e("RATINGS-ERROR", result.error.getDescription());
                }
            }
        });

        t.start();
    }
}
