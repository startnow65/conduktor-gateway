package com.hellofresh;

public class Main {
    public static void main(String[] args) {

        GatewayEncryption gatewayEncryption = new GatewayEncryption("http://127.0.0.1:8200", "root", "v1/kv/");
        //String getEncryptionKey = gatewayEncryption.GetEncryptionKeyByUserId("user-id");
        //String getEncryptionKey = gatewayEncryption.GetEncryptionKeyByUserId("user-id-2");
        //String getEncryptionKey = gatewayEncryption.GetEncryptionKeyByUserId("user-id-3");
        //String getEncryptionKey = gatewayEncryption.GetEncryptionKeyByUserId("user-id-4");
        //String getEncryptionKey = gatewayEncryption.GetEncryptionKeyByUserId(randomUUID().toString());
//        if(getEncryptionKey!= null || getEncryptionKey.equals("")){
//            System.out.println("Use this encryption key: " + getEncryptionKey);
//        }else
//        {
//            System.out.println(getEncryptionKey);
//        }


        try {
            String encryptAES = gatewayEncryption.encryptAES("Hello World!", "user-id");
            System.out.println(encryptAES);
            String decryptAES = gatewayEncryption.decryptAES(encryptAES, "user-id");
            System.out.println(decryptAES);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }


    }
}