package org.hellofresh;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Base64;
import java.io.OutputStream;
import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;

public class GatewayEncryption {
    private static final String VAULT_URL = "http://127.0.0.1:8200"; // Replace with your Vault URL
    private  static String token = "root"; // Replace with your Vault token
    String kvPath = "v1/kv/"; // Replace with the desired Vault API path


    private String callVaultAPI(String token, String path) throws IOException {
        URL url = new URL(VAULT_URL + "/" + path);
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("GET");
        connection.setRequestProperty("X-Vault-Token", token);

        StringBuilder response = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(connection.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                response.append(line);
            }
        }catch (IOException e) {}

        connection.disconnect();


        return response.toString();
    }

    private  void putValueInVault(String token, String path, String value) throws IOException {
        URL url = new URL(VAULT_URL + "/" + path);
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("POST");
        connection.setRequestProperty("X-Vault-Token", token);
        connection.setRequestProperty("Content-Type", "application/json");

        connection.setDoOutput(true);
        String requestPayload = "{\"ciphertext\": \"" + value + "\"}";
        OutputStream outputStream = connection.getOutputStream();
        outputStream.write(requestPayload.getBytes());
        outputStream.flush();

        int responseCode = connection.getResponseCode();
        if (responseCode == HttpURLConnection.HTTP_OK) {
            System.out.println("Value has been successfully put in Vault.");
        } else {
            System.out.println("Error putting value in Vault. Response Code: " + responseCode);
        }

        connection.disconnect();
    }

    public String createDataKey(String token) throws IOException {
        URL url = new URL(VAULT_URL + "/v1/encryption/datakey/plaintext/gateway");
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("POST");
        connection.setRequestProperty("X-Vault-Token", token);
        connection.setDoOutput(true);
        String requestPayload = "{\"plaintext\":\"\"}"; // Empty plaintext indicates creating a new data key
        connection.getOutputStream().write(requestPayload.getBytes());

        StringBuilder response = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(connection.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                response.append(line);
            }
        }

        connection.disconnect();

        return response.toString();
    }

    public static String decryptCiphertext(String token, String ciphertext) throws IOException {
        URL url = new URL(VAULT_URL + "/v1/encryption/decrypt/gateway");
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("POST");
        connection.setRequestProperty("X-Vault-Token", token);

        connection.setDoOutput(true);
        String requestPayload = "{\"ciphertext\":\"" + ciphertext + "\"}";
        connection.getOutputStream().write(requestPayload.getBytes());

        StringBuilder response = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(connection.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                response.append(line);
            }
        }

        connection.disconnect();
        // Extract the plaintext from the response JSON
        ObjectMapper objectMapper = new ObjectMapper();
        JsonNode jsonNode = objectMapper.readTree(response.toString());
        return jsonNode.path("data").path("plaintext").asText();
    }

    public String GetEncryptionKeyByUserId(String userid) {
        System.out.println(userid);
        String response="";
        try {
            response = callVaultAPI(token, kvPath.concat(userid));
            // Extract ciphertext from the response JSON
            ObjectMapper objectMapper = new ObjectMapper();
            JsonNode jsonNode = objectMapper.readTree(response.toString());
            String ciphertext = jsonNode.path("data").path("ciphertext").asText();

            if(ciphertext== null || ciphertext.equals("")){
                String callVaultAPIForDataKey = createDataKey(token);

                // Extract the plaintext from the response JSON
                objectMapper = new ObjectMapper();
                jsonNode = objectMapper.readTree(callVaultAPIForDataKey.toString());
                String plaintext = jsonNode.path("data").path("plaintext").asText();
                ciphertext = jsonNode.path("data").path("ciphertext").asText();

                putValueInVault(token, kvPath.concat(userid), ciphertext);


                return plaintext;

            }else
            {
                return decryptCiphertext(token, ciphertext);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }

        return response;
    }

    public  String encryptAES(String plaintext, String userid) throws Exception {
        String s = GetEncryptionKeyByUserId(userid);
        String adjustKeyLength = adjustKeyLength(String.valueOf(s));
        byte[] keyBytes = adjustKeyLength.getBytes(StandardCharsets.UTF_8);
        SecretKeySpec secretKey = new SecretKeySpec(keyBytes, "AES");
        Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
        cipher.init(Cipher.ENCRYPT_MODE, secretKey);
        byte[] encryptedBytes = cipher.doFinal(plaintext.getBytes(StandardCharsets.UTF_8));
        return Base64.getEncoder().encodeToString(encryptedBytes);
    }

    public  String decryptAES(String ciphertext, String userid) throws Exception {
        String s = GetEncryptionKeyByUserId(userid);
        String adjustKeyLength = adjustKeyLength(String.valueOf(s));
        byte[] keyBytes = adjustKeyLength.getBytes(StandardCharsets.UTF_8);
        SecretKeySpec secretKey = new SecretKeySpec(keyBytes, "AES");
        Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
        cipher.init(Cipher.DECRYPT_MODE, secretKey);
        byte[] decodedBytes = Base64.getDecoder().decode(ciphertext);
        byte[] decryptedBytes = cipher.doFinal(decodedBytes);
        return new String(decryptedBytes, StandardCharsets.UTF_8);
    }
    private String adjustKeyLength(String key) {
        if (key.length() < 16) {
            // Pad the key with zeros to reach the desired length
            while (key.length() < 16) {
                key += "0";
            }
        } else if (key.length() > 16) {
            // Truncate the key to the desired length
            key = key.substring(0, 16);
        }
        return key;
    }
}




