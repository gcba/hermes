package gcba.ratings.example;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import gcba.ratings.sdk.RatingsResult;
import gcba.ratings.sdk.Ratings;

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

                Ratings rating = new Ratings("https://7333ab98.ngrok.io", app, platform, range, token);
                RatingsResult result;

                // Rating only; no user

                result = rating.create(5);

                // Rating and description only; no user

                result = rating.create(4, "Bueno");

                // Rating, description and comment; no user

                result = rating.create(3, "Regular", "Lorem ipsum dolor...");

                // Rating, description and comment; user name and mibaId only

                rating.setUser("Juan Pérez", "dc62591b-1cd3-4c6c-a943-f682e8860e08");

                result = rating.create(2, "Regular", "Lorem ipsum dolor...");

                // Rating, description and comment; user name and email only

                rating.setUser("Juan Pérez", null, "juan@example.com");

                result = rating.create(1, "Muy Malo", "Lorem ipsum dolor...");

                // Rating, description and comment; user name, email and mibaId

                rating.setUser("Juan Pérez", "ae0bfd64-7b37-4bb5-a628-b6cfe28a68af", "juan@perez.com");

                result = rating.create(5, "Muy Bueno", "Lorem ipsum dolor...");
            }
        });

        t.start();
    }
}
