package com.example.golangdevice;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.os.Bundle;
import android.util.Log;
import android.widget.TextView;
import model.Model;
import android.text.method.ScrollingMovementMethod;





public class MainActivity extends AppCompatActivity {

    private TextView viewGolang;
    private ConnectivityManager cmgr;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        cmgr = (ConnectivityManager)getSystemService(Context.CONNECTIVITY_SERVICE);
        String strisWifi;
        Log.v("bradlog","isNetwork = " + isConnectNetwork());
        Boolean  isNetwork = isConnectNetwork();
        String strNetwork = isNetwork.toString();

        if(isConnectNetwork()){
            //Log.v("bradlog","isWifi = " + isWifiConnected());
            Boolean  isWifi = isWifiConnected();
            strisWifi = isWifi.toString();
        }else {
            Boolean  isWifi = isWifiConnected();
            strisWifi = isWifi.toString();
        }

        viewGolang = (TextView) findViewById(R.id.AviewGolang);
        viewGolang.setMovementMethod(ScrollingMovementMethod.getInstance());
      

        // Call Go function.
        String Platform = Model.platform();
        String Memory = Model.ghwMemory();
        String Usable = Model.ghwUsable();
        String logicalCnt = Model.cpUlogical();
        String physicalCnt = Model.cpuPhy();
        String cpuName = Model.cpuName();
        String cpuPhysicalID = Model.cpuPhysicalID();
        String cpuVendorID = Model.cpuVendorID();
        String cpUmhz = Model.cpUmhz();
        String mac1Address = Model.mac1AddressMask();


        String mac1AddressMask = Model.mac1AddressMask();
        String mac2AddressMask = Model.mac2AddressMask();
        String mac3AddressMask = Model.mac3AddressMask();
        String mac4AddressMask = Model.mac4AddressMask();
        String mac5AddressMask = Model.mac5AddressMask();
        String goarch = Model.goarch();
        String ShellCPUInfo = Model.androidShellCPUInfo();
        String Shellgetprop = Model.androidShellgetprop();
        String IMEI = Model.androidShellIMEI();
        String SreenSize = Model.androidShellSreenSize();
        viewGolang.setText(
                        "Network: "+strNetwork+ "\n\n"+
                        "strisWifi: "+strisWifi+ "\n\n"+

                        "1.Platform: "+Platform+ "\n\n"+
                        "2.Memory: "+Memory+"\n\n"+
                        "3.Usable: "+Usable+"\n\n"+
                        "4.logicalCnt: "+logicalCnt+"\n\n"+
                        "5.physicalCnt: "+physicalCnt+"\n\n"+
                        "6.CPU Name: "+cpuName+"\n\n"+
                        "7.CPU Physical ID: "+cpuPhysicalID+"\n\n"+
                        "8.CPU Vendor ID: "+cpuVendorID+"\n\n"+
                        "9.CPU mhz: "+cpUmhz+"\n\n"+
                        "10.Intranet: "+mac1Address+"\n\n"+
                                "11.mac1AddressMask: "+mac1AddressMask+"\n\n"+
                                "12.mac2AddressMask: "+mac2AddressMask+"\n\n"+
                                "13.mac3AddressMask: "+mac3AddressMask+"\n\n"+
                                "14.mac4AddressMask: "+mac4AddressMask+"\n\n"+
                                "15.mac5AddressMask: "+mac5AddressMask+"\n\n"+
                                "16.system goarch: "+goarch+"\n\n"+
                                "17.android cpuinfo: "+"      "+ShellCPUInfo+"\n\n"+
                                "18.android Shellgetprop: "+"      "+Shellgetprop+"\n\n"+
                                "19.android IMEI: "+"      "+IMEI+"\n\n"+
                                "19.android IMEI: "+"      "+SreenSize+"\n\n"


        );
    }


    private boolean isConnectNetwork(){
        NetworkInfo networkinfo = cmgr.getActiveNetworkInfo();
        return networkinfo != null && networkinfo.isConnectedOrConnecting();
    }
    private boolean isWifiConnected(){
        NetworkInfo networkinfo = cmgr.getNetworkInfo(ConnectivityManager.TYPE_WIFI);
        return networkinfo.isConnected();
    }


}