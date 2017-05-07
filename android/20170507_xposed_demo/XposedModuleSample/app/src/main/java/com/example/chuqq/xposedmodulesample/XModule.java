package com.example.chuqq.xposedmodulesample;

import android.content.ContentResolver;
import android.provider.Settings;
import android.telephony.TelephonyManager;
import android.util.Log;

import de.robv.android.xposed.IXposedHookLoadPackage;
import de.robv.android.xposed.XC_MethodReplacement;
import de.robv.android.xposed.XposedHelpers;
import de.robv.android.xposed.callbacks.XC_LoadPackage;


/**
 * Created by chuqq on 17/5/6.
 */

public class XModule implements IXposedHookLoadPackage {
    @Override
    public void handleLoadPackage(XC_LoadPackage.LoadPackageParam lpparam) throws Throwable {
        //只hook测试app
        if (lpparam.packageName.equals("com.example.chuqq.xposedmodulesample")) {
//        if (true) {
            XposedHelpers.findAndHookMethod(TelephonyManager.class, "getDeviceId", new XC_MethodReplacement() {
                @Override
                protected Object replaceHookedMethod(MethodHookParam param) throws Throwable {
                    Log.d("chuqq", "replaceHookedMethod: getDeviceId");
                    return "123456789012345";
                }
            });
            XposedHelpers.findAndHookMethod(Settings.Secure.class, "getString", ContentResolver.class, String.class, new XC_MethodReplacement() {
                @Override
                protected Object replaceHookedMethod(MethodHookParam param) throws Throwable {
                    // 暂时没判断第二个参数是Secure.ANDROID_ID
                    Log.d("chuqq", "replaceHookedMethod: getString");
                    return "aaaaaaaaaaaaaaaa";
                }
            });
            XposedHelpers.findAndHookMethod(TelephonyManager.class, "getSubscriberId", new XC_MethodReplacement() {
                @Override
                protected Object replaceHookedMethod(MethodHookParam param) throws Throwable {
                    return "this is imsi";
                }
            });
        }
    }
}
