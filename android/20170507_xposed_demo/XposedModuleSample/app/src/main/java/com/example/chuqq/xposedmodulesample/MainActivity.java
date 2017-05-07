package com.example.chuqq.xposedmodulesample;

import android.content.Context;
import android.provider.Settings.Secure;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.telephony.TelephonyManager;
import android.util.Log;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        TelephonyManager tm = (TelephonyManager)getSystemService(Context.TELEPHONY_SERVICE);
        String deviceId = tm.getDeviceId();
        Log.d("chuqq ", "onCreate: DEVICE_ID: " + deviceId);

        String androidId = Secure.getString(this.getContentResolver(), Secure.ANDROID_ID);
        Log.d("chuqq ", "onCreate: android_id: " + androidId);
//        Secure.putString(this.getContentResolver(), Secure.ANDROID_ID, "aaaaaaaaaaaaaaaa");

        // 显示在textview中
        TextView textView2 = (TextView)findViewById(R.id.textview2);
        textView2.setText("DEVICE_ID: " + deviceId + "\nANDROID_ID: " + androidId);
    }
}
