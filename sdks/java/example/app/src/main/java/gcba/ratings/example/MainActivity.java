package gcba.ratings.example;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import gcba.ratings.sdk.Rating;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        setContentView(R.layout.activity_main);

        Thread t = new Thread(new Runnable() {
            public void run() {
                String app = "e10adc3949ba59abbe56e057f20f883e";
                String platform = app;
                String range = app;
                String token = app;

                Rating rating = new Rating("https://0cc1a1ea.ngrok.io", app, platform, range, token);

                rating.create(5);
                rating.create(3, "Regular");
                rating.create(1, "Muy Malo", "No me gustó");

                rating.setUser("Juan Pérez", "e10adc394");

                rating.create(1);
                rating.create(4, "Bueno");
                rating.create(5, "Muy Bueno", "Me encantó");

                rating.setUser("Juan Pérez", "e10adc394", "juan@perez.com");

                rating.create(3);
                rating.create(2, "Malo");
                rating.create(4, "Bueno", "Me gustó");
            }
        });

        t.start();
    }
}
