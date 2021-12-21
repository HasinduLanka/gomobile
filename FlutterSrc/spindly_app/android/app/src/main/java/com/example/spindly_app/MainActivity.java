package com.example.spindly_app;


import android.os.Bundle;
import android.util.Log;

import GoApp.GoApp;
import io.flutter.embedding.android.FlutterActivity;

public class MainActivity extends FlutterActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        GoRun();
    }

    private void GoRun() {
        String cwd = getCacheDir().getAbsolutePath() + "/";
        Thread goThread = new Thread(new Runnable() {
            public void run() {
                Log.d("onCreateTag", "onCreate: 1 + 2 = " + GoApp.add(1, 2));
                GoApp.run(cwd);
            }
        });
        goThread.start();
    }
}
